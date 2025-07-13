package master

import "time"

type BatchRequestModel struct {
	Name        string    `json:"name" validate:"true"`
	StartPeriod time.Time `json:"start_period" validate:"true"`
	EndPeriod   time.Time `json:"end_period" validate:"true"`
}

type BatchResponseModel struct {
	Id          uint      `json:"id" validate:"true"`
	Name        string    `json:"name" validate:"true"`
	StartPeriod time.Time `json:"start_period" validate:"true"`
	EndPeriod   time.Time `json:"end_period" validate:"true"`
	DateCreated time.Time `json:"date_created" validate:"true"`
	DateUpdated time.Time `json:"date_updated" validate:"true"`
}
