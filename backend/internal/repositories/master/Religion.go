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

	if err := repository.DB.Create(entity).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	success = true
	return nil

}

func (repository *ReligionRepository) GetAll(ctx context.Context, entity *[]entities.Religion) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Find(entity).Error; err != nil {
		return err
	}

	success = true
	return nil
}

func (repository *ReligionRepository) GetByID(ctx context.Context, entity *entities.Religion, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Take(entity).Error; err != nil {
		return err
	}

	success = true
	return nil
}

func (repository *ReligionRepository) Update(ctx context.Context, newEntity *entities.Religion, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Updates(newEntity).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	success = true
	return nil
}

func (repository *ReligionRepository) Delete(ctx context.Context, entity *entities.Religion, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Delete(entity).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	success = true
	return nil
}
