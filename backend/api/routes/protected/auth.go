package protected

import (
	handlers "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/auth"
	"github.com/gofiber/fiber/v2"
)

type AuthConfig struct {
	App               *fiber.App
	PermissionHandler *handlers.PermissionHandler
}

func (config *AuthConfig) SetupAuthRoute() {

	group := config.App.Group("/api/auth")

	group.Post("/permission", config.PermissionHandler.CreatePermissionHandler)
	group.Get("/permissions", config.PermissionHandler.GetAllPermissionsHandler)
	group.Get("/permission/:id", config.PermissionHandler.GetPermissionByIdHandler)
	group.Put("/permission/:id", config.PermissionHandler.UpdatePermissionHandler)
	group.Delete("/permission/:id", config.PermissionHandler.DeletePermissionHandler)

}
