package master

import "time"

type ReligionyModelRequest struct {
	Name string `json:"name" validate:"true"`
}

type ReligionyModelResponse struct {
	Name        string    `json:"name" validate:"true"`
	DateCreated time.Time `json:"date_created" validate:"true"`
}
