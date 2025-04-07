package service

import (
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("customEmail", function.CustomEmailValidator)
	validate.RegisterValidation("customPassword", function.CustomPasswordValidator)
}

func (s *service) CreateStaffUser(req domain.CreateStaffUserRequest) (domain.CreateStaffUserResponse, error) {
	var validationErrors domain.ValidationErrorResponse
	var response domain.CreateStaffUserResponse

	// Validate struct using go-playground/validator with custom rules
	err := validate.Struct(req)
	if err != nil {
		validationErrors = ProcessValidationError(err)
		return response, domain.ValidationError{Errors: validationErrors}
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return response, err
		}
		req.Password = string(hashedPassword)
	}

	_, err = s.repo.CreateStaffUser(req)
	if err != nil {
		return response, err
	}

	response = domain.CreateStaffUserResponse{
		Code:    200,
		Message: "successfully created user",
	}

	return response, nil
}
