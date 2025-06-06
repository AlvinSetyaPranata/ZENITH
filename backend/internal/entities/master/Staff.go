package entities

import (
	"time"

	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/auth"
)

type Staff struct {
	Npwp           uint   `gorm:"column:npwp;primaryKey"`
	Name           string `gorm:"column:name;unique"`
	Address        string `gorm:"column:address"`
	GenderId       int
	ReligionId     int
	CountryId      int
	ProvinceId     int
	CityId         int
	StatusId       int
	UserId         int
	FacultyId      int
	StudyProgramId int
	Gender         *Gender       `gorm:"foreignKey:GenderId"`
	Religion       *Religion     `gorm:"foreignKey:ReligionId"`
	Country        *Country      `gorm:"foreignKey:CountryId"`
	Province       *Province     `gorm:"foreignKey:ProvinceId"`
	City           *City         `gorm:"foreignKey:CityId"`
	Status         *Status       `gorm:"foreignKey:StatusId"`
	User           *auth.User    `gorm:"foreignKey:UserId"`
	Faculty        *Faculty      `gorm:"foreignKey:FacultyId"`
	StudyProgram   *StudyProgram `gorm:"foreignKey:StudyProgramId"`
	DateCreated    time.Time     `gorm:"column:date_created;autoCreateTime"`
	DateUpdated    time.Time     `gorm:"column:date_created"`
}
