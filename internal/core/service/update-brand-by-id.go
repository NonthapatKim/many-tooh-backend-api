package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
)

func (s *service) UpdateBrandById(req domain.UpdateBrandByIdRequest) (domain.UpdateBrandByIdResponse, error) {
	var checkList []domain.CheckExistsRequest

	if req.AccessToken == "" {
		return domain.UpdateBrandByIdResponse{}, errors.New("token is required")
	}

	if req.BrandId == "" {
		return domain.UpdateBrandByIdResponse{}, errors.New("brand_id is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UpdateBrandByIdResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.UpdateBrandByIdResponse{}, errors.New("invalid token: missing user_id")
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
			return domain.UpdateBrandByIdResponse{}, err
		}
		if !exists.Exists {
			return domain.UpdateBrandByIdResponse{}, fmt.Errorf("error: %s not found in %s", check.Column, check.Table)
		}
	}

	_, err = s.repo.UpdateBrandById(req)
	if err != nil {
		return domain.UpdateBrandByIdResponse{}, err
	}

	response := domain.UpdateBrandByIdResponse{
		Code:    200,
		Message: "successfully updated brand",
	}

	return response, nil
}
