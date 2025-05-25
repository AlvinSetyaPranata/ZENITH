package handlers

import (
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ReligionHandler struct {
	ReligionPresenter *presenters.ReligionPresenter
	Log               *zap.SugaredLogger
}

func NewReligionHandler(religionPresenter *presenters.ReligionPresenter, log *zap.SugaredLogger) *ReligionHandler {
	return &ReligionHandler{
		ReligionPresenter: religionPresenter,
		Log:               log,
	}
}

// Core Religion handlers

func (handler *ReligionHandler) Create(ctx *fiber.Ctx) error {
	data, status, messege := handler.ReligionPresenter.CreateNewReligionData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})

}

func (handler *ReligionHandler) GetAll(ctx *fiber.Ctx) error {
	data, status, messsege := handler.ReligionPresenter.GetAllReligionData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messsege,
		"data":    data,
	})
}

func (handler *ReligionHandler) GetByID(ctx *fiber.Ctx) error {
	data, status, messege := handler.ReligionPresenter.GetReligionDataByID(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}

func (handler *ReligionHandler) Update(ctx *fiber.Ctx) error {
	data, status, messege := handler.ReligionPresenter.UpdateReligionData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}

func (handler *ReligionHandler) Delete(ctx *fiber.Ctx) error {
	data, status, messege := handler.ReligionPresenter.DeleteReligionData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}
