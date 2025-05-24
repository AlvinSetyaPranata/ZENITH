package repositories

import "gorm.io/gorm"

type BaseRepository[T any] struct {
	DB *gorm.DB
}
