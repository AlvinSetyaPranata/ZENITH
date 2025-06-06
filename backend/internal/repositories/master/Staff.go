package master

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/master"
	"gorm.io/gorm"
)

type StaffRepository struct {
	DB *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
	return &StaffRepository{
		DB: db,
	}
}

// Core repository

func (Repository *StaffRepository) Create(ctx context.Context, newStaffEntity *entities.Staff) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Create(newStaffEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}

func (Repository *StaffRepository) GetAll(ctx context.Context, staffEntities *[]entities.Staff) error {

	err := Repository.DB.Preload("Gender").Preload("City").Preload("Religion").Preload("Province").Preload("Country").Preload("Status").Preload("User").Preload("Faculty").Preload("StudyProgram").Find(staffEntities)

	if err != nil {
		return err.Error
	}

	return nil

}
func (Repository *StaffRepository) GetById(ctx context.Context, staffEntity *entities.Staff, id string) error {

	err := Repository.DB.Preload("Gender").Preload("City").Preload("Religion").Preload("Province").Preload("Country").Preload("Status").Preload("User").Preload("Faculty").Preload("StudyProgram").Where("nidn = ?", id).Take(staffEntity)

	if err != nil {
		return err.Error
	}

	return nil

}

func (Repository *StaffRepository) Update(ctx context.Context, newStaffEntity *entities.Staff, id string) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Where("nidn = ?", id).Updates(newStaffEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}

func (Repository *StaffRepository) Delete(ctx context.Context, staffEntity *entities.Staff, id string) error {
	tx := Repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := Repository.DB.Where("nidn = ?", id).Delete(staffEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}
