package repositories

import "github.com/FeisalDy/go-ddd/internal/domain/entities"

type RoleRepository interface {
	FindAll() ([]entities.Role, error)
	FindById(id uint) (*entities.Role, error)
	FindByName(name string) (*entities.Role, error)
}
