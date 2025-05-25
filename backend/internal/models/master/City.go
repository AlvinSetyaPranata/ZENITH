package master

import "time"

type CityModelRequest struct {
	Name string `json:"name" validate:"true"`
}

type CityModelResponse struct {
	Id          uint      `json:"id" validate:"true"`
	Name        string    `json:"name" validate:"true"`
	DateCreated time.Time `json:"date_created" validate:"true"`
}
