package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
)

func (s *service) AddBrand(req domain.AddBrandRequest) (domain.AddBrandResponse, error) {
	if req.AccessToken == "" {
		return domain.AddBrandResponse{}, errors.New("token is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.AddBrandResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.AddBrandResponse{}, errors.New("invalid token: missing user_id")
	}

	reqUserExists := domain.CheckExistsRequest{
		Table:  "users",
		Column: "user_id",
		Id:     &userId,
	}

	exists, err := s.repo.CheckExists(reqUserExists)
	if err != nil {
		return domain.AddBrandResponse{}, err
	}
	if !exists.Exists {
		return domain.AddBrandResponse{}, errors.New("error: user not found")
	}

	_, err = s.repo.AddBrand(req)
	if err != nil {
		return domain.AddBrandResponse{}, err
	}

	response := domain.AddBrandResponse{
		Code:    200,
		Message: "successfully created product category",
	}

	return response, nil
}
