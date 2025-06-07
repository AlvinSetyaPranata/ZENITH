package auth

import (
	"context"
	"fmt"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/auth"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func NewRoleRepository(db *gorm.DB, log *zap.SugaredLogger) *RoleRepository {
	return &RoleRepository{
		DB:  db,
		Log: log,
	}
}

// Core Repositories

func (repository *RoleRepository) Create(ctx context.Context, roleEntity *entities.Role) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Create(roleEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil

}

func (repository *RoleRepository) GetAllRolesData(ctx context.Context, roleEntities *[]entities.Role) error {
	if err := repository.DB.Preload("Permissions").Find(roleEntities); err != nil {
		return err.Error
	}

	return nil
}

func (repository *RoleRepository) GetRoleById(ctx context.Context, roleEntity *entities.Role, id string) error {
	if err := repository.DB.Preload("Permissions").Where("id = ?", id).Take(roleEntity); err != nil {
		return err.Error
	}

	return nil
}

func (repository *RoleRepository) Update(ctx context.Context, newRoleEntity *entities.Role, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Updates(newRoleEntity); err != nil {
		return err.Error
	}

	// Update association
	if err := repository.DB.Association("Permissions").Replace(newRoleEntity.Permissions); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true

	return nil
}

func (repository *RoleRepository) Delete(ctx context.Context, roleEntity *entities.Role, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Model(roleEntity).Association("Permissions").Clear(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err := repository.DB.Where("id = ?", id).Delete(roleEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true
	return nil
}
