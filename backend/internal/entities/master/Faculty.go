package entities

import "time"

type Faculty struct {
	Id          uint      `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name;unique"`
	DateCreated time.Time `gorm:"column:date_created;autoCreateTime"`
}
