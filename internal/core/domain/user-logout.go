package domain

type UserLogoutResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
