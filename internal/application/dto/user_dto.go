package dto

import "time"

type CreateUserRequest struct {
	Username  string
	Email     string
	Password  string
	AvatarUrl string
	Bio       string
	Status    string
}

type UserResult struct {
	ID        uint
	Username  string
	Email     string
	AvatarUrl string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
