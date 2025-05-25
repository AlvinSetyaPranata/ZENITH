package routes

import (
	handlers "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/master"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App             *fiber.App
	CityHandler     *handlers.CityHandler
	ReligionHandler *handlers.ReligionHandler
}

func (r *RouteConfig) SetupRoute() {

	// City routes
	r.App.Post("/api/city", r.CityHandler.Create)
	r.App.Get("/api/cities", r.CityHandler.Get)
	r.App.Get("/api/city/:id", r.CityHandler.GetById)
	r.App.Patch("/api/city/:id", r.CityHandler.Update)
	r.App.Delete("/api/city/:id", r.CityHandler.Delete)

	// Religion routes
	r.App.Post("/api/religion", r.ReligionHandler.Create)
}
