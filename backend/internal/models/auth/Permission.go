package auth

import "time"

type PermissionRequest struct {
	Name string `json:"name" validate:"true"`
}

type PermissionResponse struct {
	Id          string    `json:"id" validate:"true"`
	Name        string    `json:"name" validate:"true"`
	DateCreated time.Time `json:"date_created" validate:"true"`
	DateUpdated time.Time `json:"date_updated" validate:"true"`
}
