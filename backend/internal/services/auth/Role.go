package auth

import (
	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/auth"
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/auth"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/auth"
	"github.com/AlvinSetyaPranata/ZENITH/backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type RoleService struct {
	RoleRepository       *repositories.RoleRepository
	PermissionRepository *repositories.PermissionRepository
	Log                  *zap.SugaredLogger
}

func NewRoleService(roleRepository *repositories.RoleRepository, permissionRepository *repositories.PermissionRepository, log *zap.SugaredLogger) *RoleService {
	return &RoleService{
		RoleRepository:       roleRepository,
		PermissionRepository: permissionRepository,
		Log:                  log,
	}
}

// Core services

func (Service *RoleService) CreateNewRoleData(ctx *fiber.Ctx, roleRequestModel *models.RoleModelRequest) (*entities.Role, *[]entities.Permission, int, string) {

	if err := ctx.BodyParser(roleRequestModel); err != nil {
		return nil, nil, 400, "Invalid Request"
	}

	// get permission

	permissionEntities := make([]entities.Permission, 0, len(roleRequestModel.Permissions))

	permissionEntity := new(entities.Permission)

	for _, permissionName := range roleRequestModel.Permissions {
		if err := Service.PermissionRepository.GetByName(ctx.UserContext(), permissionEntity, permissionName); err != nil {
			return nil, nil, 404, err.Error()
		}
		permissionEntities = append(permissionEntities, *permissionEntity)
		permissionEntity = new(entities.Permission)
	}

	roleEntity := &entities.Role{
		Name:        roleRequestModel.Name,
		Permissions: permissionEntities,
	}

	if err := Service.RoleRepository.Create(ctx.UserContext(), roleEntity); err != nil {

		if errType := utils.HandlerError(err); errType == fiber.StatusConflict {
			return nil, nil, 409, err.Error()
		}

		return nil, nil, 500, err.Error()
	}

	return roleEntity, &permissionEntities, 201, "New role has already created, successfully!"
}

func (Service *RoleService) GetAllRolesData(ctx *fiber.Ctx) (*[]entities.Role, int, string) {

	rolesEntities := new([]entities.Role)

	if err := Service.RoleRepository.GetAllRolesData(ctx.UserContext(), rolesEntities); err != nil {
		return nil, 500, err.Error()
	}

	return rolesEntities, 200, "All roles data"

}

func (Service *RoleService) GetRoleById(ctx *fiber.Ctx) (*entities.Role, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	roleEntity := new(entities.Role)

	if err := Service.RoleRepository.GetRoleById(ctx.UserContext(), roleEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return roleEntity, 200, "Role data with given, ID"
}

func (Service *RoleService) UpdateRoleData(ctx *fiber.Ctx, newRoleDataRequest *models.RoleModelRequest) (*entities.Role, *[]entities.Permission, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(newRoleDataRequest); err != nil {
		return nil, nil, 400, "Invalid Request!"
	}

	// Set the permissions

	currentPermissions := make([]entities.Permission, 0, len(newRoleDataRequest.Permissions))

	currentPermission := new(entities.Permission)

	for _, permissionName := range newRoleDataRequest.Permissions {
		if err := Service.PermissionRepository.GetByName(ctx.UserContext(), currentPermission, permissionName); err != nil {
			return nil, nil, 404, err.Error()
		}

		currentPermissions = append(currentPermissions, *currentPermission)
		currentPermission = new(entities.Permission)
	}

	currentEntity := new(entities.Role)

	if err := Service.RoleRepository.GetRoleById(ctx.UserContext(), currentEntity, id); err != nil {
		if errType := utils.HandlerError(err); errType == fiber.StatusNotFound {
			return nil, nil, 404, "No Role data with given id!"
		}

		return nil, nil, 500, err.Error()
	}

	newRoleEntity := &entities.Role{
		Id:          currentEntity.Id,
		Name:        newRoleDataRequest.Name,
		Permissions: currentPermissions,
		DateCreated: currentEntity.DateCreated,
	}

	if err := Service.RoleRepository.Update(ctx.UserContext(), newRoleEntity, id); err != nil {
		return nil, nil, 500, err.Error()
	}

	return newRoleEntity, &currentPermissions, 200, "New role data with given id, has been updated successfully!"

}

func (Service *RoleService) DeleteRoleData(ctx *fiber.Ctx) (int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request!"
	}

	currentEntity := new(entities.Role)

	if err := Service.RoleRepository.GetRoleById(ctx.UserContext(), currentEntity, id); err != nil {
		if errType := utils.HandlerError(err); errType == fiber.StatusNotFound {
			return 404, "No Role data with given id!"
		}

		return 500, err.Error()
	}

	if err := Service.RoleRepository.Delete(ctx.UserContext(), currentEntity, id); err != nil {
		return 500, err.Error()
	}

	return 204, "Role data with given id, has been deleted successfully!"

}
