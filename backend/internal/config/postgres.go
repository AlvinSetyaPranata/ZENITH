package config

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(log *zap.SugaredLogger) *gorm.DB {
	dsn := "host=localhost user=alvin password=@Alvin123 dbname=zenith port=5432"

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
