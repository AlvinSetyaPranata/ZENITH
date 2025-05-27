package master

import (
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type StatusPresenter struct {
	StatusService *services.StatusService
	Log           *zap.SugaredLogger
}

func NewStatusPresenter(statusService *services.StatusService, log *zap.SugaredLogger) *StatusPresenter {
	return &StatusPresenter{
		StatusService: statusService,
		Log:           log,
	}
}

// Core Presenter

func (presenter *StatusPresenter) CreateNewStatusData(ctx *fiber.Ctx) (*model.StatusModelResponse, int, string) {

	statusModelRequest := new(model.StatusModelRequest)

	statusEntity, status, messege := presenter.StatusService.CreateStatusData(ctx, statusModelRequest)

	if status != 201 {
		return nil, status, messege
	}

	response := &model.StatusModelResponse{
		Id:          statusEntity.Id,
		Name:        statusEntity.Name,
		DateCreated: statusEntity.DateCreated,
	}

	return response, status, messege

}

func (presenter *StatusPresenter) GetAllStatusData(ctx *fiber.Ctx) (*[]model.StatusModelResponse, int, string) {

	statussEntity, status, messege := presenter.StatusService.GetAllStatusData(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]model.StatusModelResponse, 0, len(*statussEntity))

	for _, status := range *statussEntity {
		response = append(response, model.StatusModelResponse{
			Id:          status.Id,
			Name:        status.Name,
			DateCreated: status.DateCreated,
		})
	}

	return &response, 200, "Statuss data"

}

func (presenter *StatusPresenter) GetStatusDataByID(ctx *fiber.Ctx) (*model.StatusModelResponse, int, string) {

	statusEntity, status, messege := presenter.StatusService.GetStatusByID(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &model.StatusModelResponse{
		Id:          statusEntity.Id,
		Name:        statusEntity.Name,
		DateCreated: statusEntity.DateCreated,
	}

	return response, 200, "Status data with given ID"

}

func (presenter *StatusPresenter) UpdateStatusData(ctx *fiber.Ctx) (*model.StatusModelResponse, int, string) {

	statusRequestData := new(model.CityModelRequest)

	statusEntity, status, messege := presenter.StatusService.Update(ctx, (*model.StatusModelRequest)(statusRequestData))

	if status != 200 {
		return nil, status, messege
	}

	response := &model.StatusModelResponse{
		Id:          statusEntity.Id,
		Name:        statusEntity.Name,
		DateCreated: statusEntity.DateCreated,
	}

	return response, 200, "Status with given id, has been updated successfully!"

}

func (presenter *StatusPresenter) DeleteStatusData(ctx *fiber.Ctx) (int, string) {

	status, messege := presenter.StatusService.Delete(ctx)

	if status != 200 {
		return status, messege
	}

	return 204, ""
}
