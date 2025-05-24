package entities

import "time"

type StudyProgram struct {
	Id          uint      `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name;unique"`
	DateCreated time.Time `gorm:"column:date_created;autoCreateTime"`
}
