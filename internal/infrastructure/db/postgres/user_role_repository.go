package postgres

import (
	"errors"
	"fmt"

	"github.com/FeisalDy/go-ddd/internal/domain/entities"
	"github.com/FeisalDy/go-ddd/internal/domain/repositories"
	"gorm.io/gorm"
)

type UserRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) repositories.UserRoleRepository {
	return &UserRoleRepository{db: db}
}

func (r *UserRoleRepository) Assign(userID uint, roleID uint) error {
	ur := entities.UserRole{UserID: userID, RoleID: roleID}
	if err := r.db.Where("user_id = ? AND role_id = ?", userID, roleID).FirstOrCreate(&ur).Error; err != nil {
		return fmt.Errorf("failed to assign role to user: %w", err)
	}
	return nil
}

func (r *UserRoleRepository) Remove(userID uint, roleID uint) error {
	res := r.db.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&entities.UserRole{})
	if res.Error != nil {
		return fmt.Errorf("failed to remove role from user: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return errors.New("assignment not found")
	}
	return nil
}

func (r *UserRoleRepository) GetRolesForUser(userID uint) ([]entities.Role, error) {
	var roles []entities.Role
	if err := r.db.Table("roles").
		Joins("JOIN user_roles ur ON ur.role_id = roles.id").
		Where("ur.user_id = ?", userID).
		Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("failed to get roles for user: %w", err)
	}
	return roles, nil
}

func (r *UserRoleRepository) GetUsersForRole(roleID uint) ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Table("users").
		Joins("JOIN user_roles ur ON ur.user_id = users.id").
		Where("ur.role_id = ?", roleID).
		Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get users for role: %w", err)
	}
	return users, nil
}

func (r *UserRoleRepository) HasRole(userID uint, roleID uint) (bool, error) {
	var count int64
	if err := r.db.Model(&entities.UserRole{}).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Count(&count).Error; err != nil {
		return false, fmt.Errorf("failed to check role assignment: %w", err)
	}
	return count > 0, nil
}
