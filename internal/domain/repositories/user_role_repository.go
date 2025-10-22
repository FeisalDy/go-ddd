package repositories

import "github.com/FeisalDy/go-ddd/internal/domain/entities"

type UserRoleRepository interface {
	Assign(userID uint, roleID uint) error
	Remove(userID uint, roleID uint) error
	GetRolesForUser(userID uint) ([]entities.Role, error)
	GetUsersForRole(roleID uint) ([]entities.User, error)
	HasRole(userID uint, roleID uint) (bool, error)
}
