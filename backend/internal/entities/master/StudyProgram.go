package entities

import "time"

type StudyProgram struct {
	Id          uint `gorm:"column:id;primaryKey"`
	FacultyId   int
	Faculty     *Faculty  `gorm:"foreignKey:FacultyId"`
	Name        string    `gorm:"column:name;unique"`
	DateCreated time.Time `gorm:"column:date_created;autoCreateTime"`
	DateUpdated time.Time `gorm:"column:date_updated"`
}
