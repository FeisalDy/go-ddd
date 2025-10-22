package interfaces

import "github.com/FeisalDy/go-ddd/internal/application/dto"

type UserRoleService interface {
	AssignRoles(userID uint, roleIDs []uint) ([]dto.RoleResult, error)
	RemoveRole(userID uint, roleID uint) ([]dto.RoleResult, error)
	ListUserRoles(userID uint) ([]dto.RoleResult, error)
}
