package handlers

import (
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
)

type LectureHandler struct {
	LecturePresenter *presenters.LecturePresenter
}

func NewLectureHandler(lecturePresenter *presenters.LecturePresenter) *LectureHandler {
	return &LectureHandler{
		LecturePresenter: lecturePresenter,
	}
}

// Core Handler

func (Presenter *LectureHandler) CreateLectureHandler(ctx *fiber.Ctx) error {

	LectureData, status, messege := Presenter.LecturePresenter.CreateLecturePresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    LectureData,
	})
}

func (Presenter *LectureHandler) GetAllLectureHandler(ctx *fiber.Ctx) error {
	LectureData, status, messege := Presenter.LecturePresenter.GetAllLecturePresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    LectureData,
	})
}
func (Presenter *LectureHandler) GetLectureByIdHandler(ctx *fiber.Ctx) error {
	LectureData, status, messege := Presenter.LecturePresenter.GetLectureByIdPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    LectureData,
	})
}
func (Presenter *LectureHandler) UpdateLectureHandler(ctx *fiber.Ctx) error {
	LectureData, status, messege := Presenter.LecturePresenter.UpdateLecturePresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    LectureData,
	})
}
func (Presenter *LectureHandler) DeleteLecture(ctx *fiber.Ctx) error {
	status, messege := Presenter.LecturePresenter.DeleteLecturePresenter(ctx)

	if status != 204 {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)
}
