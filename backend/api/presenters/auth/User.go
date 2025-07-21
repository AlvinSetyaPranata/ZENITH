package auth

import (
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/auth"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/auth"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserPresenter struct {
	UserService *services.UserService
	Log         *zap.SugaredLogger
}

func NewUserPressenter(userService *services.UserService, log *zap.SugaredLogger) *UserPresenter {
	return &UserPresenter{
		UserService: userService,
		Log:         log,
	}
}

// Core Presenters

func (Presenter *UserPresenter) CreatePresenter(ctx *fiber.Ctx) (*models.UserResponseModel, int, string) {

	userRequestModel := new(models.UserRequestModel)

	userEntity, status, messege := Presenter.UserService.CreateUser(ctx, userRequestModel)

	if status != 201 {
		return nil, status, messege
	}

	response := &models.UserResponseModel{
		Id:          userEntity.Id,
		Email:       userEntity.Email,
		RoleId:      userEntity.Role.Id,
		DateCreated: userEntity.DateCreated,
		DateUpdated: userEntity.Role.DateCreated,
	}

	return response, status, messege
}

func (Presenter *UserPresenter) GetAllPresenter(ctx *fiber.Ctx) (*[]models.UserResponseModel, int, string) {
	userEntities, status, messege := Presenter.UserService.GetAllUser(ctx)

	response := make([]models.UserResponseModel, 0, len(*userEntities))

	for _, user := range *userEntities {
		response = append(response, models.UserResponseModel{
			Id:          user.Id,
			Email:       user.Email,
			RoleId:      uint(user.RoleId),
			DateCreated: user.DateCreated,
			DateUpdated: user.DateUpdated,
		})
	}

	return &response, status, messege
}

func (Presenter *UserPresenter) GetByIdPresenter(ctx *fiber.Ctx) (*models.UserResponseModel, int, string) {
	userEntity, status, messege := Presenter.UserService.GetUserById(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.UserResponseModel{
		Id:          userEntity.Id,
		RoleId:      uint(userEntity.RoleId),
		DateCreated: userEntity.DateCreated,
		DateUpdated: userEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *UserPresenter) UpdatePresenter(ctx *fiber.Ctx) (*models.UserResponseModel, int, string) {

	userUpdateRequestModel := new(models.UserRequestModel)

	userEntity, status, messege := Presenter.UserService.UpdateUser(ctx, userUpdateRequestModel)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.UserResponseModel{
		Id:          userEntity.Id,
		Email:       userEntity.Email,
		RoleId:      uint(userEntity.RoleId),
		DateCreated: userEntity.DateCreated,
		DateUpdated: userEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *UserPresenter) DeletePresenter(ctx *fiber.Ctx) (int, string) {

	return Presenter.UserService.DeleteUserService(ctx)

}

// Credentials presenters

func (Presenter *UserPresenter) LoginPresenter(ctx *fiber.Ctx) (*models.UserCredentialResponseModel, string, string, int, string) {

	loginRequestModel := new(models.UserCredentialRequestModel)

	userEntity, access_token, refresh_token, status, messege := Presenter.UserService.LoginService(ctx, loginRequestModel)

	if status != 200 {
		return nil, "", "", status, messege
	}

	permissionsResponse := make([]models.PermissionResponse, len(userEntity.Role.Permissions))

	for i, p := range userEntity.Role.Permissions {
		permissionsResponse[i] = models.PermissionResponse{
			Id:          p.Id,
			Name:        p.Name,
			DateCreated: p.DateCreated,
			DateUpdated: p.DateUpdated,
		}
	}

	response := &models.UserCredentialResponseModel{
		Id:    userEntity.Id,
		Email: userEntity.Email,
		Role: &models.RoleModelResponse{
			Id:          userEntity.Role.Id,
			Name:        userEntity.Role.Name,
			Permissions: permissionsResponse,
			DateCreated: userEntity.DateCreated,
			DateUpdated: userEntity.DateUpdated,
		},
	}

	return response, access_token, refresh_token, status, messege
}

func (Presenter *UserPresenter) LogoutPresenter(ctx *fiber.Ctx) (int, string) {
	logoutRequestModel := new(models.LoogutCredentialModel)

	status, messege := Presenter.UserService.LogoutService(ctx, logoutRequestModel)

	return status, messege

}

func (Presenter *UserPresenter) RefreshTokenPresenter(ctx *fiber.Ctx) (string, int, string) {

	token, status, messege := Presenter.UserService.RefreshTokenService(ctx)

	return token, status, messege

}
