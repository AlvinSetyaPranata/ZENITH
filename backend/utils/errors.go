package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandlerError(err error) int {
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return fiber.StatusConflict

	case errors.Is(err, gorm.ErrRecordNotFound):
		return fiber.StatusNotFound

	case errors.Is(err, gorm.ErrInvalidData):
		return fiber.StatusBadRequest

	default:
		return fiber.StatusInternalServerError

	}
}
