package master

import (
	entities "github.com/AlvinSetyaPranata/ZENITH/backend/pkg/entities/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/pkg/repositories"
)

type CountryRepository struct {
	repositories.PostgresRepository[entities.Country]
}

func NewCountryRepository() *CountryRepository {
	return &CountryRepository{}
}
