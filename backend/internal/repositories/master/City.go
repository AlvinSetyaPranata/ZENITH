package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
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

func (repository *CityRepository) Create(ctx context.Context, data *entities.City) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Create(&data).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	success = true
	return nil
}

func (repository *CityRepository) GetAll(ctx context.Context, entities *[]entities.City) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Find(entities).Error; err != nil {
		return err
	}

	success = true
	return nil
}

func (repository *CityRepository) GetById(ctx context.Context, entity *entities.City, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := true

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

func (repository *CityRepository) Update(ctx context.Context, updatedEntity *entities.City, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Updates(updatedEntity).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	success = true
	return nil

}

func (repository *CityRepository) Delete(ctx context.Context, entityToDelete *entities.City, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Delete(entityToDelete); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true

	return nil
}
