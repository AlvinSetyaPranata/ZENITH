package master

import "time"

type StudyProgramRequestModel struct {
	Name string `json:"name" validate:"required"`
}

type StudyProgramResponseModel struct {
	Id          uint      `json:"id" validate:"true"`
	Name        string    `json:"name" validate:"true"`
	DateCreated time.Time `json:"created_at" validate:"true"`
}
