package master

import (
	entities "github.com/AlvinSetyaPranata/ZENITH/backend/pkg/entities/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/pkg/repositories"
)

type CityRepository struct {
	repositories.PostgresRepository[entities.City]
}

func NewCityRepository() *CityRepository {
	return &CityRepository{}
}
