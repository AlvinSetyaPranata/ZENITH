package auth

import (
	"context"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/auth"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func NewPermissionRepository(db *gorm.DB, log *zap.SugaredLogger) *PermissionRepository {
	return &PermissionRepository{
		DB:  db,
		Log: log,
	}
}

// Core Repositories

func (repository *PermissionRepository) Create(ctx context.Context, permisisonEntity *entities.Permission) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Create(permisisonEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil

}

func (repository *PermissionRepository) GetAll(ctx context.Context, permissionEntities *[]entities.Permission) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Find(permissionEntities); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}

func (repository *PermissionRepository) GetById(ctx context.Context, permissionEntity *entities.Permission, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Take(permissionEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil

}

func (repository *PermissionRepository) Update(ctx context.Context, newPermissionEntity *entities.Permission, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Updates(newPermissionEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}

func (repository *PermissionRepository) Delete(ctx context.Context, permissionEntity *entities.Permission, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Delete(permissionEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}
