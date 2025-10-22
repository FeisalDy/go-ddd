package postgres

import (
	"fmt"

	"github.com/FeisalDy/go-ddd/internal/domain/entities"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) WithTx(tx *gorm.DB) *RoleRepository {
	return &RoleRepository{db: tx}
}

func (r *RoleRepository) FindAll() ([]entities.Role, error) {
	var roles []entities.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("falied to get roles: %w", err)
	}
	return roles, nil
}

func (r *RoleRepository) FindById(id uint) (*entities.Role, error) {
	var role entities.Role
	if err := r.db.First(&role, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get role by id: %w", err)
	}
	return &role, nil
}

func (r *RoleRepository) FindByName(name string) (*entities.Role, error) {
	var role entities.Role
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, fmt.Errorf("failed to get role by name: %w", err)
	}
	return &role, nil
}
