package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
)

type CityService struct {
	CityRepository *repositories.CityRepository
}

func NewCityService(cityRepository repositories.CityRepository) *CityService {
	return &CityService{
		CityRepository: &cityRepository,
	}
}

// business logic

func (service *CityService) AddCity(ctx context.Context, cityRequest *model.CityModelRequest) (*model.CityModelResponse, int, string) {
	isCreated := service.CityRepository.Create(ctx, *cityRequest)

	if !isCreated {
		return nil, 500, "Server Error"
	}

	return nil, 200, "New city data added"

}

func (service *CityService) GetAllCities(ctx context.Context) *[]entities.City {

	cities := new([]entities.City)

	if err := service.CityRepository.GetAll(ctx, cities); err != nil {
		return nil
	}

	return cities

}
