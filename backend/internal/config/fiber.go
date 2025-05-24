package config

import "github.com/gofiber/fiber/v2"

func CreateFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	return app

}
