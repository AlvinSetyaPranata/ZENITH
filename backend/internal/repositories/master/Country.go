package master

import (
	// entities "github.com/AlvinSetyaPranata/ZENITH/backend/pkg/entities/master"
	"gorm.io/gorm"
)

type CountryRepository struct {
	DB *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{
		DB: db,
	}
}
