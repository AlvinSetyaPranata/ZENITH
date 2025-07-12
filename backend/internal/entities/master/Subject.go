package entities

import "time"

type Subject struct {
	Id            uint         `gorm:"column:id;primaryKey"`
	Name          string       `gorm:"column:name;not null;uniqueIndex:idx_name_studytime"`
	LectureId     int          `gorm:"column:lecture_id;not null"`
	SubjectTimeId int          `gorm:"column:subject_time_id;not null;uniqueIndex:idx_name_studytime"`
	Lecture       *Lecture     `gorm:"foreignKey:LectureId"`
	SubjectTime   *SubjectTime `gorm:"foreignKey:SubjectTimeId"`
	DateCreated   time.Time    `gorm:"column:date_created;autoCreateTime"`
	DateUpdated   time.Time    `gorm:"column:date_updated;autoUpdateTime"`
}
