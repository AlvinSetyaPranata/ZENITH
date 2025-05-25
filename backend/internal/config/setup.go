package config

import (
	handlers "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/master"
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/api/routes"
	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

type BoostrapConfig struct {
	App *fiber.App
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func BoostrapRoute(config *BoostrapConfig) {

	// Repositories
	config.Log.Debug("Boostraping all repositories")
	cityRepository := repositories.NewCityRepository(config.DB, config.Log)
	religionRepository := repositories.NewReligionRepository(config.DB, config.Log)

	// Services
	config.Log.Debug("Boostraping all repositories")

	cityService := services.NewCityService(cityRepository, config.Log)
	religionService := services.NewReligionService(religionRepository, config.Log)

	// Presenters

	config.Log.Debug("Boostraping all presenters")
	cityPresenter := presenters.NewCityPresenter(cityService, config.Log)
	religionPresenter := presenters.NewReligionPresenter(religionService, config.Log)

	// Handlers
	config.Log.Debug("Boostraping all handlers")
	cityHandler := handlers.NewCityHandler(cityPresenter, config.Log)
	religionHandler := handlers.NewReligionHandler(religionPresenter, config.Log)

	config.Log.Debug("So far, no problem, good!")

	route := routes.RouteConfig{
		App:             config.App,
		CityHandler:     cityHandler,
		ReligionHandler: religionHandler,
	}

	config.Log.Debug("Configuring routers")
	route.SetupRoute()
	config.Log.Debug("Routers has been configured succesfully!")
}

func Migrate(config *BoostrapConfig) bool {

	config.Log.Debug("Migrating schema to database")

	err := config.DB.AutoMigrate(
		&entities.City{},
		&entities.Religion{},
	)

	if err != nil {
		config.Log.Panicf("Error occured during migrating schema: %s", err)
		return false
	}

	config.Log.Debug("Schema migrated succesfully!")

	return true

}

func Boostrap(config *BoostrapConfig) bool {
	BoostrapRoute(config)
	migration_status := Migrate(config)

	return migration_status

}
