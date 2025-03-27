package domain

import "time"

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResult struct {
	UserId    string     `json:"user_id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	UserType  string     `json:"user_type"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserLoginResponse struct {
	Code        int    `json:"code" example:"200"`
	Message     string `json:"message" example:"successfully login"`
	Accesstoken string
}
