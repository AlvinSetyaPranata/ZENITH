package master

import (
	"context"

	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
)

type CityPresenter struct {
	Service *services.CityService
}

func NewCityPresenter(service *services.CityService) *CityPresenter {
	return &CityPresenter{
		Service: service,
	}
}

// Presenters

func (p *CityPresenter) GetAllCities(ctx context.Context) *master.CityModelResponse {
	return nil
}
