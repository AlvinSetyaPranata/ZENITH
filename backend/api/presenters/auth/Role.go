package auth

import (
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/auth"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/auth"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type RolePresenter struct {
	RoleService *services.RoleService
	Log         *zap.SugaredLogger
}

func NewRolePresenter(roleService *services.RoleService, log *zap.SugaredLogger) *RolePresenter {
	return &RolePresenter{
		RoleService: roleService,
		Log:         log,
	}
}

// Core Presenters

func (Presenter *RolePresenter) CreateRolePresenter(ctx *fiber.Ctx) (*models.RoleModelResponse, int, string) {

	roleCreationRequest := new(models.RoleModelRequest)

	roleEntity, status, messege := Presenter.RoleService.CreateNewRoleData(ctx, roleCreationRequest)

	if status != 201 {
		return nil, status, messege
	}

	response := &models.RoleModelResponse{
		Id:          roleEntity.Id,
		Name:        roleEntity.Name,
		DateCreated: roleEntity.DateCreated,
		DateUpdated: roleEntity.DateCreated,
	}

	return response, status, messege
}

func (Presenter *RolePresenter) GetAllRolesPresenter(ctx *fiber.Ctx) (*[]models.RoleModelResponse, int, string) {
	roleEntity, status, messege := Presenter.RoleService.GetAllRolesData(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]models.RoleModelResponse, 0, len(*roleEntity))

	for _, role := range response {
		response = append(response, models.RoleModelResponse{
			Id:          role.Id,
			Name:        role.Name,
			Permissions: role.Permissions,
			DateCreated: role.DateCreated,
		})
	}

	return &response, status, messege
}

func (Presenter *RolePresenter) GetRoleByIdPresenter(ctx *fiber.Ctx) (*models.RoleModelResponse, int, string) {
	roleEntity, status, messege := Presenter.RoleService.GetRoleById(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.RoleModelResponse{
		Id:          roleEntity.Id,
		Name:        roleEntity.Name,
		DateCreated: roleEntity.DateCreated,
		DateUpdated: roleEntity.DateCreated,
	}

	return response, status, messege
}

func (Presenter *RolePresenter) UpdateRole(ctx *fiber.Ctx) (*models.RoleModelResponse, int, string) {

	roleRequest := new(models.RoleModelRequest)

	roleEntity, status, messege := Presenter.RoleService.UpdateRoleData(ctx, roleRequest)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.RoleModelResponse{
		Id:          roleEntity.Id,
		Name:        roleEntity.Name,
		DateCreated: roleEntity.DateCreated,
		DateUpdated: roleEntity.DateCreated,
	}

	return response, status, messege
}

func (Presenter *RolePresenter) DeleteRole(ctx *fiber.Ctx) (int, string) {
	roleRequest := new(models.RoleModelRequest)

	return Presenter.RoleService.DeleteRoleData(ctx, roleRequest)

}
