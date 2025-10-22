package dto

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string     `json:"token"`
	User  UserResult `json:"user"`
}
