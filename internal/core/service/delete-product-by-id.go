package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/function"
)

func (s *service) DeleteProductById(req domain.DeleteProductByIdRequest) (domain.DeleteProductByIdResponse, error) {
	var checkList []domain.CheckExistsRequest

	if req.AccessToken == "" {
		return domain.DeleteProductByIdResponse{}, errors.New("token is required")
	}

	if req.ProductId == "" {
		return domain.DeleteProductByIdResponse{}, errors.New("product_id is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.DeleteProductByIdResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.DeleteProductByIdResponse{}, errors.New("invalid token: missing user_id")
	}

	checkList = []domain.CheckExistsRequest{
		{
			Table:  "users",
			Column: "user_id",
			Id:     &userId,
		},
		{
			Table:  "products",
			Column: "product_id",
			Id:     &req.ProductId,
		},
	}

	for _, check := range checkList {
		exists, err := s.repo.CheckExists(check)
		if err != nil {
			return domain.DeleteProductByIdResponse{}, err
		}
		if !exists.Exists {
			return domain.DeleteProductByIdResponse{}, fmt.Errorf("error: %s not found in %s", check.Column, check.Table)
		}
	}

	_, err = s.repo.DeleteProductById(req)
	if err != nil {
		return domain.DeleteProductByIdResponse{}, err
	}

	response := domain.DeleteProductByIdResponse{
		Code:    200,
		Message: "successfully updated product",
	}

	return response, nil
}
