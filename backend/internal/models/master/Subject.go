package master

import (
	"time"
)

type SubjectRequestModel struct {
	Name    string `json:"name" validate:"required"`
	Lecture int    `json:"lecture_id" validate:"true"`
}

type SubjectResponseModel struct {
	Id          uint      `json:"id" validate:"true"`
	Name        string    `json:"name" validate:"required"`
	Lecture     int       `json:"lecture_id" validate:"true"`
	DateCreated time.Time `json:"created_at" validate:"true"`
	DateUpdated time.Time `json:"updated_at" validate:"true"`
}
