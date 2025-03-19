package domain

type ValidationError struct {
	Errors ValidationErrorResponse `json:"errors"`
}

type ValidationErrorResponse struct {
	Fullname  *string `json:"fullname"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	Incorrect *string `json:"incorrect"`
	AuthError *string `json:"auth_error"`
	UserError *string `json:"user_error"`
}

func (v ValidationError) Error() string {
	return "validation failed"
}
