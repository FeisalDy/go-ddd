package repositories

import "github.com/FeisalDy/go-ddd/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.ValidatedUser) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	ExistsByEmail(email string) (bool, error)
	FindByID(id uint) (*entities.User, error)
}
