package entities

import "time"

type Class struct {
	Id          uint       `gorm:"column:id;primaryKey"`
	Name        uint       `gorm:"column:name;uniqueIndex:idx_name_grade"`
	Grade       uint       `gorm:"column:grade;uniqueIndex:idx_name_grade"`
	DateCreated *time.Time `gorm:"column:date_created;autoCreatedTime"`
	DateUpdated *time.Time `gorm:"column:date_updated"`
}
