package config

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func CreateFiber(log *zap.SugaredLogger) *fiber.App {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	log.Debug("Fiber succesfully configured")

	return app

}
