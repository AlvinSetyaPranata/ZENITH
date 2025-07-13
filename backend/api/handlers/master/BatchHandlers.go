package handlers

import (
	presenter "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type BatchHandler struct {
	BatchPresenter *presenter.BatchPresenter
	Log            *zap.SugaredLogger
}

func NewBatchHandler(BatchPresenter *presenter.BatchPresenter, log *zap.SugaredLogger) *BatchHandler {
	return &BatchHandler{
		BatchPresenter: BatchPresenter,
		Log:            log,
	}
}

// Core Handlers

func (Presenter *BatchHandler) CreateBatchHandler(ctx *fiber.Ctx) error {

	BatchData, status, messege := Presenter.BatchPresenter.CreateBatchPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    BatchData,
	})
}

func (Presenter *BatchHandler) GetAllBatchHandler(ctx *fiber.Ctx) error {
	BatchData, status, messege := Presenter.BatchPresenter.GetAllBatchPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    BatchData,
	})
}
func (Presenter *BatchHandler) GetBatchByIdHandler(ctx *fiber.Ctx) error {
	BatchData, status, messege := Presenter.BatchPresenter.GetBatchById(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    BatchData,
	})
}
func (Presenter *BatchHandler) UpdateBatchHandler(ctx *fiber.Ctx) error {
	BatchData, status, messege := Presenter.BatchPresenter.UpdateBatch(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    BatchData,
	})
}
func (Presenter *BatchHandler) DeleteBatch(ctx *fiber.Ctx) error {
	status, messege := Presenter.BatchPresenter.DeleteBatch(ctx)

	if status != 204 {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)
}
