package handlers

import (
	presenter "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type FacultyHandler struct {
	FacultyPresenter *presenter.FacultyPresenter
	Log              *zap.SugaredLogger
}

func NewFacultyHandler(facultyPresenter *presenter.FacultyPresenter, log *zap.SugaredLogger) *FacultyHandler {
	return &FacultyHandler{
		FacultyPresenter: facultyPresenter,
		Log:              log,
	}
}

// Core Handlers

func (Presenter *FacultyHandler) CreateFacultyHandler(ctx *fiber.Ctx) error {

	facultyData, status, messege := Presenter.FacultyPresenter.CreateFacultyPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    facultyData,
	})
}

func (Presenter *FacultyHandler) GetAllFacultiesHandler(ctx *fiber.Ctx) error {
	facultyData, status, messege := Presenter.FacultyPresenter.GetAllFacultiesData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    facultyData,
	})
}
func (Presenter *FacultyHandler) GetFacultyByIdHandler(ctx *fiber.Ctx) error {
	facultyData, status, messege := Presenter.FacultyPresenter.GetFacultyDataById(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    facultyData,
	})
}
func (Presenter *FacultyHandler) UpdateFacultyHandler(ctx *fiber.Ctx) error {
	facultyData, status, messege := Presenter.FacultyPresenter.UpdateFacultyData(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    facultyData,
	})
}
func (Presenter *FacultyHandler) DeleteFaculty(ctx *fiber.Ctx) error {
	status, messege := Presenter.FacultyPresenter.DeleteFacultyData(ctx)

	if status != 204 {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)
}
