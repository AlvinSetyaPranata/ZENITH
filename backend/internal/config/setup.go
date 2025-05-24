package config

import (
	handlers "github.com/AlvinSetyaPranata/ZENITH/backend/api/handlers/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/api/routes"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

type BoostrapConfig struct {
	App *fiber.App
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func Boostrap(config *BoostrapConfig) {
	cityHandler := &handlers.CityHandler{}

	config.Log.Debug("Setting up routers")
	route := routes.RouteConfig{
		App:         config.App,
		CityHandler: cityHandler,
	}

	route.SetupRoute()

}
