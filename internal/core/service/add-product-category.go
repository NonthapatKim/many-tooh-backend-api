package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
)

func (s *service) AddProductCategory(req domain.AddProductCategoryRequest) (domain.AddProductCategoryResponse, error) {
	if req.AccessToken == "" {
		return domain.AddProductCategoryResponse{}, errors.New("token is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.AddProductCategoryResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.AddProductCategoryResponse{}, errors.New("invalid token: missing user_id")
	}

	reqUserExists := domain.CheckExistsRequest{
		Table:  "users",
		Column: "user_id",
		Id:     &userId,
	}

	exists, err := s.repo.CheckExists(reqUserExists)
	if err != nil {
		return domain.AddProductCategoryResponse{}, err
	}
	if !exists.Exists {
		return domain.AddProductCategoryResponse{}, errors.New("error: user not found")
	}

	_, err = s.repo.AddProductCategory(req)
	if err != nil {
		return domain.AddProductCategoryResponse{}, err
	}

	response := domain.AddProductCategoryResponse{
		Code:    200,
		Message: "successfully created product category",
	}

	return response, nil
}
