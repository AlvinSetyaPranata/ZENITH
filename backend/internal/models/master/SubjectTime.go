package master

import "time"

type SubjectTimeRequestModel struct {
	Name      string    `json:"name" validate:"true"`
	TimeStart time.Time `json:"time_start" validate:"true"`
	TimeEnd   time.Time `json:"time_end" validate:"true"`
}

type SubjectTimeResponseModel struct {
	Id          uint      `json:"id" validate:"true"`
	Name        string    `json:"name" validate:"true"`
	TimeStart   time.Time `json:"time_start" validate:"true"`
	TimeEnd     time.Time `json:"time_end" validate:"true"`
	DateCreated time.Time `json:"date_created" validate:"true"`
	DateUpdated time.Time `json:"date_updated" validate:"true"`
}
