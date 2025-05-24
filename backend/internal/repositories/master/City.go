package master

import (
	"context"
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories"
	"gorm.io/gorm"
)

type CityRepository struct {
	repositories.BaseRepository[entities.City]
}

func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{}
}

// Repositories

func (r *CityRepository) Create(ctx context.Context, data master.CityModelRequest) bool {
	tx := r.DB.WithContext(ctx)

	defer tx.Rollback()

	cityEntity := &entities.City{
		Name:        data.Name,
		DateCreated: time.Now(),
	}

	if err := r.DB.Create(cityEntity).Error; err != nil {

		return false
	}

	if err := tx.Commit().Error; err != nil {

		return false
	}

	return true

}

func (r *CityRepository) GetAll(ctx context.Context, entity *[]entities.City) error {
	tx := r.DB.WithContext(ctx)

	defer tx.Rollback()

	return r.DB.Find(entity).Error
}
