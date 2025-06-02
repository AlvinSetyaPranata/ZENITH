package auth

import (
	presenters "github.com/AlvinSetyaPranata/ZENITH/backend/api/presenters/auth"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserPresenter *presenters.UserPresenter
}

func NewUserHandler(userPresenter *presenters.UserPresenter) *UserHandler {
	return &UserHandler{
		UserPresenter: userPresenter,
	}
}

// Core handlers

func (Handler *UserHandler) CreateUserHandler(ctx *fiber.Ctx) error {
	creationResponse, status, messege := Handler.UserPresenter.CreatePresenter(ctx)

	return ctx.Status(status).JSON(&fiber.Map{
		"messege": messege,
		"data":    creationResponse,
	})
}

func (Handler *UserHandler) GetAllUserHandler(ctx *fiber.Ctx) error {
	userData, status, messege := Handler.UserPresenter.GetAllPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    userData,
	})
}

func (Handler *UserHandler) GetUserByIdHandler(ctx *fiber.Ctx) error {
	userData, status, messege := Handler.UserPresenter.GetByIdPresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    userData,
	})
}

func (Handler *UserHandler) UpdateUserHandler(ctx *fiber.Ctx) error {
	userData, status, messege := Handler.UserPresenter.UpdatePresenter(ctx)

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
		"data":    userData,
	})
}

func (Handler *UserHandler) DeleteUserHandler(ctx *fiber.Ctx) error {
	status, messege := Handler.UserPresenter.DeletePresenter(ctx)

	if status != 204 {
		return ctx.SendStatus(status)
	}

	return ctx.Status(status).JSON(fiber.Map{
		"messege": messege,
	})
}

// Authentication handlers

func (Handler *UserHandler) LoginHandler(ctx *fiber.Ctx) error {
	loginResponse, accessToken, refreshToken, status, messege := Handler.UserPresenter.LoginPresenter(ctx)

	if status != 200 {
		return ctx.Status(status).JSON(fiber.Map{
			"messege": messege,
		})
	}

	// Set refresh token inside cookies

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return ctx.Status(status).JSON(fiber.Map{
		"messege":      messege,
		"data":         loginResponse,
		"access_token": accessToken,
	})

}
