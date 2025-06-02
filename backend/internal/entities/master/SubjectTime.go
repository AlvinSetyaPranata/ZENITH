package entities

import "time"

type SubjectTime struct {
	Id          uint      `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name;unique"`
	TimeStart   time.Time `gorm:"column:time_start"`
	TimeEnd     time.Time `gorm:"column:time_end"`
	DateCreated time.Time `gorm:"column:date_created;autoCreateTime"`
	DateUpdated time.Time `gorm:"column:date_updated;autoCreateTime"`
}
