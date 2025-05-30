package config

import (
	"github.com/AlvinSetyaPranata/ZENITH/backend/api/routes"
	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BoostrapConfig struct {
	App *fiber.App
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func Migrate(config *BoostrapConfig) bool {

	config.Log.Debug("Migrating schema to database")

	err := config.DB.AutoMigrate(
		&entities.City{},
		&entities.Religion{},
		&entities.Gender{},
		&entities.Faculty{},
		&entities.Province{},
		&entities.StudyProgram{},
		&entities.Status{},
		&entities.Country{},
	)

	if err != nil {
		config.Log.Panicf("Error occured during migrating schema: %s", err)
		return false
	}

	config.Log.Debug("Schema migrated succesfully!")

	return true

}

func Boostrap(config *BoostrapConfig) bool {
	config.Log.Debug("Boostrapping protected routes")
	routes.BoostrapMasterRoute(&routes.BoostrapConfig{App: config.App, Log: config.Log, DB: config.DB})
	routes.BoostrapAuthRoute(&routes.BoostrapConfig{App: config.App, Log: config.Log, DB: config.DB})
	config.Log.Debug("All protected routes is successfully boostraped")
	migration_status := Migrate(config)

	return migration_status

}
