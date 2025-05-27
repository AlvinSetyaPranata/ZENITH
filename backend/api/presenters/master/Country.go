package master

import (
	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CountryPresenter struct {
	Service *services.CountryService
	Log     *zap.SugaredLogger
}

func NewCountryPresenter(service *services.CountryService, log *zap.SugaredLogger) *CountryPresenter {
	return &CountryPresenter{
		Service: service,
		Log:     log,
	}
}

// Presenters

func (p *CountryPresenter) CreateCountryData(ctx *fiber.Ctx) (*master.CountryModelResponse, int, string) {

	countryRequestModel := new(master.CountryModelRequest)

	data, status, messege := p.Service.AddCountry(ctx, countryRequestModel)

	if status != 201 {
		return nil, status, messege
	}

	// Response creation
	response := &master.CountryModelResponse{
		Id:          data.Id,
		Name:        data.Name,
		DateCreated: data.DateCreated,
	}

	return response, status, messege
}

func (p *CountryPresenter) GetCitiesData(ctx *fiber.Ctx) (*[]master.CountryModelResponse, int, string) {

	entities, status, messege := p.Service.GetAllCities(ctx.UserContext())

	if status != 200 {
		return nil, status, messege
	}

	response := make([]master.CountryModelResponse, 0, len(*entities))

	for _, country := range *entities {
		response = append(response, master.CountryModelResponse{
			Id:          country.Id,
			Name:        country.Name,
			DateCreated: country.DateCreated,
		})
	}

	return &response, status, messege

}

func (p *CountryPresenter) GetCountryDataById(ctx *fiber.Ctx, id string) (*master.CountryModelResponse, int, string) {

	country, status, messege := p.Service.GetCountryById(ctx.UserContext(), id)

	if status != 200 {
		return nil, status, messege
	}

	response := &master.CountryModelResponse{
		Id:          country.Id,
		Name:        country.Name,
		DateCreated: country.DateCreated,
	}

	return response, status, messege

}

func (p *CountryPresenter) UpdateCountryData(ctx *fiber.Ctx, id string) (*master.CountryModelResponse, int, string) {

	countryModelRequest := new(master.CountryModelRequest)

	country, status, messege := p.Service.UpdateCountryData(ctx, countryModelRequest, id)

	if status != 200 {
		return nil, status, messege
	}

	response := &master.CountryModelResponse{
		Id:          country.Id,
		Name:        country.Name,
		DateCreated: country.DateCreated,
	}

	return response, status, messege

}

func (p *CountryPresenter) DeleteCountry(ctx *fiber.Ctx, id string) (int, string) {

	status, messege := p.Service.DeleteCountryData(ctx, id)

	if status != 204 {
		return status, messege
	}

	return status, ""

}
