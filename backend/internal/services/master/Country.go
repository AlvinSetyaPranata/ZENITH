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

type CountryService struct {
	CountryRepository *repositories.CountryRepository
	Log               *zap.SugaredLogger
}

func NewCountryService(countryRepository *repositories.CountryRepository, log *zap.SugaredLogger) *CountryService {
	return &CountryService{
		CountryRepository: countryRepository,
		Log:               log,
	}
}

// business logic

func (service *CountryService) AddCountry(ctx *fiber.Ctx, countryRequest *model.CountryModelRequest) (*entities.Country, int, string) {

	if err := ctx.BodyParser(countryRequest); err != nil {
		return nil, 400, "Invalid Request"
	}

	data := &entities.Country{
		Name:        countryRequest.Name,
		DateCreated: time.Now(),
	}

	if err := service.CountryRepository.Create(ctx.UserContext(), data); err != nil {
		return nil, 500, "Server Error"
	}

	return data, 201, "New country data added"

}

func (service *CountryService) GetAllCities(ctx context.Context) (*[]entities.Country, int, string) {

	cities := new([]entities.Country)

	if err := service.CountryRepository.GetAll(ctx, cities); err != nil {
		return nil, 500, err.Error()
	}

	return cities, 200, "Data of all country"

}

func (service *CountryService) GetCountryById(ctx context.Context, id string) (*entities.Country, int, string) {
	country := new(entities.Country)

	if err := service.CountryRepository.GetById(ctx, country, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 404, "Country with given id, is not found!"
	}

	return country, 200, "Country data with given id"
}

func (service *CountryService) UpdateCountryData(ctx *fiber.Ctx, countryModelRequest *model.CountryModelRequest, id string) (*entities.Country, int, string) {

	if err := ctx.BodyParser(countryModelRequest); err != nil {
		return nil, 400, "Invalid request"
	}

	country := new(entities.Country)

	// Get entity
	if err := service.CountryRepository.GetById(ctx.UserContext(), country, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 404, "Country with given id, is not found!"
	}

	fmt.Println(country.Id)

	updatedCountry := &entities.Country{
		Id:          country.Id,
		Name:        countryModelRequest.Name,
		DateCreated: country.DateCreated,
	}

	if err := service.CountryRepository.Update(ctx.UserContext(), updatedCountry, id); err != nil {
		return nil, 500, err.Error()
	}

	return updatedCountry, 200, "Country with given id, has been updated successfully!"

}

func (service *CountryService) DeleteCountryData(ctx *fiber.Ctx, id string) (int, string) {

	country := new(entities.Country)

	// Get country data
	if err := service.CountryRepository.GetById(ctx.UserContext(), country, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return 404, "Country with given id, is not found!"
	}

	if err := service.CountryRepository.Delete(ctx.UserContext(), country, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""

}
