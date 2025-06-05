package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"gorm.io/gorm"
)

type LectureRepository struct {
	DB *gorm.DB
}

func NewLectureRepository(db *gorm.DB) *LectureRepository {
	return &LectureRepository{
		DB: db,
	}
}

// Core repository

func (Repository *LectureRepository) Create(ctx context.Context, newLectureEntity *entities.Lecture) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Create(newLectureEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}

func (Repository *LectureRepository) GetAll(ctx context.Context, lectureEntities *[]entities.Lecture) error {

	err := Repository.DB.Preload("Gender").Preload("City").Preload("Religion").Preload("Province").Preload("Country").Preload("Status").Preload("User").Preload("Faculty").Preload("StudyProgram").Find(lectureEntities)

	if err != nil {
		return err.Error
	}

	return nil

}
func (Repository *LectureRepository) GetById(ctx context.Context, lectureEntity *entities.Lecture, id string) error {

	err := Repository.DB.Preload("Gender").Preload("City").Preload("Religion").Preload("Province").Preload("Country").Preload("Status").Preload("User").Preload("Faculty").Preload("StudyProgram").Where("nidn = ?", id).Take(lectureEntity)

	if err != nil {
		return err.Error
	}

	return nil

}

func (Repository *LectureRepository) Update(ctx context.Context, newLectureEntity *entities.Lecture, id string) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Where("nidn = ?", id).Updates(newLectureEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}

func (Repository *LectureRepository) Delete(ctx context.Context, lectureEntity *entities.Lecture, id string) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Where("nidn = ?", id).Delete(lectureEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}
