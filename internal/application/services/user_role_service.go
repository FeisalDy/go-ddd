package services

import (
	"fmt"

	"github.com/FeisalDy/go-ddd/internal/application/dto"
	appif "github.com/FeisalDy/go-ddd/internal/application/interfaces"
	"github.com/FeisalDy/go-ddd/internal/domain/repositories"
)

type UserRoleService struct {
	userRepo     repositories.UserRepository
	roleRepo     repositories.RoleRepository
	userRoleRepo repositories.UserRoleRepository
}

func NewUserRoleService(
	userRepo repositories.UserRepository,
	roleRepo repositories.RoleRepository,
	userRoleRepo repositories.UserRoleRepository,
) appif.UserRoleService {
	return &UserRoleService{
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		userRoleRepo: userRoleRepo,
	}
}

func (s *UserRoleService) AssignRoles(userID uint, roleIDs []uint) ([]dto.RoleResult, error) {
	if _, err := s.userRepo.FindByID(userID); err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	for _, rid := range roleIDs {
		if _, err := s.roleRepo.FindById(rid); err != nil {
			return nil, fmt.Errorf("role %d not found: %w", rid, err)
		}
		if err := s.userRoleRepo.Assign(userID, rid); err != nil {
			return nil, fmt.Errorf("failed to assign role %d: %w", rid, err)
		}
	}
	return s.ListUserRoles(userID)
}

func (s *UserRoleService) RemoveRole(userID uint, roleID uint) ([]dto.RoleResult, error) {
	if _, err := s.userRepo.FindByID(userID); err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	if _, err := s.roleRepo.FindById(roleID); err != nil {
		return nil, fmt.Errorf("role not found: %w", err)
	}
	if err := s.userRoleRepo.Remove(userID, roleID); err != nil {
		return nil, err
	}
	return s.ListUserRoles(userID)
}

func (s *UserRoleService) ListUserRoles(userID uint) ([]dto.RoleResult, error) {
	if _, err := s.userRepo.FindByID(userID); err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	roles, err := s.userRoleRepo.GetRolesForUser(userID)
	if err != nil {
		return nil, err
	}
	res := make([]dto.RoleResult, 0, len(roles))
	for _, r := range roles {
		res = append(res, dto.RoleResult{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			CreatedAt:   r.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
		})
	}
	return res, nil
}
