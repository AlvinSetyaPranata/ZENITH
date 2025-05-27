package handlers

import (
	presenter "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type StudyProgramHandler struct {
	StudyProgramPresenter *presenter.StudyProgramPresenter
	Log                   *zap.SugaredLogger
}

func NewStudyProgramHandler(StudyProgramPresenter *presenter.StudyProgramPresenter, log *zap.SugaredLogger) *StudyProgramHandler {
	return &StudyProgramHandler{
		StudyProgramPresenter: StudyProgramPresenter,
		Log:                   log,
	}
}

// Core Handlers

func (Presenter *StudyProgramHandler) CreateStudyProgramHandler(ctx *fiber.Ctx) error {

	StudyProgramData, status, messege := Presenter.StudyProgramPresenter.CreateStudyProgramPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StudyProgramData,
	})
}

func (Presenter *StudyProgramHandler) GetAllStudyProgramHandler(ctx *fiber.Ctx) error {
	StudyProgramData, status, messege := Presenter.StudyProgramPresenter.GetAllStudyProgramsData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StudyProgramData,
	})
}
func (Presenter *StudyProgramHandler) GetStudyProgramByIdHandler(ctx *fiber.Ctx) error {
	StudyProgramData, status, messege := Presenter.StudyProgramPresenter.GetStudyProgramDataById(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StudyProgramData,
	})
}
func (Presenter *StudyProgramHandler) UpdateStudyProgramHandler(ctx *fiber.Ctx) error {
	StudyProgramData, status, messege := Presenter.StudyProgramPresenter.UpdateStudyProgramData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StudyProgramData,
	})
}
func (Presenter *StudyProgramHandler) DeleteStudyProgram(ctx *fiber.Ctx) error {
	status, messege := Presenter.StudyProgramPresenter.DeleteStudyProgramData(ctx)

	if status != 204 {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)
}
