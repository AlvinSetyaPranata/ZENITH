package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StudyProgramRepository struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func NewStudyProgRepository(db *gorm.DB, log *zap.SugaredLogger) *StudyProgramRepository {
	return &StudyProgramRepository{
		DB:  db,
		Log: log,
	}
}

// Repositories

func (repository *StudyProgramRepository) Create(ctx context.Context, data *entities.StudyProgram) error {
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

func (repository *StudyProgramRepository) GetAll(ctx context.Context, entities *[]entities.StudyProgram) error {
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

func (repository *StudyProgramRepository) GetById(ctx context.Context, entity *entities.StudyProgram, id string) error {
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

func (repository *StudyProgramRepository) Update(ctx context.Context, updatedEntity *entities.StudyProgram, id string) error {
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

func (repository *StudyProgramRepository) Delete(ctx context.Context, entityToDelete *entities.StudyProgram, id string) error {
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
