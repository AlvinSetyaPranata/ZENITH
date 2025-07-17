package auth

import "time"

type Token struct {
	ID        int `gorm:"primaryKey"`
	UserID    uint
	User      User      `gorm:"foreignKey:UserID;"`
	Revoked   bool      `gorm:"column:revoked;default:false"`
	ExpiredAt int64     `gorm:"column:expired_at"`
	CreatedAt time.Time `gorm:"column:date_created;autoCreateTime"`
}
