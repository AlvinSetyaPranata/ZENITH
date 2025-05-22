package config

import "github.com/gofiber/fiber/v2"

func CreateFiber() *fiber.App {
	app := fiber.New()

	return app

}
