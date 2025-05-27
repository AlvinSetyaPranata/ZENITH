package routes

import (
	handlers "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/master"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                 *fiber.App
	CityHandler         *handlers.CityHandler
	ReligionHandler     *handlers.ReligionHandler
	GenderHandler       *handlers.GenderHandler
	FacultyHandler      *handlers.FacultyHandler
	ProvinceHandler     *handlers.ProvincesHandler
	StudyProgramHandler *handlers.StudyProgramHandler
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

	// Province routes
	r.App.Get("/api/provinces", r.ProvinceHandler.GetAllProvincesHandler)
	r.App.Get("/api/province/:id", r.ProvinceHandler.GetprovinceByIdHandler)
	r.App.Post("/api/province", r.ProvinceHandler.CreateProvincesHandler)
	r.App.Put("/api/province/:id", r.ProvinceHandler.UpdateProvincesHandler)
	r.App.Delete("/api/province/:id", r.ProvinceHandler.Deleteprovince)

	// Study Program routes
	r.App.Get("/api/study-programs", r.StudyProgramHandler.GetAllStudyProgramHandler)
	r.App.Get("/api/study-program/:id", r.StudyProgramHandler.GetStudyProgramByIdHandler)
	r.App.Post("/api/study-program", r.StudyProgramHandler.CreateStudyProgramHandler)
	r.App.Put("/api/study-program/:id", r.StudyProgramHandler.UpdateStudyProgramHandler)
	r.App.Delete("/api/study-program/:id", r.StudyProgramHandler.DeleteStudyProgram)
}
