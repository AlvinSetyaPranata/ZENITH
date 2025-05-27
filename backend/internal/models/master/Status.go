package master

import "time"

type StatusModelRequest struct {
	Name string `json:"name" validate:"true"`
}

type StatusModelResponse struct {
	Id          uint      `json:"id" validate:"true"`
	Name        string    `json:"name" validate:"true"`
	DateCreated time.Time `json:"date_created" validate:"true"`
}
