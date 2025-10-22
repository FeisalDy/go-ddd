package interfaces

import "github.com/FeisalDy/go-ddd/internal/application/dto"

type UserService interface {
	CreateUser(user *dto.CreateUserRequest) (*dto.UserResult, error)
}
