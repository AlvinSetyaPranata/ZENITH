package handlers

import (
	presenter "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type SubjectTimeHandler struct {
	SubjectTimePresenter *presenter.SubjectTimePresenter
	Log                  *zap.SugaredLogger
}

func NewSubjectTimeHandler(SubjectTimePresenter *presenter.SubjectTimePresenter, log *zap.SugaredLogger) *SubjectTimeHandler {
	return &SubjectTimeHandler{
		SubjectTimePresenter: SubjectTimePresenter,
		Log:                  log,
	}
}

// Core Handlers

func (Presenter *SubjectTimeHandler) CreateSubjectTimeHandler(ctx *fiber.Ctx) error {

	SubjectTimeData, status, messege := Presenter.SubjectTimePresenter.CreateSubjectTimePresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    SubjectTimeData,
	})
}

func (Presenter *SubjectTimeHandler) GetAllSubjectTimeHandler(ctx *fiber.Ctx) error {
	SubjectTimeData, status, messege := Presenter.SubjectTimePresenter.GetAllSubjectTimePresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    SubjectTimeData,
	})
}
func (Presenter *SubjectTimeHandler) GetSubjectTimeByIdHandler(ctx *fiber.Ctx) error {
	SubjectTimeData, status, messege := Presenter.SubjectTimePresenter.GetSubjectTimeById(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    SubjectTimeData,
	})
}
func (Presenter *SubjectTimeHandler) UpdateSubjectTimeHandler(ctx *fiber.Ctx) error {
	SubjectTimeData, status, messege := Presenter.SubjectTimePresenter.UpdateSubjectTime(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    SubjectTimeData,
	})
}
func (Presenter *SubjectTimeHandler) DeleteSubjectTime(ctx *fiber.Ctx) error {
	status, messege := Presenter.SubjectTimePresenter.DeleteSubjectTime(ctx)

	if status != 204 {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)
}
