package entities

import "time"

type UserRole struct {
	UserID    uint
	RoleID    uint
	CreatedAt time.Time
}
