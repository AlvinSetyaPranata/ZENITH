package routes

import (
	authHandler "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/auth"
	masterHandler "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/master"
	authPresenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/auth"
	masterPresenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	protectedRoute "github.com/AlvinSetyaPranata/ZENITH/backend/api/routes/protected"
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

	// Services
	config.Log.Debug("Boostraping all master data services")

	cityService := masterService.NewCityService(cityRepository, config.Log)
	religionService := masterService.NewReligionService(religionRepository, config.Log)
	genderService := masterService.NewGenderService(genderRepisotory, config.Log)
	facultyService := masterService.NewFacultyService(facultyRepository, config.Log)
	provinceService := masterService.NewProvinceService(provinceRepository, config.Log)
	studyProgramService := masterService.NewStudyProgramService(studyProgramRepository, config.Log)
	countryService := masterService.NewCountryService(countryRepository, config.Log)
	statusService := masterService.NewStatusService(statusRepository, config.Log)

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

	config.Log.Debug("So far, no problem, good!")

	route := protectedRoute.MasterConfig{
		App:                 config.App,
		CityHandler:         cityHandler,
		ReligionHandler:     religionHandler,
		GenderHandler:       genderHandler,
		FacultyHandler:      facultyHandler,
		ProvinceHandler:     provinceHandler,
		StudyProgramHandler: studyProgramHandler,
		CountryHandler:      countryHandler,
		StatusHandler:       statusHandler,
	}

	config.Log.Debug("Setuping master data routers")
	route.SetupMasterRoute()
	config.Log.Debug("Master data routers has been configured succesfully!")

}

func BoostrapAuthRoute(config *BoostrapConfig) {
	// Boostrap auth route

	permissionRepository := authRepository.NewPermissionRepository(config.DB, config.Log)

	// Boostrap auth services
	permissionService := authServices.NewPermissionService(permissionRepository, config.Log)

	// Boostrap auth presenters
	permissionPresenter := authPresenters.NewPermissionPresenter(permissionService, config.Log)

	// Boostrap auth presenters
	permissionHandler := authHandler.NewPermissionHandler(permissionPresenter)

	config.Log.Debug("Auth route successfully boostrapped, yay!")

	config.Log.Debug("Setting up, auth route")
	route := protectedRoute.AuthConfig{
		App:               config.App,
		PermissionHandler: permissionHandler,
	}
	route.SetupAuthRoute()
	config.Log.Debug("Auth route successfully boostrapped, yay!")

}
