package repositories

import "gorm.io/gorm"

type PostgresRepository[T any] struct {
	DB *gorm.DB
}
