package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FacultyRepository struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func NewFacultyRepository(db *gorm.DB, log *zap.SugaredLogger) *FacultyRepository {
	return &FacultyRepository{
		DB:  db,
		Log: log,
	}
}

// Core repositories

func (repository *FacultyRepository) Create(ctx context.Context, faculty *entities.Faculty) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Create(faculty); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}

func (repository *FacultyRepository) GetAll(ctx context.Context, facultyEntity *[]entities.Faculty) error {

	if err := repository.DB.Find(facultyEntity); err != nil {
		return err.Error
	}

	return nil
}

func (repository *FacultyRepository) GetById(ctx context.Context, facultyEntity *entities.Faculty, id string) error {
	if err := repository.DB.Where("id = ?", id).First(facultyEntity); err != nil {
		return err.Error
	}

	return nil
}

func (repository *FacultyRepository) Update(ctx context.Context, facultyEntity *entities.Faculty) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Save(facultyEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}
func (repository *FacultyRepository) Delete(ctx context.Context, facultyEntity *entities.Faculty, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Delete(facultyEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}
