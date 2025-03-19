package domain

import "time"

type CreateStaffUserRequest struct {
	UserId    string
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateStaffUserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
