package handlers

import (
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/master"
	"github.com/gofiber/fiber/v2"
)

type StaffHandler struct {
	StaffPresenter *presenters.StaffPresenter
}

func NewStaffHandler(staffPresenter *presenters.StaffPresenter) *StaffHandler {
	return &StaffHandler{
		StaffPresenter: staffPresenter,
	}
}

// Core Handler

func (Presenter *StaffHandler) CreateStaffHandler(ctx *fiber.Ctx) error {

	StaffData, status, messege := Presenter.StaffPresenter.CreateStaffPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StaffData,
	})
}

func (Presenter *StaffHandler) GetAllStaffHandler(ctx *fiber.Ctx) error {
	StaffData, status, messege := Presenter.StaffPresenter.GetAllStaffPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StaffData,
	})
}
func (Presenter *StaffHandler) GetStaffByIdHandler(ctx *fiber.Ctx) error {
	StaffData, status, messege := Presenter.StaffPresenter.GetStaffByIdPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StaffData,
	})
}
func (Presenter *StaffHandler) UpdateStaffHandler(ctx *fiber.Ctx) error {
	StaffData, status, messege := Presenter.StaffPresenter.UpdateStaffPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    StaffData,
	})
}
func (Presenter *StaffHandler) DeleteStaff(ctx *fiber.Ctx) error {
	status, messege := Presenter.StaffPresenter.DeleteStaffPresenter(ctx)

	if status != 204 {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	return ctx.SendStatus(status)
}
