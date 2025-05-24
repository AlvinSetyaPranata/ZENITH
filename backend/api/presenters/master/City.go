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
