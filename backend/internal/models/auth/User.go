package auth

import (
	"time"
)

type UserCredentialRequestModel struct {
	Email    string `json:"email" validate:"true"`
	Password string `json:"password" validate:"true"`
}

type UserRequestModel struct {
	Email    string `json:"email" validate:"true"`
	Password string `json:"password" validate:"true"`
	Role     uint   `json:"role_id" validate:"true"`
}

type UserResponseModel struct {
	Id          uint      `json:"id" validate:"true"`
	Email       string    `json:"email" validate:"true"`
	RoleId      uint      `json:"role_id" validate:"true"`
	DateCreated time.Time `json:"date_created" validate:"true"`
	DateUpdated time.Time `json:"date_updated" validate:"true"`
}

type UserCredentialResponseModel struct {
	Id    uint               `json:"id" validate:"true"`
	Email string             `json:"email" validate:"true"`
	Role  *RoleModelResponse `json:"role" validate:"true"`
}
