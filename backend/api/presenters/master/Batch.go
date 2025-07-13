package master

import (
	models "github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	services "github.com/AlvinSetyaPranata/ZENITH/backend/internal/services/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type BatchPresenter struct {
	BatchService *services.BatchService
	Log          *zap.SugaredLogger
}

func NewBatchPresenter(batchService *services.BatchService, log *zap.SugaredLogger) *BatchPresenter {
	return &BatchPresenter{
		BatchService: batchService,
		Log:          log,
	}
}

// Core Presenters

func (Presenter *BatchPresenter) CreateBatchPresenter(ctx *fiber.Ctx) (*models.BatchResponseModel, int, string) {
	batchRequestModel := new(models.BatchRequestModel)

	batchEntity, status, messege := Presenter.BatchService.CreateNewBatch(ctx, batchRequestModel)

	if status != 201 {
		return nil, status, messege
	}

	response := &models.BatchResponseModel{
		Id:          batchEntity.Id,
		Name:        batchEntity.Name,
		StartPeriod: batchEntity.StartPeriod,
		EndPeriod:   batchEntity.EndPeriod,
		DateCreated: batchEntity.DateCreated,
		DateUpdated: batchEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *BatchPresenter) GetAllBatchPresenter(ctx *fiber.Ctx) (*[]models.BatchResponseModel, int, string) {
	batchEntities, status, messege := Presenter.BatchService.GetAllBatch(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := make([]models.BatchResponseModel, 0, len(*batchEntities))

	for _, batch := range *batchEntities {
		response = append(response, models.BatchResponseModel{
			Id:          batch.Id,
			Name:        batch.Name,
			StartPeriod: batch.StartPeriod,
			EndPeriod:   batch.EndPeriod,
			DateCreated: batch.DateCreated,
			DateUpdated: batch.DateUpdated,
		})
	}

	return &response, status, messege
}

func (Presenter *BatchPresenter) GetBatchById(ctx *fiber.Ctx) (*models.BatchResponseModel, int, string) {
	batchEntity, status, messege := Presenter.BatchService.GetBatchById(ctx)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.BatchResponseModel{
		Id:          batchEntity.Id,
		Name:        batchEntity.Name,
		StartPeriod: batchEntity.StartPeriod,
		EndPeriod:   batchEntity.EndPeriod,
		DateCreated: batchEntity.DateCreated,
		DateUpdated: batchEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *BatchPresenter) UpdateBatch(ctx *fiber.Ctx) (*models.BatchResponseModel, int, string) {

	BatchRequestModel := new(models.BatchRequestModel)

	batchEntity, status, messege := Presenter.BatchService.UpdateBatch(ctx, BatchRequestModel)

	if status != 200 {
		return nil, status, messege
	}

	response := &models.BatchResponseModel{
		Id:          batchEntity.Id,
		Name:        batchEntity.Name,
		StartPeriod: batchEntity.StartPeriod,
		EndPeriod:   batchEntity.EndPeriod,
		DateCreated: batchEntity.DateCreated,
		DateUpdated: batchEntity.DateUpdated,
	}

	return response, status, messege
}

func (Presenter *BatchPresenter) DeleteBatch(ctx *fiber.Ctx) (int, string) {

	status, messege := Presenter.BatchService.DeleteBatch(ctx)

	if status != 200 {
		return status, messege
	}

	return status, ""
}
