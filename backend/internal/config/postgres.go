package config

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := "host=localhost user=alvin password=@Alvin123 dbname=zenith port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		Log.Panic("Failed to connect to database: ", err.Error())
	}

	con, err := db.DB()

	if err != nil {
		Log.Panic("Failed to access database: ", err.Error())
	}

	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(100)
	con.SetConnMaxLifetime(time.Second * time.Duration(300))

	return db
}
