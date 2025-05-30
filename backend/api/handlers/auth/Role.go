package auth

import (
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/auth"
	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	RolePresenter *presenters.RolePresenter
}

func NewRoleHandler(rolePresenter *presenters.RolePresenter) *RoleHandler {
	return &RoleHandler{
		RolePresenter: rolePresenter,
	}
}

// Core handlers

func (Presenter *RoleHandler) CreateNewRoleHandler(ctx *fiber.Ctx) error {

	response, status, messege := Presenter.RolePresenter.CreateRolePresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    response,
	})
}

func (Presenter *RoleHandler) GetAllRolesHandler(ctx *fiber.Ctx) error {

	response, status, messege := Presenter.RolePresenter.GetAllRolesPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    response,
	})
}

func (Presenter *RoleHandler) GetRoleById(ctx *fiber.Ctx) error {

	response, status, messege := Presenter.RolePresenter.GetRoleByIdPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    response,
	})
}

func (Presenter *RoleHandler) UpdateRole(ctx *fiber.Ctx) error {

	response, status, messege := Presenter.RolePresenter.UpdateRole(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    response,
	})
}

func (Presenter *RoleHandler) DeleteRole(ctx *fiber.Ctx) error {

	status, _ := Presenter.RolePresenter.DeleteRole(ctx)

	return ctx.SendStatus(status)
}
