package handlers

import (
	presenter "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ProvincesHandler struct {
	provincePresenter *presenter.ProvincePresenter
	Log               *zap.SugaredLogger
}

func NewProvincesHandler(provincePresenter *presenter.ProvincePresenter, log *zap.SugaredLogger) *ProvincesHandler {
	return &ProvincesHandler{
		provincePresenter: provincePresenter,
		Log:               log,
	}
}

// Core Handlers

func (Presenter *ProvincesHandler) CreateProvincesHandler(ctx *fiber.Ctx) error {

	provinceData, status, messege := Presenter.provincePresenter.CreateProvincePresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    provinceData,
	})
}

func (Presenter *ProvincesHandler) GetAllProvincesHandler(ctx *fiber.Ctx) error {
	provinceData, status, messege := Presenter.provincePresenter.GetAllProvincesData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    provinceData,
	})
}
func (Presenter *ProvincesHandler) GetprovinceByIdHandler(ctx *fiber.Ctx) error {
	provinceData, status, messege := Presenter.provincePresenter.GetProvinceDataById(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    provinceData,
	})
}
func (Presenter *ProvincesHandler) UpdateProvincesHandler(ctx *fiber.Ctx) error {
	provinceData, status, messege := Presenter.provincePresenter.UpdateProvinceData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    provinceData,
	})
}
func (Presenter *ProvincesHandler) Deleteprovince(ctx *fiber.Ctx) error {
	status, messege := Presenter.provincePresenter.DeleteProvinceData(ctx)

	if status != 204 {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)
}
