package auth

import "time"

type User struct {
	Id          uint   `gorm:"column:id;primaryKey"`
	Email       string `gorm:"column:email;unique"`
	Password    string `gorm:"column:password"`
	RoleId      int
	Role        Role      `gorm:"foreignKey:RoleId;"`
	DateCreated time.Time `gorm:"column:date_created;autoCreateTime"`
	DateUpdated time.Time `gorm:"column:date_updated;"`
}
