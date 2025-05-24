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

	return ctx.Status(200).JSON(fiber.Map{
		"messege": "It works!",
	})

	// request := new(model.CityModelRequest)

	// if err := ctx.BodyParser(request); err != nil {
	// 	return ctx.Status(400).JSON(fiber.Map{
	// 		"messege": err.Error(),
	// 	})
	// }

	// _, status, messege := h.Service.AddCity(ctx.UserContext(), request)

	// return ctx.Status(status).JSON(fiber.Map{
	// 	"messege": messege,
	// })

}

func (h *CityHandler) Create(ctx *fiber.Ctx) error {

	status, data, messege := h.Presenter.CreateCityData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})

}
