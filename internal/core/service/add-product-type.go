package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
)

func (s *service) AddProductType(req domain.AddProductTypeRequest) (domain.AddProductTypeResponse, error) {
	if req.AccessToken == "" {
		return domain.AddProductTypeResponse{}, errors.New("token is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.AddProductTypeResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.AddProductTypeResponse{}, errors.New("invalid token: missing user_id")
	}

	reqUserExists := domain.CheckExistsRequest{
		Table:  "users",
		Column: "user_id",
		Id:     &userId,
	}

	exists, err := s.repo.CheckExists(reqUserExists)
	if err != nil {
		return domain.AddProductTypeResponse{}, err
	}
	if !exists.Exists {
		return domain.AddProductTypeResponse{}, errors.New("error: user not found")
	}

	_, err = s.repo.AddProductType(req)
	if err != nil {
		return domain.AddProductTypeResponse{}, err
	}

	response := domain.AddProductTypeResponse{
		Code:    200,
		Message: "successfully created product type",
	}

	return response, nil
}
