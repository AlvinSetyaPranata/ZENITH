package handlers

import (
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
)

type SubjectHandler struct {
	SubjectPresenter *presenters.SubjectPresenter
}

func NewSubjectHandler(lecturePresenter *presenters.SubjectPresenter) *SubjectHandler {
	return &SubjectHandler{
		SubjectPresenter: lecturePresenter,
	}
}

// Core Handler

func (Presenter *SubjectHandler) CreateSubjectHandler(ctx *fiber.Ctx) error {

	SubjectData, status, messege := Presenter.SubjectPresenter.CreateSubjectPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    SubjectData,
	})
}

func (Presenter *SubjectHandler) GetAllSubjectHandler(ctx *fiber.Ctx) error {
	SubjectData, status, messege := Presenter.SubjectPresenter.GetAllSubjectPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    SubjectData,
	})
}
func (Presenter *SubjectHandler) GetSubjectByIdHandler(ctx *fiber.Ctx) error {
	SubjectData, status, messege := Presenter.SubjectPresenter.GetSubjectByIdPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    SubjectData,
	})
}
func (Presenter *SubjectHandler) UpdateSubjectHandler(ctx *fiber.Ctx) error {
	SubjectData, status, messege := Presenter.SubjectPresenter.UpdateSubjectPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    SubjectData,
	})
}
func (Presenter *SubjectHandler) DeleteSubject(ctx *fiber.Ctx) error {
	status, messege := Presenter.SubjectPresenter.DeleteSubjectPresenter(ctx)

	if status != 204 {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)
}
