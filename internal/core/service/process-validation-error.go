package service

import (
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"
	"github.com/go-playground/validator/v10"
)

func ProcessValidationError(err error) domain.ValidationErrorResponse {
	var validationErrors domain.ValidationErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Field() {
		case "Fullname":
			message := "กรุณากรอกชื่อ"
			validationErrors.Fullname = &message
		case "Email":
			if err.Tag() == "customEmail" {
				message := "รูปแบบอีเมลไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง"
				validationErrors.Email = &message
			} else {
				message := "กรุณากรอกอีเมล"
				validationErrors.Email = &message
			}
		case "Password":
			if err.Tag() == "customPassword" {
				message := "รหัสผ่านต้องมีความยาวอย่างน้อย 8 ตัวอักษร และต้องมีตัวเลขและอักขระพิเศษอย่างน้อยหนึ่งตัว"
				validationErrors.Password = &message
			} else {
				message := "กรุณากรอกรหัสผ่าน"
				validationErrors.Password = &message
			}
		}
	}
	return validationErrors
}
