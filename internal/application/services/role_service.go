package services

import (
	"github.com/FeisalDy/go-ddd/internal/application/dto"
	"github.com/FeisalDy/go-ddd/internal/domain/repositories"
)

type RoleService struct {
	repo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) FindAll() ([]dto.RoleResult, error) {
	roles, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var result []dto.RoleResult
	for _, role := range roles {
		result = append(result, dto.RoleResult{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
		})
	}
	return result, nil
}

func (s *RoleService) FindByID(id uint) (*dto.RoleResult, error) {
	role, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return &dto.RoleResult{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}

func (s *RoleService) FindByName(name string) (*dto.RoleResult, error) {
	role, err := s.repo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return &dto.RoleResult{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}
