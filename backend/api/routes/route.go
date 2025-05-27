package routes

import (
	handlers "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/master"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App             *fiber.App
	CityHandler     *handlers.CityHandler
	ReligionHandler *handlers.ReligionHandler
	GenderHandler   *handlers.GenderHandler
	FacultyHandler  *handlers.FacultyHandler
}

func (r *RouteConfig) SetupRoute() {

	// City routes
	r.App.Get("/api/cities", r.CityHandler.Get)
	r.App.Get("/api/city/:id", r.CityHandler.GetById)
	r.App.Post("/api/city", r.CityHandler.Create)
	r.App.Put("/api/city/:id", r.CityHandler.Update)
	r.App.Delete("/api/city/:id", r.CityHandler.Delete)

	// Religion routes
	r.App.Get("/api/religions", r.ReligionHandler.GetAll)
	r.App.Get("/api/religion/:id", r.ReligionHandler.GetByID)
	r.App.Post("/api/religion", r.ReligionHandler.Create)
	r.App.Put("/api/religion/:id", r.ReligionHandler.Update)
	r.App.Delete("/api/religion/:id", r.ReligionHandler.Delete)

	// Gender routes
	r.App.Get("/api/genders", r.GenderHandler.GetAllGenderData)
	r.App.Get("/api/gender/:id", r.GenderHandler.GetGenderDataById)
	r.App.Post("/api/gender", r.GenderHandler.CreateGenderData)
	r.App.Put("/api/gender/:id", r.GenderHandler.UpdateGenderData)
	r.App.Delete("/api/gender/:id", r.GenderHandler.DeleteGenderData)

	// Faculty routes
	r.App.Get("/api/faculties", r.FacultyHandler.GetAllFacultiesHandler)
	r.App.Get("/api/faculty/:id", r.FacultyHandler.GetFacultyByIdHandler)
	r.App.Post("/api/faculty", r.FacultyHandler.CreateFacultyHandler)
	r.App.Put("/api/faculty/:id", r.FacultyHandler.UpdateFacultyHandler)
	r.App.Delete("/api/faculty/:id", r.FacultyHandler.DeleteFaculty)
}
