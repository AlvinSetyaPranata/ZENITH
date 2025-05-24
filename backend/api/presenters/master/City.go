package master

import (
	"fmt"

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

	if status != 201 || data == nil {
		return status, nil, fmt.Sprintf("Error just occured: %s", messege)
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
