package master

import (
	"fmt"
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	repositories "github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type BatchService struct {
	BatchRepository *repositories.BatchRepository
	Log             *zap.SugaredLogger
}

func NewBatchService(batchRepository *repositories.BatchRepository, log *zap.SugaredLogger) *BatchService {
	return &BatchService{
		BatchRepository: batchRepository,
		Log:             log,
	}
}

// Core Services

func (Service *BatchService) CreateNewBatch(ctx *fiber.Ctx, BatchRequestModel *models.BatchRequestModel) (*entities.Batch, int, string) {
	if err := ctx.BodyParser(BatchRequestModel); err != nil {
		fmt.Println(err.Error())
		return nil, 400, "Invalid Request"
	}

	batchEntity := &entities.Batch{
		Name:        BatchRequestModel.Name,
		StartPeriod: BatchRequestModel.StartPeriod,
		EndPeriod:   BatchRequestModel.EndPeriod,
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	}

	if err := Service.BatchRepository.Create(ctx.UserContext(), batchEntity); err != nil {

		if errType := utils.HandlerError(err); errType == 409 {
			return nil, 409, "Batch with given name is already exists"
		}

		return nil, 500, err.Error()
	}

	return batchEntity, 201, "New batch has been added successfully!"
}

func (Service *BatchService) GetAllBatch(ctx *fiber.Ctx) (*[]entities.Batch, int, string) {

	batchEntities := new([]entities.Batch)

	if err := Service.BatchRepository.GetAll(ctx.UserContext(), batchEntities); err != nil {
		return nil, 500, err.Error()
	}

	return batchEntities, 200, "Data of all batch"
}

func (Service *BatchService) GetBatchById(ctx *fiber.Ctx) (*entities.Batch, int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	batchEntitiy := new(entities.Batch)

	if err := Service.BatchRepository.GetById(ctx.UserContext(), batchEntitiy, id); err != nil {
		if errType := utils.HandlerError(err); errType == 404 {
			return nil, 404, "Batch with given id, is not found"
		}

		return nil, 500, err.Error()
	}

	return batchEntitiy, 200, "Data of batch with given id"
}

func (Service *BatchService) UpdateBatch(ctx *fiber.Ctx, BatchRequestModel *models.BatchRequestModel) (*entities.Batch, int, string) {
	id := ctx.Params("id", "")

	if id == "" {
		return nil, 400, "Invalid Request!"
	}

	if err := ctx.BodyParser(BatchRequestModel); err != nil {
		return nil, 400, "Invalid Request!"
	}

	// Get current batch entitiy
	currentBatch := new(entities.Batch)

	if err := Service.BatchRepository.GetById(ctx.UserContext(), currentBatch, id); err != nil {
		return nil, 404, "Batch with given id, is not found!"
	}

	newBatchEntity := &entities.Batch{
		Id:          currentBatch.Id,
		Name:        BatchRequestModel.Name,
		StartPeriod: BatchRequestModel.StartPeriod,
		EndPeriod:   BatchRequestModel.EndPeriod,
		DateCreated: currentBatch.DateCreated,
		DateUpdated: time.Now(),
	}

	if err := Service.BatchRepository.Update(ctx.UserContext(), newBatchEntity, id); err != nil {
		return nil, 500, err.Error()
	}

	return newBatchEntity, 200, "Batch with given id, has been updated successfully!"
}

func (Service *BatchService) DeleteBatch(ctx *fiber.Ctx) (int, string) {

	id := ctx.Params("id", "")

	if id == "" {
		return 400, "Invalid Request!"
	}

	deletedBatchEntity := new(entities.Batch)

	// Search for batch with given id

	if err := Service.BatchRepository.GetById(ctx.UserContext(), deletedBatchEntity, id); err != nil {
		return 404, "Batch with given id, is not found!"
	}

	if err := Service.BatchRepository.Delete(ctx.UserContext(), deletedBatchEntity, id); err != nil {
		return 500, err.Error()
	}

	return 204, ""
}
