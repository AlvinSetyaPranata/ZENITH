package handlers

import (
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type GenderHandler struct {
	GenderPresenter *presenters.GenderPresenter
	Log             *zap.SugaredLogger
}

func NewGenderHandler(genderPresenter *presenters.GenderPresenter, log *zap.SugaredLogger) *GenderHandler {
	return &GenderHandler{
		GenderPresenter: genderPresenter,
		Log:             log,
	}
}

// Core Handlers

func (handler *GenderHandler) CreateGenderData(ctx *fiber.Ctx) error {
	response, status, message := handler.GenderPresenter.CreateGenderData(ctx)
	if status != 201 {
		return ctx.Status(status).JSON(fiber.Map{"error": message})
	}
	return ctx.Status(status).JSON(response)
}
func (handler *GenderHandler) GetAllGenderData(ctx *fiber.Ctx) error {
	response, status, message := handler.GenderPresenter.GetAllGenderData(ctx)
	if status != 200 {
		return ctx.Status(status).JSON(fiber.Map{"error": message})
	}
	return ctx.Status(status).JSON(response)
}
func (handler *GenderHandler) GetGenderDataById(ctx *fiber.Ctx) error {
	response, status, message := handler.GenderPresenter.GetGenderDataById(ctx)
	if status != 200 {
		return ctx.Status(status).JSON(fiber.Map{"error": message})
	}
	return ctx.Status(status).JSON(response)
}
func (handler *GenderHandler) UpdateGenderData(ctx *fiber.Ctx) error {
	response, status, message := handler.GenderPresenter.UpdateGenderData(ctx)
	if status != 200 {
		return ctx.Status(status).JSON(fiber.Map{"error": message})
	}
	return ctx.Status(status).JSON(response)
}
func (handler *GenderHandler) DeleteGenderData(ctx *fiber.Ctx) error {
	status, message := handler.GenderPresenter.DeleteGenderData(ctx)
	if status != 200 {
		return ctx.Status(status).JSON(fiber.Map{"error": message})
	}
	return ctx.SendStatus(status)
}
