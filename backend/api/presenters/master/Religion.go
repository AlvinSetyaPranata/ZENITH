package master

import (
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ReligionPresenter struct {
	ReligionService *services.ReligionService
	Log             *zap.SugaredLogger
}

func NewReligionPresenter(religionService *services.ReligionService, log *zap.SugaredLogger) *ReligionPresenter {
	return &ReligionPresenter{
		ReligionService: religionService,
		Log:             log,
	}
}

// Core Presenter

func (presenter ReligionPresenter) CreateNewReligionData(ctx *fiber.Ctx) (*model.ReligionyModelResponse, int, string) {

	religionRequest := new(model.ReligionyModelRequest)

	religionEntity, status, messege := presenter.ReligionService.CreateReligionData(ctx, religionRequest)

	if status != 201 {
		return nil, status, messege
	}

	response := &model.ReligionyModelResponse{
		Name:        religionEntity.Name,
		DateCreated: religionEntity.DateCreated,
	}

	return response, status, messege

}
