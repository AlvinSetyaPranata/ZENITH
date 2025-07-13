package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BatchRepository struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func NewBatchRepository(db *gorm.DB, log *zap.SugaredLogger) *BatchRepository {
	return &BatchRepository{
		DB:  db,
		Log: log,
	}
}

// Core repositories

func (Repository *BatchRepository) Create(ctx context.Context, batchEntity *entities.Batch) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Create(batchEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}

func (Repository *BatchRepository) GetAll(ctx context.Context, batchEntities *[]entities.Batch) error {
	if err := Repository.DB.Find(batchEntities); err != nil {
		return err.Error
	}

	return nil

}

func (repository *BatchRepository) GetById(ctx context.Context, batchEntity *entities.Batch, id string) error {

	if err := repository.DB.Where("id = ?", id).Take(batchEntity); err != nil {
		return err.Error
	}

	return nil

}

func (repository *BatchRepository) Update(ctx context.Context, newBatchEntity *entities.Batch, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Updates(newBatchEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}

func (repository *BatchRepository) Delete(ctx context.Context, batchEntity *entities.Batch, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Delete(batchEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}
