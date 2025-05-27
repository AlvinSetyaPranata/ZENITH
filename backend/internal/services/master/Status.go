package master

import (
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	model "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type StatusService struct {
	StatusRepository *repositories.StatusRepository
	Log              *zap.SugaredLogger
}

func NewStatusService(statusRepository *repositories.StatusRepository, log *zap.SugaredLogger) *StatusService {
	return &StatusService{
		StatusRepository: statusRepository,
		Log:              log,
	}
}

// Core Status Service

func (service *StatusService) CreateStatusData(ctx *fiber.Ctx, statusRequest *model.StatusModelRequest) (*entities.Status, int, string) {

	if err := ctx.BodyParser(statusRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	statusEntity := &entities.Status{
		Name:        statusRequest.Name,
		DateCreated: time.Now(),
	}

	if err := service.StatusRepository.Create(ctx.UserContext(), statusEntity); err != nil {
		return nil, 500, err.Error()
	}

	return statusEntity, 201, "New Status data is successfully created!"
}

func (service *StatusService) GetAllStatusData(ctx *fiber.Ctx) (*[]entities.Status, int, string) {

	statusEntities := new([]entities.Status)

	if err := service.StatusRepository.GetAll(ctx.UserContext(), statusEntities); err != nil {
		return nil, 500, err.Error()
	}

	return statusEntities, 200, "Statuss Data"

}

func (service *StatusService) GetStatusByID(ctx *fiber.Ctx) (*entities.Status, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	statusEntity := new(entities.Status)

	if err := service.StatusRepository.GetByID(ctx.UserContext(), statusEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return statusEntity, 200, "Status data by ID"

}

func (service *StatusService) Update(ctx *fiber.Ctx, statusRequest *model.StatusModelRequest) (*entities.Status, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(statusRequest); err != nil {
		return nil, 400, "Invalid Request!"
	}

	// Get current status data

	currentStatusData := new(entities.Status)

	if err := service.StatusRepository.GetByID(ctx.UserContext(), currentStatusData, id); err != nil {
		return nil, 404, "Status with given ID is not found!"
	}

	updatedStatusData := &entities.Status{
		Id:          currentStatusData.Id,
		Name:        statusRequest.Name,
		DateCreated: currentStatusData.DateCreated,
	}

	if err := service.StatusRepository.Update(ctx.UserContext(), updatedStatusData, id); err != nil {
		return nil, 500, err.Error()
	}

	return updatedStatusData, 200, "Status data with given id, has been updated successfully!"

}

func (service *StatusService) Delete(ctx *fiber.Ctx) (int, string) {
	id := ctx.Params("id", "")
	deletedEntity := new(entities.Status)

	if id == "" {
		return 400, "Invalid Request!"
	}

	if err := service.StatusRepository.Delete(ctx.UserContext(), deletedEntity, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}
