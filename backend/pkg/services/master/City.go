package master

import (
	"github.com/AlvinSetyaPranata/ZENITH/backend/pkg/repositories/master"
	"gorm.io/gorm"
)

type CityService struct {
	DB             *gorm.DB
	CityRepository *master.CityRepository
}

func NewCityService(db *gorm.DB, cityRepository master.CityRepository) *CityService {
	return &CityService{
		DB: db,
	}
}

// business logic
