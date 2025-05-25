package master

import (
	// entities "github.com/AlvinSetyaPranata/ZENITH/backend/pkg/entities/master"
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ReligionRepository struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func NewReligionRepository(db *gorm.DB, log *zap.SugaredLogger) *ReligionRepository {
	return &ReligionRepository{
		DB:  db,
		Log: log,
	}
}

// Core Repositories

func (repository *ReligionRepository) Create(ctx context.Context, entity *entities.Religion) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	success = true

	if err := tx.Create(entity).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}
