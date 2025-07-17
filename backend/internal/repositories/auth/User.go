package auth

import (
	"context"
	"log"

	entities "github.com/AlvinSetyaPranata/ZENITH/backend/internal/entities/auth"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB  *gorm.DB
	Log *zap.SugaredLogger
}

func NewUserRepository(db *gorm.DB, log *zap.SugaredLogger) *UserRepository {
	return &UserRepository{
		DB:  db,
		Log: log,
	}
}

// Core Repositories

func (repository *UserRepository) Create(ctx context.Context, userEntity *entities.User) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Create(userEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true

	return nil

}

func (repository *UserRepository) GetAllUser(ctx context.Context, userEntites *[]entities.User) error {

	if err := repository.DB.Preload("Role").Find(userEntites); err != nil {
		return err.Error
	}

	return nil

}

func (repository *UserRepository) GetUserById(ctx context.Context, userEntity *entities.User, id string) error {

	if err := repository.DB.Preload("Role").Where("id = ?", id).Take(userEntity); err != nil {
		return err.Error
	}

	return nil
}

func (repository *UserRepository) GetUserByEmail(ctx context.Context, userEntity *entities.User, email string) error {

	if err := repository.DB.Preload("Role").Preload("Role.Permissions").Where("email = ?", email).Take(userEntity); err != nil {
		return err.Error
	}

	return nil

}

func (repository *UserRepository) GetUserByRole(ctx context.Context, userEntity *entities.User, role string) error {

	if err := repository.DB.Preload("Role").Where("role_id = ?", role).Take(userEntity); err != nil {
		return err.Error
	}

	return nil

}

func (repository *UserRepository) Update(ctx context.Context, newUserEntity *entities.User, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Model(newUserEntity).Association("Role").Replace(newUserEntity.Role); err != nil {
		return err
	}

	if err := repository.DB.Where("id = ?", id).Updates(newUserEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	success = true

	return nil
}

func (repository *UserRepository) Delete(ctx context.Context, userEntity *entities.User, id string) error {
	tx := repository.DB.WithContext(ctx).Begin()
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Where("id = ?", id).Delete(userEntity); err != nil {
		return err.Error
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}

	return nil

}

func (repository *UserRepository) StoreToken(ctx context.Context, tokenEntity *entities.Token) error {
	tx := repository.DB.WithContext(ctx)
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Create(tokenEntity); err != nil {
		return err.Error
	}

	success = true

	return nil
}

func (repository *UserRepository) InvalidateToken(ctx context.Context, tokenEntity *entities.Token, userId uint) {
	tx := repository.DB.WithContext(ctx)
	success := false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	if err := repository.DB.Model(&entities.Token{}).Where("user_id = ?", userId).Where("revoked = ?", false).Order("date_created DESC").Updates(tokenEntity).Error; err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}

	success = true
}

func (repository *UserRepository) GetToken(ctx context.Context, userId uint) error {
	var token entities.Token
	err := repository.DB.WithContext(ctx).Where("user_id = ?", userId).Where("revoked = ?", false).Order("date_created DESC").First(&token).Error
	return err
}
