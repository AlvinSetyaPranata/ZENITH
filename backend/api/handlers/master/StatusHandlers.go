package handlers

import (
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type StatusHandler struct {
	StatusPresenter *presenters.StatusPresenter
	Log             *zap.SugaredLogger
}

func NewStatusHandler(statusPresenter *presenters.StatusPresenter, log *zap.SugaredLogger) *StatusHandler {
	return &StatusHandler{
		StatusPresenter: statusPresenter,
		Log:             log,
	}
}

// Core Status handlers

func (handler *StatusHandler) CreateStatusData(ctx *fiber.Ctx) error {

	data, status, messege := handler.StatusPresenter.CreateNewStatusData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})

}

func (handler *StatusHandler) GetAllStatusesData(ctx *fiber.Ctx) error {
	data, status, messsege := handler.StatusPresenter.GetAllStatusData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messsege,
		"data":    data,
	})
}

func (handler *StatusHandler) GetStatusDataByID(ctx *fiber.Ctx) error {
	data, status, messege := handler.StatusPresenter.GetStatusDataByID(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}

func (handler *StatusHandler) UpdateStatusData(ctx *fiber.Ctx) error {
	data, status, messege := handler.StatusPresenter.UpdateStatusData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}

func (handler *StatusHandler) DeleteStatusData(ctx *fiber.Ctx) error {
	status, messege := handler.StatusPresenter.DeleteStatusData(ctx)

	if messege != "" {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)

}
