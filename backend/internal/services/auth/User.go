package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/auth"
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/auth"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/auth"
	utils "github.com/AlvinSetyaPranata/ZENITH/backend/utils"
	authUtils "github.com/AlvinSetyaPranata/ZENITH/backend/utils/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository *repositories.UserRepository
	RoleRepository *repositories.RoleRepository
	Log            *zap.SugaredLogger
}

func NewUserService(userRepository *repositories.UserRepository, roleRepository *repositories.RoleRepository, log *zap.SugaredLogger) *UserService {
	return &UserService{
		UserRepository: userRepository,
		RoleRepository: roleRepository,
		Log:            log,
	}
}

// Core Services
func (Service *UserService) CreateUser(ctx *fiber.Ctx, userRequestModel *models.UserRequestModel) (*entities.User, int, string) {

	if err := ctx.BodyParser(userRequestModel); err != nil {
		return nil, 400, "Invlaid Request!"
	}

	hashedPassword, err := authUtils.HashPassword(userRequestModel.Password)

	if err != nil {
		return nil, 400, "Invalid Request!"
	}

	timeNow := time.Now()

	roleId := strconv.FormatUint(uint64(userRequestModel.Role), 10)

	currentRole := new(entities.Role)

	if err := Service.RoleRepository.GetRoleById(ctx.UserContext(), currentRole, roleId); err != nil {
		return nil, 404, "Role With given id, was not found!"
	}

	userEntity := &entities.User{
		Email:       userRequestModel.Email,
		Password:    hashedPassword,
		Role:        *currentRole,
		DateCreated: timeNow,
		DateUpdated: timeNow,
	}

	// Check if there is user with given role already

	userWithRole := new(entities.User)

	if err := Service.UserRepository.GetUserByRole(ctx.UserContext(), userWithRole, roleId); err != nil {
		if errType := utils.HandlerError(err); errType != 404 {
			return nil, 500, err.Error()
		}
	}

	if userWithRole.Id != 0 && authUtils.CheckRoleCannotDuplicate(currentRole.Name) {
		return nil, 404, "User with following role, is already registered!"
	}

	if err := Service.UserRepository.Create(ctx.UserContext(), userEntity); err != nil {
		if errType := utils.HandlerError(err); errType == 409 {
			return nil, 409, "User with given email is already registered!"
		}
		return nil, 500, err.Error()
	}

	return userEntity, 201, "New user has alread been registered"

}

func (Service *UserService) GetAllUser(ctx *fiber.Ctx) (*[]entities.User, int, string) {

	userEntities := new([]entities.User)

	if err := Service.UserRepository.GetAllUser(ctx.UserContext(), userEntities); err != nil {
		return nil, 500, err.Error()
	}

	return userEntities, 200, "Data for all registered user"

}

func (Service *UserService) GetUserById(ctx *fiber.Ctx) (*entities.User, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	userEntity := new(entities.User)

	if err := Service.UserRepository.GetUserById(ctx.UserContext(), userEntity, id); err != nil {

		if errType := utils.HandlerError(err); errType == 404 {
			return nil, 404, "User with given id is not exists!"
		}

		return nil, 500, err.Error()
	}

	return userEntity, 200, "Data of user from given ID"
}

func (Service *UserService) UpdateUser(ctx *fiber.Ctx, userDataRequest *models.UserRequestModel) (*entities.User, int, string) {
	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(userDataRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	hashedPassword, err := authUtils.HashPassword(userDataRequest.Password)

	if err != nil {
		return nil, 500, err.Error()
	}

	userEntity := &entities.User{
		Email:       userDataRequest.Email,
		Password:    hashedPassword,
		DateUpdated: time.Now(),
	}

	if err := Service.UserRepository.Update(ctx.UserContext(), userEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return userEntity, 204, "User data with given ID, is successfully updated!"

}

func (Service *UserService) DeleteUserService(ctx *fiber.Ctx) (int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request"
	}

	if err := Service.UserRepository.Delete(ctx.UserContext(), &entities.User{}, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}

// Credentials services

func (Service *UserService) LoginService(ctx *fiber.Ctx, loginRequestModel *models.UserCredentialRequestModel) (*entities.User, string, string, int, string) {
	if err := ctx.BodyParser(loginRequestModel); err != nil {
		return nil, "", "", 400, "Invalid Request!"
	}

	currentUser := new(entities.User)

	if err := Service.UserRepository.GetUserByEmail(ctx.UserContext(), currentUser, loginRequestModel.Email); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", "", 404, "User with given email is not found"
		}
	}

	// Check credentials

	if status := authUtils.CheckPassword(loginRequestModel.Password, currentUser.Password); !status {
		return nil, "", "", 401, "Username or Password is incorrect!"
	}

	// Check if the token with given id is still valid
	// If it's then invalidate it

	if err := Service.UserRepository.GetToken(ctx.UserContext(), currentUser.Id); err == nil {
		Service.UserRepository.InvalidateToken(ctx.UserContext(), &entities.Token{
			Revoked: true,
		}, currentUser.Id)
	}

	// Create new one

	userID := strconv.FormatUint(uint64(currentUser.Id), 10)

	access_token, refresh_token, expiredTime := authUtils.GenerateToken(userID, currentUser.Role.Name)

	// Store refresh token to DB

	refreshToken := &entities.Token{
		User:      *currentUser,
		Revoked:   false,
		ExpiredAt: expiredTime,
		CreatedAt: time.Now(),
	}

	if err := Service.UserRepository.StoreToken(ctx.UserContext(), refreshToken); err != nil {
		return nil, "", "", 500, "Internal Server Error!"
	}

	if access_token == "" || refresh_token == "" {
		return nil, "", "", 500, "Internal Server Error!"
	}

	return currentUser, access_token, refresh_token, 200, "Login successfull!"

}

func (Service *UserService) LogoutService(ctx *fiber.Ctx, logoutRequetModel *models.LoogutCredentialModel) (int, string) {
	if err := ctx.BodyParser(logoutRequetModel); err != nil {
		Service.Log.Debugf("Error occured: %s", err.Error())
		return 400, "Bad Request!"
	}

	currentUser := new(entities.User)

	if err := Service.UserRepository.GetUserById(ctx.UserContext(), currentUser, logoutRequetModel.Id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, "User with given email is not found"
		}
		fmt.Println("Error getting user by ID:", err)
		return 500, "Failed to retrieve user information!"
	}

	err := Service.UserRepository.GetToken(ctx.UserContext(), currentUser.Id)

	if err == nil {
		Service.UserRepository.InvalidateToken(ctx.UserContext(), &entities.Token{
			Revoked: true,
		}, currentUser.Id)
		return 200, "Successfully, logged out the user!"
	}

	log.Fatalf("Error checking token existence: %s", err)
	return 500, "Failed to logout user!"
}

func (Service *UserService) RefreshTokenService(ctx *fiber.Ctx) (string, int, string) {

	refresh_token := ctx.Cookies("refresh_token")

	if refresh_token == "" {
		return "", 400, "Refresh Token is invalid!"
	}

	userID, userRole, err := authUtils.ParseToken(refresh_token, false)

	if err != nil {
		return "", 400, err.Error()
	}

	token, tokenErr := authUtils.GenerateNewAccessToken(userID, userRole)

	if tokenErr != nil {
		return "", 500, tokenErr.Error()
	}

	return token, 200, ""

}
