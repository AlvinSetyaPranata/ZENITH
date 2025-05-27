package master

import (
	// entities "github.com/AlvinSetyaPranata/ZENITH/backend/pkg/entities/master"
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StatusRepository struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func NewStatusRepository(db *gorm.DB, log *zap.SugaredLogger) *StatusRepository {
	return &StatusRepository{
		DB:  db,
		Log: log,
	}
}

// Core Repositories

func (repository *StatusRepository) Create(ctx context.Context, entity *entities.Status) error {
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

func (repository *StatusRepository) GetAll(ctx context.Context, entity *[]entities.Status) error {
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

func (repository *StatusRepository) GetByID(ctx context.Context, entity *entities.Status, id string) error {
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

func (repository *StatusRepository) Update(ctx context.Context, newEntity *entities.Status, id string) error {
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

func (repository *StatusRepository) Delete(ctx context.Context, entity *entities.Status, id string) error {
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
