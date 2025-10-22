package services

import (
	"errors"
	"fmt"

	"github.com/FeisalDy/go-ddd/internal/application/dto"
	"github.com/FeisalDy/go-ddd/internal/application/interfaces"
	"github.com/FeisalDy/go-ddd/internal/domain/entities"
	"github.com/FeisalDy/go-ddd/internal/domain/repositories"
	"github.com/FeisalDy/go-ddd/internal/domain/services"
)

type UserService struct {
	repo   repositories.UserRepository
	hasher services.HashService
}

func NewUserService(repo repositories.UserRepository, hasher services.HashService) interfaces.UserService {
	return &UserService{repo: repo, hasher: hasher}
}

func (s *UserService) CreateUser(user *dto.CreateUserRequest) (*dto.UserResult, error) {
	isExists, err := s.repo.ExistsByEmail(user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists: %w", err)
	}

	if isExists {
		return nil, errors.New("user with this email already exists")
	}
	hashedPassword, err := s.hasher.Hash(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	newUser := entities.NewUser(user.Username, user.Email, hashedPassword, user.AvatarUrl, user.Bio)

	validatedUser, err := entities.NewValidatedUser(newUser)
	if err != nil {
		return nil, fmt.Errorf("invalid user data: %w", err)
	}

	createdUser, err := s.repo.Create(validatedUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &dto.UserResult{
		ID:        createdUser.ID,
		Email:     createdUser.Email,
		Username:  createdUser.Username,
		AvatarUrl: createdUser.AvatarUrl,
		Bio:       createdUser.Bio,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}
