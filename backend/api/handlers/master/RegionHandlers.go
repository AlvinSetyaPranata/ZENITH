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
