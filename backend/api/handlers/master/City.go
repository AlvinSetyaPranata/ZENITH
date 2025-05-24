package handlers

import (
	"github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CityHandler struct {
	Presenter *master.CityPresenter
	Log       *zap.SugaredLogger
}

func NewCityHandler(presenter *master.CityPresenter, log *zap.SugaredLogger) *CityHandler {
	return &CityHandler{
		Presenter: presenter,
		Log:       log,
	}
}

// Handlers

func (h *CityHandler) Get(ctx *fiber.Ctx) error {

	status, data, messege := h.Presenter.GetCitiesData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})

}

func (h *CityHandler) Create(ctx *fiber.Ctx) error {

	status, data, messege := h.Presenter.CreateCityData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})

}
