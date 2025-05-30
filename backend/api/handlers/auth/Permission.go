package auth

import (
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/auth"
	"github.com/gofiber/fiber/v2"
)

type PermissionHandler struct {
	PermissionPresenter *presenters.PermissionPresenter
}

func NewPermissionHandler(permissionPresenter *presenters.PermissionPresenter) *PermissionHandler {
	return &PermissionHandler{
		PermissionPresenter: permissionPresenter,
	}
}

// Core handler

func (Handler *PermissionHandler) CreatePermissionHandler(ctx *fiber.Ctx) error {
	data, status, messege := Handler.PermissionPresenter.CreatePermissionPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}

func (Handler *PermissionHandler) GetAllPermissionsHandler(ctx *fiber.Ctx) error {
	data, status, messege := Handler.PermissionPresenter.GetAllPermissionsPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}
func (Handler *PermissionHandler) GetPermissionByIdHandler(ctx *fiber.Ctx) error {
	data, status, messege := Handler.PermissionPresenter.GetPermissionByIdPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}
func (Handler *PermissionHandler) UpdatePermissionHandler(ctx *fiber.Ctx) error {
	data, status, messege := Handler.PermissionPresenter.GetPermissionByIdPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    data,
	})
}
func (Handler *PermissionHandler) DeletePermissionHandler(ctx *fiber.Ctx) error {
	status, messege := Handler.PermissionPresenter.DeletePermissionPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
	})
}
