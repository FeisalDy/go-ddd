package postgres

import (
	"fmt"

	"github.com/FeisalDy/go-ddd/internal/domain/entities"
	"github.com/FeisalDy/go-ddd/internal/domain/repositories"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) WithTx(tx *gorm.DB) repositories.UserRepository {
	return &UserRepository{db: tx}
}

func (r *UserRepository) Create(user *entities.ValidatedUser) (*entities.User, error) {
	u := user.User
	if err := r.db.Create(&u).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &user.User, nil
}

func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&entities.User{}).
		Where("email = ?", email).
		Count(&count).Error; err != nil {
		return false, fmt.Errorf("failed to check existing user: %w", err)
	}
	return count > 0, nil
}
