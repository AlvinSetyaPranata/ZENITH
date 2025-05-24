package master

import (
	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CityPresenter struct {
	Service *services.CityService
	Log     *zap.SugaredLogger
}

func NewCityPresenter(service *services.CityService, log *zap.SugaredLogger) *CityPresenter {
	return &CityPresenter{
		Service: service,
		Log:     log,
	}
}

// Presenters

func (p *CityPresenter) CreateCityData(ctx *fiber.Ctx) (int, *master.CityModelResponse, string) {

	// Validating request
	request := new(master.CityModelRequest)

	if err := ctx.BodyParser(request); err != nil {
		return 400, nil, "Invalid request"
	}

	data, status, messege := p.Service.AddCity(ctx.UserContext(), request)

	if status != 201 {
		return status, nil, messege
	}

	// Response creation
	response := &master.CityModelResponse{
		Name:        data.Name,
		DateCreated: data.DateCreated,
	}

	return status, response, messege
}

func (p *CityPresenter) GetCitiesData(ctx *fiber.Ctx) (int, *[]master.CityModelResponse, string) {
	entities, status, messege := p.Service.GetAllCities(ctx.UserContext())

	if status != 200 {
		return status, nil, messege
	}

	response := make([]master.CityModelResponse, 0, len(*entities))

	for _, city := range *entities {
		response = append(response, master.CityModelResponse{
			Name:        city.Name,
			DateCreated: city.DateCreated,
		})
	}

	return status, &response, messege

}

func (p *CityPresenter) GetCityDataById(ctx *fiber.Ctx, id string) (int, *master.CityModelResponse, string) {

	city, status, messege := p.Service.GetCityById(ctx.UserContext(), id)

	if status != 200 {
		return status, nil, messege
	}

	response := &master.CityModelResponse{
		Name:        city.Name,
		DateCreated: city.DateCreated,
	}

	return status, response, messege

}

func (p *CityPresenter) UpdateCityData(ctx *fiber.Ctx, id string) (int, *master.CityModelResponse, string) {

	request := new(master.CityModelRequest)

	if err := ctx.BodyParser(request); err != nil {
		return 400, nil, "Invalid request"
	}

	city, status, messege := p.Service.UpdateCityData(ctx.UserContext(), request, id)

	if status != 200 {
		return status, nil, messege
	}

	response := &master.CityModelResponse{
		Name:        city.Name,
		DateCreated: city.DateCreated,
	}

	return status, response, messege

}

func (p *CityPresenter) DeleteCity(ctx *fiber.Ctx, id string) (int, *master.CityModelResponse, string) {

	city, status, messege := p.Service.DeleteCityData(ctx.UserContext(), id)

	if status != 200 {
		return status, nil, messege
	}

	response := &master.CityModelResponse{
		Name:        city.Name,
		DateCreated: city.DateCreated,
	}

	return status, response, messege

}
