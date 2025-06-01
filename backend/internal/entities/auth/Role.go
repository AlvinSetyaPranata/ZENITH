package auth

import (
	"time"
)

type Role struct {
	Id          uint         `gorm:"column:id;primaryKey"`
	Name        string       `gorm:"column:name;unique"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
	DateCreated time.Time    `gorm:"column:date_created;autoCreateTime"`
}
