package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
)

func (s *service) DeleteProductCategoryById(req domain.DeleteProductCategoryByIdRequest) (domain.DeleteProductCategoryByIdResponse, error) {
	var checkList []domain.CheckExistsRequest

	if req.AccessToken == "" {
		return domain.DeleteProductCategoryByIdResponse{}, errors.New("token is required")
	}

	if req.CategoryId == "" {
		return domain.DeleteProductCategoryByIdResponse{}, errors.New("category_id is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.DeleteProductCategoryByIdResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.DeleteProductCategoryByIdResponse{}, errors.New("invalid token: missing user_id")
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
			return domain.DeleteProductCategoryByIdResponse{}, err
		}
		if !exists.Exists {
			return domain.DeleteProductCategoryByIdResponse{}, fmt.Errorf("error: %s not found in %s", check.Column, check.Table)
		}
	}

	_, err = s.repo.DeleteProductCategoryById(req)
	if err != nil {
		return domain.DeleteProductCategoryByIdResponse{}, err
	}

	response := domain.DeleteProductCategoryByIdResponse{
		Code:    200,
		Message: "successfully deleted product",
	}

	return response, nil
}
