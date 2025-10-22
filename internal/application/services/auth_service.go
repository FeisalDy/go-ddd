package services

import (
	"errors"
	"fmt"

	"github.com/FeisalDy/go-ddd/internal/application/dto"
	"github.com/FeisalDy/go-ddd/internal/application/interfaces"
	"github.com/FeisalDy/go-ddd/internal/domain/repositories"
	"github.com/FeisalDy/go-ddd/internal/domain/services"
)

type AuthService struct {
	userRepo     repositories.UserRepository
	hasher       services.HashService
	tokenService services.JWTService
}

func NewAuthService(
	userRepo repositories.UserRepository,
	hasher services.HashService,
	tokenService services.JWTService,
) interfaces.AuthService {
	return &AuthService{
		userRepo:     userRepo,
		hasher:       hasher,
		tokenService: tokenService,
	}
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare password
	if err := s.hasher.Compare(user.Password, req.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate token
	token, err := s.tokenService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResult{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			AvatarUrl: user.AvatarUrl,
			Bio:       user.Bio,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}
