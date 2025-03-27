package service

import (
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/function"
	"gorm.io/gorm"
)

func (s *service) UserLogin(req domain.UserLoginRequest) (domain.UserLoginResponse, error) {
	var validationErrors domain.ValidationErrorResponse
	var response domain.UserLoginResponse
	var message string

	loginResult, err := s.repo.UserLogin(req)
	if err != nil {
		if err.Error() == "incorrect password" {
			message = "บัญชีผู้ใช้งานหรือรหัสผ่านไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง"
			validationErrors.Incorrect = &message
			return response, domain.ValidationError{Errors: validationErrors}
		}

		if err == gorm.ErrRecordNotFound {
			message = "บัญชีผู้ใช้งานหรือรหัสผ่านไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง"
			validationErrors.Incorrect = &message
			return response, domain.ValidationError{Errors: validationErrors}
		}
		return response, err
	}

	tokenString, err := function.GenerateAccessToken(loginResult.UserId, loginResult.UserType, loginResult.Username)
	if err != nil {
		return response, err
	}

	response = domain.UserLoginResponse{
		Accesstoken: tokenString,
	}

	return response, nil
}
