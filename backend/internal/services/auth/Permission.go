package auth

import (
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/auth"
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/auth"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/auth"
	"github.com/AlvinSetyaPranata/ZENITH/backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type PermissionService struct {
	PermissionRepository *repositories.PermissionRepository
	Log                  *zap.SugaredLogger
}

func NewPermissionService(permissionRepository *repositories.PermissionRepository, log *zap.SugaredLogger) *PermissionService {
	return &PermissionService{
		PermissionRepository: permissionRepository,
		Log:                  log,
	}
}

// Core services

func (Service *PermissionService) CreatePermissionService(ctx *fiber.Ctx, permissionRequest *models.PermissionRequest) (*entities.Permission, int, string) {

	if err := ctx.BodyParser(permissionRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	permissionEntity := &entities.Permission{
		Name:        permissionRequest.Name,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	if err := Service.PermissionRepository.Create(ctx.UserContext(), permissionEntity); err != nil {
		if errType := utils.HandlerError(err); errType == fiber.StatusConflict {
			return nil, 409, "Permission with given name is already registered!"
		}

		return nil, 500, err.Error()
	}

	return permissionEntity, 201, "New permission, has been registered successfully!"

}

func (Service *PermissionService) GetAllPermissionsService(ctx *fiber.Ctx) (*[]entities.Permission, int, string) {

	permissionEntities := new([]entities.Permission)

	if err := Service.PermissionRepository.GetAll(ctx.UserContext(), permissionEntities); err != nil {
		return nil, 500, err.Error()
	}

	return permissionEntities, 200, "Data of all permissions"
}

func (Service *PermissionService) GetPermissionById(ctx *fiber.Ctx) (*entities.Permission, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	permissionEntity := new(entities.Permission)

	if err := Service.PermissionRepository.GetById(ctx.UserContext(), permissionEntity, id); err != nil {
		if errType := utils.HandlerError(err); errType == fiber.StatusNotFound {
			return nil, 404, "Permission with given id, is not found!"
		}
	}

	return permissionEntity, 200, "Data permission with given id"
}

func (Service *PermissionService) UpdatePermissionService(ctx *fiber.Ctx, permissionRequest *models.PermissionRequest) (*entities.Permission, int, string) {
	id := ctx.Params("id", "")

	if err := ctx.BodyParser(permissionRequest); err != nil || id == "" {
		return nil, 400, "Invalid Request!"
	}

	currentPermission := new(entities.Permission)

	if err := Service.PermissionRepository.GetById(ctx.UserContext(), currentPermission, id); err != nil {
		if errType := utils.HandlerError(err); errType == fiber.StatusNotFound {
			return nil, 404, "Permission with given id, is not found!"
		}
	}

	newPermissionEntity := &entities.Permission{
		Name:        permissionRequest.Name,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	if err := Service.PermissionRepository.Update(ctx.UserContext(), newPermissionEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return newPermissionEntity, 200, "Permission with given id, has been updated!"

}

func (Service *PermissionService) DeletePermissionService(ctx *fiber.Ctx) (int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request"
	}

	if err := Service.PermissionRepository.Delete(ctx.UserContext(), &entities.Permission{}, id); err != nil {
		if errType := utils.HandlerError(err); errType == fiber.StatusNotFound {
			return 404, "User with given id, is not found!"
		}

		return 500, err.Error()
	}

	return 204, ""
}
