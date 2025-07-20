package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func CreateFiber(log *zap.SugaredLogger) *fiber.App {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	cors_config := cors.Config{
		AllowOrigins:     "http://127.0.0.1:3000,http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
	}

	app.Use(cors.New(cors_config))

	log.Debug("Fiber succesfully configured")

	return app

}
