package master

import (
	"time"

	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/auth"
)

type StudentRequestModel struct {
	Npm          uint   `json:"npm" validate:"true"`
	Name         string `json:"name" validate:"true"`
	Gender       int    `json:"gender_id" validate:"true"`
	Address      string `json:"address" validate:"true"`
	City         int    `json:"city_id" validate:"true"`
	Religion     int    `json:"religion_id" validate:"true"`
	Country      int    `json:"country_id" validate:"true"`
	Province     int    `json:"province_id" validate:"true"`
	Status       int    `json:"status_id" validate:"true"`
	User         int    `json:"user_id" validate:"true"`
	Faculty      int    `json:"faculty_id" validate:"true"`
	StudyProgram int    `json:"study_program_id" validate:"true"`
}

type StudentResponseModel struct {
	Npm          uint                      `json:"npm" validate:"true"`
	Name         string                    `json:"name" validate:"true"`
	Gender       GenderResponseModel       `json:"gender_id" validate:"true"`
	Address      string                    `json:"address" validate:"true"`
	City         CityModelResponse         `json:"city" validate:"true"`
	Religion     ReligionModelResponse     `json:"religion" validate:"true"`
	Country      CountryModelResponse      `json:"country" validate:"true"`
	Province     ProvinceResponseModel     `json:"province" validate:"true"`
	Status       StatusModelResponse       `json:"status" validate:"true"`
	User         auth.UserResponseModel    `json:"user" validate:"true"`
	Faculty      FacultyResponseModel      `json:"faculty" validate:"true"`
	StudyProgram StudyProgramResponseModel `json:"study_program" validate:"true"`
	DateCreated  time.Time                 `json:"date_created" validate:"true"`
	DateUpdated  time.Time                 `json:"date_updated" validate:"true"`
}
