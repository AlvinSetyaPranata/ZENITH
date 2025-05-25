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

func (presenter *ReligionPresenter) CreateNewReligionData(ctx *fiber.Ctx) (*model.ReligionModelResponse, int, string) {

	religionRequest := new(model.ReligionyModelRequest)

	religionEntity, status, messege := presenter.ReligionService.CreateReligionData(ctx, religionRequest)

	if status != 201 {
		return nil, status, messege
	}

	response := &model.ReligionModelResponse{
		Name:        religionEntity.Name,
		DateCreated: religionEntity.DateCreated,
	}

	return response, status, messege

}

func (presenter *ReligionPresenter) GetAllReligionData(ctx *fiber.Ctx) (*[]model.ReligionModelResponse, int, string) {

	religionsEntity, status, messege := presenter.ReligionService.GetAllReligionData(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]model.ReligionModelResponse, 0, len(*religionsEntity))

	for _, city := range *religionsEntity {
		response = append(response, model.ReligionModelResponse{
			Name:        city.Name,
			DateCreated: city.DateCreated,
		})
	}

	return &response, 200, "Religions data"

}

func (presenter *ReligionPresenter) GetReligionDataByID(ctx *fiber.Ctx) (*model.ReligionModelResponse, int, string) {

	religionEntity, status, messege := presenter.ReligionService.GetReligionByID(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &model.ReligionModelResponse{
		Name:        religionEntity.Name,
		DateCreated: religionEntity.DateCreated,
	}

	return response, 200, "Religion data with given ID"

}

func (presenter *ReligionPresenter) UpdateReligionData(ctx *fiber.Ctx) (*model.ReligionModelResponse, int, string) {

	religionEntity, status, messege := presenter.ReligionService.Update(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &model.ReligionModelResponse{
		Name:        religionEntity.Name,
		DateCreated: religionEntity.DateCreated,
	}

	return response, 200, "Religion with given id, has been updated successfully!"

}

func (presenter *ReligionPresenter) DeleteReligionData(ctx *fiber.Ctx) (*model.ReligionModelResponse, int, string) {

	religionEntity, status, messege := presenter.ReligionService.Delete(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &model.ReligionModelResponse{
		Name:        religionEntity.Name,
		DateCreated: religionEntity.DateCreated,
	}

	return response, 200, "Religion with given id, has been deleted successfully!"
}
