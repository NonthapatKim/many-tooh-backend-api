package function

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func CustomEmailValidator(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return emailRegex.MatchString(email)
}

func CustomPasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecialChar := regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(password)
	hasLowerCase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpperCase := regexp.MustCompile(`[A-Z]`).MatchString(password)

	return hasNumber && hasSpecialChar && hasLowerCase && hasUpperCase
}
