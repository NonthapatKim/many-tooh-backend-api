package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
)

func (s *service) UpdateProductTypeById(req domain.UpdateProductTypeByIdRequest) (domain.UpdateProductTypeByIdResponse, error) {
	var checkList []domain.CheckExistsRequest

	if req.AccessToken == "" {
		return domain.UpdateProductTypeByIdResponse{}, errors.New("token is required")
	}

	if req.ProductTypeId == "" {
		return domain.UpdateProductTypeByIdResponse{}, errors.New("product_type_id is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UpdateProductTypeByIdResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.UpdateProductTypeByIdResponse{}, errors.New("invalid token: missing user_id")
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
			return domain.UpdateProductTypeByIdResponse{}, err
		}
		if !exists.Exists {
			return domain.UpdateProductTypeByIdResponse{}, fmt.Errorf("error: %s not found in %s", check.Column, check.Table)
		}
	}

	_, err = s.repo.UpdateProductTypeById(req)
	if err != nil {
		return domain.UpdateProductTypeByIdResponse{}, err
	}

	response := domain.UpdateProductTypeByIdResponse{
		Code:    200,
		Message: "successfully updated product category",
	}

	return response, nil
}
