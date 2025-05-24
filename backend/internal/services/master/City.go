package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"go.uber.org/zap"
)

type CityService struct {
	CityRepository *repositories.CityRepository
	Log            *zap.SugaredLogger
}

func NewCityService(cityRepository *repositories.CityRepository, log *zap.SugaredLogger) *CityService {
	return &CityService{
		CityRepository: cityRepository,
		Log:            log,
	}
}

// business logic

func (service *CityService) AddCity(ctx context.Context, cityRequest *model.CityModelRequest) (*entities.City, int, string) {
	data, err := service.CityRepository.Create(ctx, *cityRequest)

	if err != nil {
		return nil, 500, "Server Error"
	}

	return data, 201, "New city data added"

}

func (service *CityService) GetAllCities(ctx context.Context) *[]entities.City {

	cities := new([]entities.City)

	if err := service.CityRepository.GetAll(ctx, cities); err != nil {
		return nil
	}

	return cities

}
