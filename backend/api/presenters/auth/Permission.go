package auth

import (
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/auth"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/auth"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type PermissionPresenter struct {
	PermissionService *services.PermissionService
	Log               *zap.SugaredLogger
}

func NewPermissionPresenter(permissionService *services.PermissionService, log *zap.SugaredLogger) *PermissionPresenter {
	return &PermissionPresenter{
		PermissionService: permissionService,
		Log:               log,
	}
}

// Core Permission

func (Presenter *PermissionPresenter) CreatePermissionPresenter(ctx *fiber.Ctx) (*models.PermissionResponse, int, string) {

	permissionRequestModel := new(models.PermissionRequest)

	permissionEntity, status, messege := Presenter.PermissionService.CreatePermissionService(ctx, permissionRequestModel)

	if status != 201 {
		return nil, status, messege
	}

	response := &models.PermissionResponse{
		Id:          permissionEntity.Id,
		Name:        permissionEntity.Name,
		DateCreated: permissionEntity.DateCreated,
		DateUpdated: permissionEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *PermissionPresenter) GetAllPermissionsPresenter(ctx *fiber.Ctx) (*[]models.PermissionResponse, int, string) {

	permissionEntities, status, messege := Presenter.PermissionService.GetAllPermissionsService(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]models.PermissionResponse, 0, len(*permissionEntities))

	for _, permission := range *permissionEntities {
		response = append(response, models.PermissionResponse{
			Id:          permission.Id,
			Name:        permission.Name,
			DateCreated: permission.DateCreated,
			DateUpdated: permission.DateUpdated,
		})
	}

	return &response, 200, messege
}

func (Presenter *PermissionPresenter) GetPermissionByIdPresenter(ctx *fiber.Ctx) (*models.PermissionResponse, int, string) {
	permissionEntity, status, messege := Presenter.PermissionService.GetPermissionById(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.PermissionResponse{
		Id:          permissionEntity.Id,
		Name:        permissionEntity.Name,
		DateCreated: permissionEntity.DateCreated,
		DateUpdated: permissionEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *PermissionPresenter) UpdatePermissionPresenter(ctx *fiber.Ctx) (*models.PermissionResponse, int, string) {

	permissionRequestModel := new(models.PermissionRequest)

	permissionEntity, status, messege := Presenter.PermissionService.UpdatePermissionService(ctx, permissionRequestModel)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.PermissionResponse{
		Id:          permissionEntity.Id,
		Name:        permissionEntity.Name,
		DateCreated: permissionEntity.DateCreated,
		DateUpdated: permissionEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *PermissionPresenter) DeletePermissionPresenter(ctx *fiber.Ctx) (int, string) {
	return Presenter.PermissionService.DeletePermissionService(ctx)
}
