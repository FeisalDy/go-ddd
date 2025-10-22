package interfaces

import "github.com/FeisalDy/go-ddd/internal/application/dto"

type RoleService interface {
	FindAll() ([]dto.RoleResult, error)
	FindByID(id uint) (*dto.RoleResult, error)
	FindByName(name string) (*dto.RoleResult, error)
}
