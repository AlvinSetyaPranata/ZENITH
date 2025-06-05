package routes

import (
	authHandler "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/auth"
	masterHandler "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/master"
	authPresenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/auth"
	masterPresenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	protectedRoute "github.com/AlvinSetyaPranata/ZENITH/backend/api/routes/protected"
	publicRoute "github.com/AlvinSetyaPranata/ZENITH/backend/api/routes/public"
	authRepository "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/auth"
	masterRepository "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	authServices "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/auth"
	masterService "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BoostrapConfig struct {
	App *fiber.App
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func BoostrapMasterRoute(config *BoostrapConfig) {

	// masterRepository
	config.Log.Debug("Boostraping all master data masterRepository")
	cityRepository := masterRepository.NewCityRepository(config.DB, config.Log)
	religionRepository := masterRepository.NewReligionRepository(config.DB, config.Log)
	genderRepisotory := masterRepository.NewGenderRepository(config.DB, config.Log)
	facultyRepository := masterRepository.NewFacultyRepository(config.DB, config.Log)
	provinceRepository := masterRepository.NewProvinceRepository(config.DB, config.Log)
	studyProgramRepository := masterRepository.NewStudyProgRepository(config.DB, config.Log)
	countryRepository := masterRepository.NewCountryRepository(config.DB, config.Log)
	statusRepository := masterRepository.NewStatusRepository(config.DB, config.Log)
	subjectTimeRepository := masterRepository.NewSubjectTimeRepository(config.DB, config.Log)
	lectureRepository := masterRepository.NewLectureRepository(config.DB)
	studentRepository := masterRepository.NewStudentRepository(config.DB)

	// Services
	config.Log.Debug("Boostraping all master data services")

	cityService := masterService.NewCityService(cityRepository, config.Log)
	religionService := masterService.NewReligionService(religionRepository, config.Log)
	genderService := masterService.NewGenderService(genderRepisotory, config.Log)
	facultyService := masterService.NewFacultyService(facultyRepository, config.Log)
	provinceService := masterService.NewProvinceService(provinceRepository, config.Log)
	studyProgramService := masterService.NewStudyProgramService(studyProgramRepository, facultyRepository, config.Log)
	countryService := masterService.NewCountryService(countryRepository, config.Log)
	statusService := masterService.NewStatusService(statusRepository, config.Log)
	subjectTimeService := masterService.NewSubjectTimeService(subjectTimeRepository, config.Log)
	lectureService := masterService.NewLectureService(lectureRepository, config.Log)
	studentService := masterService.NewStudentService(studentRepository, config.Log)

	// Presenters

	config.Log.Debug("Boostraping all master data presenters")
	cityPresenter := masterPresenters.NewCityPresenter(cityService, config.Log)
	religionPresenter := masterPresenters.NewReligionPresenter(religionService, config.Log)
	genderPresenter := masterPresenters.NewGenderPresenter(genderService, config.Log)
	facultyPresenter := masterPresenters.NewFacultyPresenter(facultyService, config.Log)
	provincePresenter := masterPresenters.NewProvincePresenter(provinceService, config.Log)
	studyProgramPresenter := masterPresenters.NewStudyProgramPresenter(studyProgramService, config.Log)
	countryPresenter := masterPresenters.NewCountryPresenter(countryService, config.Log)
	statusPresenter := masterPresenters.NewStatusPresenter(statusService, config.Log)
	subjectTimePresenter := masterPresenters.NewSubjectTimePresenter(subjectTimeService, config.Log)
	lecturePresenter := masterPresenters.NewLecturePresenter(lectureService)
	studentPresenter := masterPresenters.NewStudentPresenter(studentService)

	// Handlers
	config.Log.Debug("Boostraping all master data handlers")
	cityHandler := masterHandler.NewCityHandler(cityPresenter, config.Log)
	religionHandler := masterHandler.NewReligionHandler(religionPresenter, config.Log)
	genderHandler := masterHandler.NewGenderHandler(genderPresenter, config.Log)
	facultyHandler := masterHandler.NewFacultyHandler(facultyPresenter, config.Log)
	provinceHandler := masterHandler.NewProvincesHandler(provincePresenter, config.Log)
	studyProgramHandler := masterHandler.NewStudyProgramHandler(studyProgramPresenter, config.Log)
	countryHandler := masterHandler.NewCountryHandler(countryPresenter, config.Log)
	statusHandler := masterHandler.NewStatusHandler(statusPresenter, config.Log)
	subjectTimeHandler := masterHandler.NewSubjectTimeHandler(subjectTimePresenter, config.Log)
	lectureHandler := masterHandler.NewLectureHandler(lecturePresenter)
	studentHandler := masterHandler.NewStudentHandler(studentPresenter)

	config.Log.Debug("So far, no problem, good!")

	protectedRouteConfig := protectedRoute.MasterConfig{
		App:                 config.App,
		CityHandler:         cityHandler,
		ReligionHandler:     religionHandler,
		GenderHandler:       genderHandler,
		FacultyHandler:      facultyHandler,
		ProvinceHandler:     provinceHandler,
		StudyProgramHandler: studyProgramHandler,
		CountryHandler:      countryHandler,
		StatusHandler:       statusHandler,
		SubjectTimeHandler:  subjectTimeHandler,
		LectureHandler:      lectureHandler,
		StudentHandler:      studentHandler,
	}

	config.Log.Debug("Setuping master data routers")
	protectedRouteConfig.SetupMasterRoute()
	config.Log.Debug("Master data routers has been configured succesfully!")

}

func BoostrapAuthRoute(config *BoostrapConfig) {
	// Boostrap auth route

	permissionRepository := authRepository.NewPermissionRepository(config.DB, config.Log)
	roleRepository := authRepository.NewRoleRepository(config.DB, config.Log)
	userRepository := authRepository.NewUserRepository(config.DB, config.Log)

	// Boostrap auth services
	permissionService := authServices.NewPermissionService(permissionRepository, config.Log)
	roleService := authServices.NewRoleService(roleRepository, permissionRepository, config.Log)
	userService := authServices.NewUserService(userRepository, roleRepository, config.Log)

	// Boostrap auth presenters
	permissionPresenter := authPresenters.NewPermissionPresenter(permissionService, config.Log)
	rolePresenter := authPresenters.NewRolePresenter(roleService, config.Log)
	userPresenter := authPresenters.NewUserPressenter(userService, config.Log)

	// Boostrap auth handlers
	permissionHandler := authHandler.NewPermissionHandler(permissionPresenter)
	roleHandler := authHandler.NewRoleHandler(rolePresenter)
	userHandler := authHandler.NewUserHandler(userPresenter)

	config.Log.Debug("Auth route successfully boostrapped, yay!")

	config.Log.Debug("Setting up, auth route")
	protectedRouteConfig := protectedRoute.AuthConfig{
		App:               config.App,
		PermissionHandler: permissionHandler,
		RoleHandler:       roleHandler,
		UserHandler:       userHandler,
	}

	publicRouteConfig := publicRoute.PublicConfig{
		App:         config.App,
		UserHandler: userHandler,
	}

	protectedRouteConfig.SetupAuthRoute()
	publicRouteConfig.SetupPublicRoute()
	config.Log.Debug("Auth route successfully boostrapped, yay!")

}
