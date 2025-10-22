package entities

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	AvatarUrl string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username, email, password, avatarUrl, bio string) *User {
	return &User{
		Username:  username,
		Email:     email,
		Password:  password,
		AvatarUrl: avatarUrl,
		Bio:       bio,
	}

}

func (u *User) validate() error {
	if err := validateEmail(u.Email); err != nil {
		return err
	}
	if err := validateUrl(u.AvatarUrl); err != nil {
		return err
	}
	return nil
}

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	urlRegex   = regexp.MustCompile(`[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
)

func validateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}

	if len(email) > 320 {
		return errors.New("email too long")
	}

	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

func validateUrl(url string) error {
	if url == "" {
		return nil
	}

	if !urlRegex.MatchString(url) {
		return errors.New("invalid url format")
	}

	return nil
}
