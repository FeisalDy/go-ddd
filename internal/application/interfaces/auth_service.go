package interfaces

import "github.com/FeisalDy/go-ddd/internal/application/dto"

type AuthService interface {
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}
