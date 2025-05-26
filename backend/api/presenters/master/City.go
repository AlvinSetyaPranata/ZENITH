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

func (p *CityPresenter) CreateCityData(ctx *fiber.Ctx) (*master.CityModelResponse, int, string) {

	cityRequestModel := new(master.CityModelRequest)

	data, status, messege := p.Service.AddCity(ctx, cityRequestModel)

	if status != 201 {
		return nil, status, messege
	}

	// Response creation
	response := &master.CityModelResponse{
		Id:          data.Id,
		Name:        data.Name,
		DateCreated: data.DateCreated,
	}

	return response, status, messege
}

func (p *CityPresenter) GetCitiesData(ctx *fiber.Ctx) (*[]master.CityModelResponse, int, string) {

	entities, status, messege := p.Service.GetAllCities(ctx.UserContext())

	if status != 200 {
		return nil, status, messege
	}

	response := make([]master.CityModelResponse, 0, len(*entities))

	for _, city := range *entities {
		response = append(response, master.CityModelResponse{
			Id:          city.Id,
			Name:        city.Name,
			DateCreated: city.DateCreated,
		})
	}

	return &response, status, messege

}

func (p *CityPresenter) GetCityDataById(ctx *fiber.Ctx, id string) (*master.CityModelResponse, int, string) {

	city, status, messege := p.Service.GetCityById(ctx.UserContext(), id)

	if status != 200 {
		return nil, status, messege
	}

	response := &master.CityModelResponse{
		Id:          city.Id,
		Name:        city.Name,
		DateCreated: city.DateCreated,
	}

	return response, status, messege

}

func (p *CityPresenter) UpdateCityData(ctx *fiber.Ctx, id string) (*master.CityModelResponse, int, string) {

	cityModelRequest := new(master.CityModelRequest)

	city, status, messege := p.Service.UpdateCityData(ctx, cityModelRequest, id)

	if status != 200 {
		return nil, status, messege
	}

	response := &master.CityModelResponse{
		Id:          city.Id,
		Name:        city.Name,
		DateCreated: city.DateCreated,
	}

	return response, status, messege

}

func (p *CityPresenter) DeleteCity(ctx *fiber.Ctx, id string) (int, string) {

	status, messege := p.Service.DeleteCityData(ctx, id)

	if status != 204 {
		return status, messege
	}

	return status, ""

}
