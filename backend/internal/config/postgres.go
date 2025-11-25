package config

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(log *zap.SugaredLogger) *gorm.DB {

	dsn := viper.GetString("DB_DSN")

	log.Debugf("DSN: %s", dsn)

	log.Debug("Opening connection to database")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to database: ", err.Error())
	}

	log.Debug("Creating database instance")
	con, err := db.DB()

	if err != nil {
		log.Panic("Failed to access database: ", err.Error())
	}

	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(100)
	con.SetConnMaxLifetime(time.Second * time.Duration(300))
	log.Debug("Database configured succesfully!")

	return db
}
