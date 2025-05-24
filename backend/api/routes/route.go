package routes

import (
	handlers "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/master"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App         *fiber.App
	CityHandler *handlers.CityHandler
}

func (r *RouteConfig) SetupRoute() {
	r.App.Post("/api/city", r.CityHandler.Create)
	r.App.Get("/api/cities", r.CityHandler.Get)
}
