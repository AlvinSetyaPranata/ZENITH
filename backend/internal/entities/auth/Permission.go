package auth

import "time"

type Permission struct {
	Id          string    `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name;unique"`
	DateCreated time.Time `gorm:"column:date_created;autoCreateTime"`
	DateUpdated time.Time `gorm:"column:date_updated"`
}
