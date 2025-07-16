package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{
		DB: db,
	}
}

// Core repository

func (Repository *StudentRepository) Create(ctx context.Context, newStudentEntity *entities.Student) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Create(newStudentEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}

func (Repository *StudentRepository) GetAll(ctx context.Context, studentEntities *[]entities.Student) error {

	err := Repository.DB.Preload("Gender").Preload("City").Preload("Religion").Preload("Province").Preload("Country").Preload("Status").Preload("User").Preload("Faculty").Preload("StudyProgram").Find(studentEntities)

	if err != nil {
		return err.Error
	}

	return nil

}
func (Repository *StudentRepository) GetById(ctx context.Context, studentEntity *entities.Student, id string) error {

	err := Repository.DB.Preload("Gender").Preload("City").Preload("Religion").Preload("Province").Preload("Country").Preload("Status").Preload("User").Preload("User.Role").Preload("Faculty").Preload("StudyProgram").Where("npm = ?", id).Take(studentEntity)

	if err != nil {
		return err.Error
	}

	return nil

}

func (Repository *StudentRepository) GetBySingleQuery(ctx context.Context, studentEntity *entities.Student, query string) error {
	err := Repository.DB.Preload("Gender").Preload("City").Preload("Religion").Preload("Province").Preload("Country").Preload("Status").Preload("User").Preload("User.Role").Preload("Faculty").Preload("StudyProgram").Where(query).Take(studentEntity)

	if err != nil {
		return err.Error
	}

	return nil

}

func (Repository *StudentRepository) Update(ctx context.Context, newStudentEntity *entities.Student, id string) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Where("npm = ?", id).Updates(newStudentEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}

func (Repository *StudentRepository) Delete(ctx context.Context, studentEntity *entities.Student, id string) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Where("npm = ?", id).Delete(studentEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}
