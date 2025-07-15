package handlers

import (
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	StudentPresenter *presenters.StudentPresenter
}

func NewStudentHandler(studentPresenter *presenters.StudentPresenter) *StudentHandler {
	return &StudentHandler{
		StudentPresenter: studentPresenter,
	}
}

// Core Handler

func (Presenter *StudentHandler) CreateStudentHandler(ctx *fiber.Ctx) error {

	StudentData, status, messege := Presenter.StudentPresenter.CreateStudentPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StudentData,
	})
}

func (Presenter *StudentHandler) GetAllStudentHandler(ctx *fiber.Ctx) error {
	StudentData, status, messege := Presenter.StudentPresenter.GetAllStudentPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StudentData,
	})
}
func (Presenter *StudentHandler) GetStudentByIdHandler(ctx *fiber.Ctx) error {
	StudentData, status, messege := Presenter.StudentPresenter.GetStudentByIdPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StudentData,
	})
}
func (Presenter *StudentHandler) GetStudentByUserIdHandler(ctx *fiber.Ctx) error {
	StudentData, status, messege := Presenter.StudentPresenter.GetStudentByUserIdPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StudentData,
	})
}
func (Presenter *StudentHandler) UpdateStudentHandler(ctx *fiber.Ctx) error {
	StudentData, status, messege := Presenter.StudentPresenter.UpdateStudentPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StudentData,
	})
}
func (Presenter *StudentHandler) DeleteStudent(ctx *fiber.Ctx) error {
	status, messege := Presenter.StudentPresenter.DeleteStudentPresenter(ctx)

	if status != 204 {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)
}
