package master

import (
	"time"
)

type StudyProgramRequestModel struct {
	Name    string `json:"name" validate:"required"`
	Faculty string `json:"faculty" validate:"required"`
}

type StudyProgramResponseModel struct {
	Id          uint                  `json:"id" validate:"true"`
	Name        string                `json:"name" validate:"true"`
	Faculty     *FacultyResponseModel `json:"faculty" validate:"true"`
	DateCreated time.Time             `json:"created_at" validate:"true"`
	DateUpdated time.Time             `json:"updated_at" validate:"true"`
}
