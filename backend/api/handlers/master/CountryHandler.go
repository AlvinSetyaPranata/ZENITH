package handlers

import (
	"github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CountryHandler struct {
	Presenter *master.CountryPresenter
	Log       *zap.SugaredLogger
}

func NewCountryHandler(presenter *master.CountryPresenter, log *zap.SugaredLogger) *CountryHandler {
	return &CountryHandler{
		Presenter: presenter,
		Log:       log,
	}
}

// Handlers

func (h *CountryHandler) GetAllCountriesData(ctx *fiber.Ctx) error {

	data, status, messege := h.Presenter.GetCitiesData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})

}

func (h *CountryHandler) CreateCountryData(ctx *fiber.Ctx) error {

	data, status, messege := h.Presenter.CreateCountryData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})

}

func (h *CountryHandler) GetCountryDataById(ctx *fiber.Ctx) error {

	// Extract id params
	id := ctx.Params("id", "")

	data, status, messege := h.Presenter.GetCountryDataById(ctx, id)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}

func (h *CountryHandler) UpdateCountryData(ctx *fiber.Ctx) error {

	// Extract id params
	id := ctx.Params("id", "")

	data, status, messege := h.Presenter.UpdateCountryData(ctx, id)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}

func (h *CountryHandler) DeleteCountryData(ctx *fiber.Ctx) error {

	// Extract id params
	id := ctx.Params("id", "")

	status, messege := h.Presenter.DeleteCountry(ctx, id)

	if status != 204 {

		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)

}
