package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
)

func (s *service) UpdateProductCategoryById(req domain.UpdateProductCategoryByIdRequest) (domain.UpdateProductCategoryByIdResponse, error) {
	var checkList []domain.CheckExistsRequest

	if req.AccessToken == "" {
		return domain.UpdateProductCategoryByIdResponse{}, errors.New("token is required")
	}

	if req.CategoryId == "" {
		return domain.UpdateProductCategoryByIdResponse{}, errors.New("category_id is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UpdateProductCategoryByIdResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.UpdateProductCategoryByIdResponse{}, errors.New("invalid token: missing user_id")
	}

	checkList = []domain.CheckExistsRequest{
		{
			Table:  "users",
			Column: "user_id",
			Id:     &userId,
		},
		{
			Table:  "product_categories",
			Column: "category_id",
			Id:     &req.CategoryId,
		},
	}

	for _, check := range checkList {
		exists, err := s.repo.CheckExists(check)
		if err != nil {
			return domain.UpdateProductCategoryByIdResponse{}, err
		}
		if !exists.Exists {
			return domain.UpdateProductCategoryByIdResponse{}, fmt.Errorf("error: %s not found in %s", check.Column, check.Table)
		}
	}

	_, err = s.repo.UpdateProductCategoryById(req)
	if err != nil {
		return domain.UpdateProductCategoryByIdResponse{}, err
	}

	response := domain.UpdateProductCategoryByIdResponse{
		Code:    200,
		Message: "successfully updated product category",
	}

	return response, nil
}
