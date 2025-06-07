package entities

import "time"

type Subject struct {
	Id          uint   `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:name;unique"`
	LectureId   int
	Lecture     *Lecture  `gorm:"foreignKey:LectureId"`
	DateCreated time.Time `gorm:"column:date_created;autoCreatedTime"`
	DateUpdated time.Time `gorm:"column:date_updated"`
}
