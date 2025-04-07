package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
)

func (s *service) DeleteBrandById(req domain.DeleteBrandByIdRequest) (domain.DeleteBrandByIdResponse, error) {
	var checkList []domain.CheckExistsRequest

	if req.AccessToken == "" {
		return domain.DeleteBrandByIdResponse{}, errors.New("token is required")
	}

	if req.BrandId == "" {
		return domain.DeleteBrandByIdResponse{}, errors.New("brand_id is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.DeleteBrandByIdResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.DeleteBrandByIdResponse{}, errors.New("invalid token: missing user_id")
	}

	checkList = []domain.CheckExistsRequest{
		{
			Table:  "users",
			Column: "user_id",
			Id:     &userId,
		},
		{
			Table:  "brands",
			Column: "brand_id",
			Id:     &req.BrandId,
		},
	}

	for _, check := range checkList {
		exists, err := s.repo.CheckExists(check)
		if err != nil {
			return domain.DeleteBrandByIdResponse{}, err
		}
		if !exists.Exists {
			return domain.DeleteBrandByIdResponse{}, fmt.Errorf("error: %s not found in %s", check.Column, check.Table)
		}
	}

	_, err = s.repo.DeleteBrandById(req)
	if err != nil {
		return domain.DeleteBrandByIdResponse{}, err
	}

	response := domain.DeleteBrandByIdResponse{
		Code:    200,
		Message: "successfully deleted product",
	}

	return response, nil
}
