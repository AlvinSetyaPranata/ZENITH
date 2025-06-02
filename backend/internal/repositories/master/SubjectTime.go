package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SubjectTimeRepository struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func NewSubjectTimeRepository(db *gorm.DB, log *zap.SugaredLogger) *SubjectTimeRepository {
	return &SubjectTimeRepository{
		DB:  db,
		Log: log,
	}
}

// Core repositories

func (Repository *SubjectTimeRepository) Create(ctx context.Context, subjectTimeEntity *entities.SubjectTime) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Create(subjectTimeEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}

func (Repository *SubjectTimeRepository) GetAll(ctx context.Context, subjectTimeEntities *[]entities.SubjectTime) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Find(subjectTimeEntities); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true

	return nil

}

func (repository *SubjectTimeRepository) GetById(ctx context.Context, subjectTimeEntity *entities.SubjectTime, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Take(subjectTimeEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil

}

func (repository *SubjectTimeRepository) Update(ctx context.Context, newSubjectTimeEntity *entities.SubjectTime, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Updates(newSubjectTimeEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}

func (repository *SubjectTimeRepository) Delete(ctx context.Context, subjectTimeEntity *entities.SubjectTime, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Delete(subjectTimeEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}
