package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	DB *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{
		DB: db,
	}
}

// Core repositories

func (Repository *SubjectRepository) Create(ctx context.Context, subjectEntity *entities.Subject) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Create(subjectEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil

}

func (Repository *SubjectRepository) GetAll(ctx context.Context, subjectEntities *[]entities.Subject) error {

	if err := Repository.DB.Preload("Lecture").Preload("SubjectTime").Find(subjectEntities); err != nil {
		return err.Error
	}

	return nil

}

func (Repository *SubjectRepository) GetById(ctx context.Context, subjectEntity *entities.Subject, id string) error {

	if err := Repository.DB.Preload("Lecture").Preload("SubjectTime").Where("id = ?", id).Take(subjectEntity); err != nil {
		return err.Error
	}

	return nil

}

func (Repository *SubjectRepository) Update(ctx context.Context, newSubjectEntity *entities.Subject, id string) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Where("id = ?", id).Updates(newSubjectEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil

}

func (Repository *SubjectRepository) Delete(ctx context.Context, subjectEntityToDelete *entities.Subject, id string) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Where("id = ?", id).Delete(subjectEntityToDelete); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil

}
