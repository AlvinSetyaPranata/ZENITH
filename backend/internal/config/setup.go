package config

import (
	"github.com/AlvinSetyaPranata/ZENITH/backend/api/routes"
	authEntities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/auth"
	masterEntities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BoostrapConfig struct {
	App *fiber.App
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func MigrateMaster(config *BoostrapConfig) bool {

	config.Log.Debug("Migrating master data schema to database")

	err := config.DB.AutoMigrate(
		&masterEntities.City{},
		&masterEntities.Religion{},
		&masterEntities.Gender{},
		&masterEntities.Faculty{},
		&masterEntities.Province{},
		&masterEntities.StudyProgram{},
		&masterEntities.Status{},
		&masterEntities.Country{},
		&masterEntities.SubjectTime{},
		&masterEntities.Lecture{},
	)

	if err != nil {
		config.Log.Panicf("Error occured during migrating schema: %s", err)
		return false
	}

	config.Log.Debug("Schema migrated succesfully!")

	return true

}

func MigrateAuth(config *BoostrapConfig) bool {
	config.Log.Debug("Migrating auth data schema to database")

	err := config.DB.AutoMigrate(
		&authEntities.Permission{},
		&authEntities.Role{},
		&authEntities.User{},
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

	master_migration_status := MigrateMaster(config)
	auth_migration_status := MigrateAuth(config)

	return master_migration_status && auth_migration_status

}
