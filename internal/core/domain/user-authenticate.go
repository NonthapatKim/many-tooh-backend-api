package domain

type UserAuthenticateRequest struct {
	AccessToken string `validate:"required"`
}

type UserAuthenticateResponse struct {
	UserId   string `json:"user_id"`
	Role     string `json:"role"`
	Username string `json:"username"`
}
