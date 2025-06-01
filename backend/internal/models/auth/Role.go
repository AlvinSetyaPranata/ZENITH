package auth

import (
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/auth"
)

type RoleModelRequest struct {
	Name        string   `json:"name" validate:"true"`
	Permissions []string `json:"permissions" validate:"true"`
}

type RoleModelResponse struct {
	Id          uint                  `json:"id" validate:"true"`
	Name        string                `json:"name" validate:"true"`
	Permissions []entities.Permission `json:"permissions" validate:"true"`
	DateCreated time.Time             `json:"date_created" validate:"true"`
	DateUpdated time.Time             `json:"date_updated" validate:"true"`
}
