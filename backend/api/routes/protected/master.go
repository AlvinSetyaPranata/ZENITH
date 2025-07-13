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
	LectureHandler      *handlers.LectureHandler
	StudentHandler      *handlers.StudentHandler
	StaffHandler        *handlers.StaffHandler
	SubjectHandler      *handlers.SubjectHandler
	BatchHandler        *handlers.BatchHandler
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

	// Lecture section route
	protected.Get("/lecturers", r.LectureHandler.GetAllLectureHandler)
	protected.Get("/lecturer/:id", r.LectureHandler.GetLectureByIdHandler)
	protected.Post("/lecturer", r.LectureHandler.CreateLectureHandler)
	protected.Put("/lecturer/:id", r.LectureHandler.UpdateLectureHandler)
	protected.Delete("/lecturer/:id", r.LectureHandler.DeleteLecture)

	// Student section route
	protected.Get("/students", r.StudentHandler.GetAllStudentHandler)
	protected.Get("/student/:id", r.StudentHandler.GetStudentByIdHandler)
	protected.Post("/student", r.StudentHandler.CreateStudentHandler)
	protected.Put("/student/:id", r.StudentHandler.UpdateStudentHandler)
	protected.Delete("/student/:id", r.StudentHandler.DeleteStudent)

	// Staff section route
	protected.Get("/staffs", r.StaffHandler.GetAllStaffHandler)
	protected.Get("/staff/:id", r.StaffHandler.GetStaffByIdHandler)
	protected.Post("/staff", r.StaffHandler.CreateStaffHandler)
	protected.Put("/staff/:id", r.StaffHandler.UpdateStaffHandler)
	protected.Delete("/staff/:id", r.StaffHandler.DeleteStaff)

	// Subject section route
	protected.Get("/subjects", r.SubjectHandler.GetAllSubjectHandler)
	protected.Get("/subject/:id", r.SubjectHandler.GetSubjectByIdHandler)
	protected.Post("/subject", r.SubjectHandler.CreateSubjectHandler)
	protected.Put("/subject/:id", r.SubjectHandler.UpdateSubjectHandler)
	protected.Delete("/subject/:id", r.SubjectHandler.DeleteSubject)

	// Batch section route
	protected.Get("/batches", r.BatchHandler.GetAllBatchHandler)
	protected.Get("/batch/:id", r.BatchHandler.GetBatchByIdHandler)
	protected.Post("/batch", r.BatchHandler.CreateBatchHandler)
	protected.Put("/batch/:id", r.BatchHandler.UpdateBatchHandler)
	protected.Delete("/batch/:id", r.BatchHandler.DeleteBatch)
}
