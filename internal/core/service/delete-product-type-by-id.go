package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
)

func (s *service) DeleteProductTypeById(req domain.DeleteProductTypeByIdRequest) (domain.DeleteProductTypeByIdResponse, error) {
	var checkList []domain.CheckExistsRequest

	if req.AccessToken == "" {
		return domain.DeleteProductTypeByIdResponse{}, errors.New("token is required")
	}

	if req.ProductTypeId == "" {
		return domain.DeleteProductTypeByIdResponse{}, errors.New("product_type_id is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.DeleteProductTypeByIdResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.DeleteProductTypeByIdResponse{}, errors.New("invalid token: missing user_id")
	}

	checkList = []domain.CheckExistsRequest{
		{
			Table:  "users",
			Column: "user_id",
			Id:     &userId,
		},
		{
			Table:  "product_type",
			Column: "product_type_id",
			Id:     &req.ProductTypeId,
		},
	}

	for _, check := range checkList {
		exists, err := s.repo.CheckExists(check)
		if err != nil {
			return domain.DeleteProductTypeByIdResponse{}, err
		}
		if !exists.Exists {
			return domain.DeleteProductTypeByIdResponse{}, fmt.Errorf("error: %s not found in %s", check.Column, check.Table)
		}
	}

	_, err = s.repo.DeleteProductTypeById(req)
	if err != nil {
		return domain.DeleteProductTypeByIdResponse{}, err
	}

	response := domain.DeleteProductTypeByIdResponse{
		Code:    200,
		Message: "successfully deleted product",
	}

	return response, nil
}
