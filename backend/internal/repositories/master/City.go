package master

import (
	"context"
	"time"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/models/master"
	"github.com/AlvinSetyaPranata/ZENITH/backend/internal/repositories"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CityRepository struct {
	repositories.BaseRepository[entities.City]
	Log *zap.SugaredLogger
}

func NewCityRepository(db *gorm.DB, log *zap.SugaredLogger) *CityRepository {
	return &CityRepository{
		BaseRepository: repositories.BaseRepository[entities.City]{
			DB: db,
		},
		Log: log,
	}
}

// Repositories

func (r *CityRepository) Create(ctx context.Context, data master.CityModelRequest) (*entities.City, error) {
	tx := r.DB.WithContext(ctx).Begin()
	success := true

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	cityEntity := &entities.City{
		Name:        data.Name,
		DateCreated: time.Now(),
	}

	if err := tx.Create(cityEntity).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	success = true
	return cityEntity, nil
}

func (r *CityRepository) GetAll(ctx context.Context, entity *[]entities.City) error {
	tx := r.DB.WithContext(ctx)

	defer tx.Rollback()

	return r.DB.Find(entity).Error
}
