package protected

import (
	handlers "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

type MasterConfig struct {
	App                 *fiber.App
	CityHandler         *handlers.CityHandler
	ReligionHandler     *handlers.ReligionHandler
	GenderHandler       *handlers.GenderHandler
	FacultyHandler      *handlers.FacultyHandler
	ProvinceHandler     *handlers.ProvincesHandler
	StudyProgramHandler *handlers.StudyProgramHandler
	CountryHandler      *handlers.CountryHandler
	StatusHandler       *handlers.StatusHandler
	SubjectTimeHandler  *handlers.SubjectTimeHandler
}

func (r *MasterConfig) SetupMasterRoute() {

	protected := r.App.Group("/api/master", middlewares.AuthenticationCheck)

	// City routes
	protected.Get("/cities", r.CityHandler.Get)
	protected.Get("/city/:id", r.CityHandler.GetById)
	protected.Post("/city", r.CityHandler.Create)
	protected.Put("/city/:id", r.CityHandler.Update)
	protected.Delete("/city/:id", r.CityHandler.Delete)

	// Religion routes
	protected.Get("/religions", r.ReligionHandler.GetAll)
	protected.Get("/religion/:id", r.ReligionHandler.GetByID)
	protected.Post("/religion", r.ReligionHandler.Create)
	protected.Put("/religion/:id", r.ReligionHandler.Update)
	protected.Delete("/religion/:id", r.ReligionHandler.Delete)

	// Gender routes
	protected.Get("/genders", r.GenderHandler.GetAllGenderData)
	protected.Get("/gender/:id", r.GenderHandler.GetGenderDataById)
	protected.Post("/gender", r.GenderHandler.CreateGenderData)
	protected.Put("/gender/:id", r.GenderHandler.UpdateGenderData)
	protected.Delete("/gender/:id", r.GenderHandler.DeleteGenderData)

	// Faculty routes
	protected.Get("/faculties", r.FacultyHandler.GetAllFacultiesHandler)
	protected.Get("/faculty/:id", r.FacultyHandler.GetFacultyByIdHandler)
	protected.Post("/faculty", r.FacultyHandler.CreateFacultyHandler)
	protected.Put("/faculty/:id", r.FacultyHandler.UpdateFacultyHandler)
	protected.Delete("/faculty/:id", r.FacultyHandler.DeleteFaculty)

	// Province routes
	protected.Get("/provinces", r.ProvinceHandler.GetAllProvincesHandler)
	protected.Get("/province/:id", r.ProvinceHandler.GetprovinceByIdHandler)
	protected.Post("/province", r.ProvinceHandler.CreateProvincesHandler)
	protected.Put("/province/:id", r.ProvinceHandler.UpdateProvincesHandler)
	protected.Delete("/province/:id", r.ProvinceHandler.Deleteprovince)

	// Study Program routes
	protected.Get("/study-programs", r.StudyProgramHandler.GetAllStudyProgramHandler)
	protected.Get("/study-program/:id", r.StudyProgramHandler.GetStudyProgramByIdHandler)
	protected.Post("/study-program", r.StudyProgramHandler.CreateStudyProgramHandler)
	protected.Put("/study-program/:id", r.StudyProgramHandler.UpdateStudyProgramHandler)
	protected.Delete("/study-program/:id", r.StudyProgramHandler.DeleteStudyProgram)

	// Country routes
	protected.Get("/countries", r.CountryHandler.GetAllCountriesData)
	protected.Get("/country/:id", r.CountryHandler.GetCountryDataById)
	protected.Post("/country", r.CountryHandler.CreateCountryData)
	protected.Put("/country/:id", r.CountryHandler.UpdateCountryData)
	protected.Delete("/country/:id", r.CountryHandler.DeleteCountryData)

	// Status routes
	protected.Get("/statuses", r.StatusHandler.GetAllStatusesData)
	protected.Get("/status/:id", r.StatusHandler.GetStatusDataByID)
	protected.Post("/status", r.StatusHandler.CreateStatusData)
	protected.Put("/status/:id", r.StatusHandler.UpdateStatusData)
	protected.Delete("/status/:id", r.StatusHandler.DeleteStatusData)

	// Subject Time section route
	protected.Get("/subjectsTime", r.SubjectTimeHandler.GetAllSubjectTimeHandler)
	protected.Get("/subjectTime/:id", r.SubjectTimeHandler.GetSubjectTimeByIdHandler)
	protected.Post("/subjectTime", r.SubjectTimeHandler.CreateSubjectTimeHandler)
	protected.Put("/subjectTime/:id", r.SubjectTimeHandler.UpdateSubjectTimeHandler)
	protected.Delete("/subjectTime/:id", r.SubjectTimeHandler.DeleteSubjectTime)

}
