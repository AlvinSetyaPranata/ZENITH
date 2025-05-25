package master

import (
	"context"
	"errors"
	"fmt"
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

func (service *CityService) AddCity(ctx *fiber.Ctx, cityRequest *model.CityModelRequest) (*entities.City, int, string) {

	if err := ctx.BodyParser(cityRequest); err != nil {
		return nil, 400, "Invalid Request"
	}

	data := &entities.City{
		Name:        cityRequest.Name,
		DateCreated: time.Now(),
	}

	if err := service.CityRepository.Create(ctx.UserContext(), data); err != nil {
		return nil, 500, "Server Error"
	}

	return data, 201, "New city data added"

}

func (service *CityService) GetAllCities(ctx context.Context) (*[]entities.City, int, string) {

	cities := new([]entities.City)

	if err := service.CityRepository.GetAll(ctx, cities); err != nil {
		return nil, 500, err.Error()
	}

	return cities, 200, "Data of all city"

}

func (service *CityService) GetCityById(ctx context.Context, id string) (*entities.City, int, string) {
	city := new(entities.City)

	if err := service.CityRepository.GetById(ctx, city, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 404, "City with given id, is not found!"
	}

	return city, 200, "City data with given id"
}

func (service *CityService) UpdateCityData(ctx *fiber.Ctx, cityModelRequest *model.CityModelRequest, id string) (*entities.City, int, string) {

	if err := ctx.BodyParser(cityModelRequest); err != nil {
		return nil, 400, "Invalid request"
	}

	city := new(entities.City)

	// Get entity
	if err := service.CityRepository.GetById(ctx.UserContext(), city, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 404, "City with given id, is not found!"
	}

	fmt.Println(city.Id)

	updatedCity := &entities.City{
		Id:          city.Id,
		Name:        cityModelRequest.Name,
		DateCreated: city.DateCreated,
	}

	if err := service.CityRepository.Update(ctx.UserContext(), updatedCity, id); err != nil {
		return nil, 500, err.Error()
	}

	return updatedCity, 200, "City with given id, has been updated successfully!"

}

func (service *CityService) DeleteCityData(ctx *fiber.Ctx, id string) (*entities.City, int, string) {

	city := new(entities.City)

	// Get city data
	if err := service.CityRepository.GetById(ctx.UserContext(), city, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 404, "City with given id, is not found!"
	}

	if err := service.CityRepository.Delete(ctx.UserContext(), city, id); err != nil {
		return nil, 500, err.Error()
	}

	return city, 200, "City with given id, has deleted successfully!"

}
