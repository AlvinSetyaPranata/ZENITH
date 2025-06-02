package protected

import (
	handlers "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/auth"
	"github.com/AlvinSetyaPranata/ZENITH/backend/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

type AuthConfig struct {
	App               *fiber.App
	PermissionHandler *handlers.PermissionHandler
	RoleHandler       *handlers.RoleHandler
	UserHandler       *handlers.UserHandler
}

func (config *AuthConfig) SetupAuthRoute() {

	group := config.App.Group("/api/auth", middlewares.AuthenticationCheck)

	// Permission routes
	group.Post("/permission", config.PermissionHandler.CreatePermissionHandler)
	group.Get("/permissions", config.PermissionHandler.GetAllPermissionsHandler)
	group.Get("/permission/:id", config.PermissionHandler.GetPermissionByIdHandler)
	group.Put("/permission/:id", config.PermissionHandler.UpdatePermissionHandler)
	group.Delete("/permission/:id", config.PermissionHandler.DeletePermissionHandler)

	// Role routes
	group.Post("/role", config.RoleHandler.CreateNewRoleHandler)
	group.Get("/roles", config.RoleHandler.GetAllRolesHandler)
	group.Put("/role/:id", config.RoleHandler.UpdateRole)
	group.Delete("/role/:id", config.RoleHandler.DeleteRole)

	// User routes
	group.Post("/user", config.UserHandler.CreateUserHandler)
	group.Get("/users", config.UserHandler.GetAllUserHandler)
	group.Get("/user/:id", config.UserHandler.GetUserByIdHandler)
	group.Put("/user/:id", config.UserHandler.UpdateUserHandler)
	group.Delete("/user/:id", config.UserHandler.DeleteUserHandler)

}
