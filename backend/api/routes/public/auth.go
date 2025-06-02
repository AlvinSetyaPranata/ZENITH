package public

import (
	handlers "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/auth"
	"github.com/gofiber/fiber/v2"
)

type PublicConfig struct {
	App         *fiber.App
	UserHandler *handlers.UserHandler
}

func (config *PublicConfig) SetupPublicRoute() {
	group := config.App.Group("/")

	group.Post("/login", config.UserHandler.LoginHandler)
}
