package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GenderRepository struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func NewGenderRepository(db *gorm.DB, log *zap.SugaredLogger) *GenderRepository {
	return &GenderRepository{
		DB:  db,
		Log: log,
	}
}

// Core gender repositories

func (repository *GenderRepository) Create(ctx context.Context, genderEntity *entities.Gender) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Create(genderEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil

}

func (repository *GenderRepository) GetAll(ctx context.Context, genderEntities *[]entities.Gender) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Find(genderEntities); err != nil {
		return err.Error
	}

	success = true
	return nil
}

func (repository *GenderRepository) GetById(ctx context.Context, genderEntity *entities.Gender, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Take(genderEntity); err != nil {
		return err.Error
	}

	success = true
	return nil
}

func (repository *GenderRepository) Update(ctx context.Context, updatedGenderEntity *entities.Gender, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Updates(updatedGenderEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil

}

func (repository *GenderRepository) Delete(ctx context.Context, genderEntity *entities.Gender, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Delete(genderEntity); err != nil {
		return err.Error
	}

	success = true
	return nil
}
