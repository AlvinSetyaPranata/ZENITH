package auth

import (
	"time"
)

type RoleModelRequest struct {
	Name        string   `json:"name" validate:"true"`
	Permissions []string `json:"permissions" validate:"true"`
}

type RoleModelResponse struct {
	Id          uint                 `json:"id" validate:"true"`
	Name        string               `json:"name" validate:"true"`
	Permissions []PermissionResponse `json:"permissions" validate:"true"`
	DateCreated time.Time            `json:"date_created" validate:"true"`
	DateUpdated time.Time            `json:"date_updated" validate:"true"`
}
