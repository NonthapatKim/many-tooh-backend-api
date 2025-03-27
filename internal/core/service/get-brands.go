package service

import (
	"errors"

	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/function"
)

func (s *service) GetBrands(req domain.GetBrandsRequest) ([]domain.GetBrandsResponse, error) {
	if req.AccessToken == "" {
		return nil, errors.New("token is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return nil, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return nil, errors.New("invalid token: missing user_id")
	}

	reqUserExists := domain.CheckExistsRequest{
		Table:  "users",
		Column: "user_id",
		Id:     &userId,
	}

	exists, err := s.repo.CheckExists(reqUserExists)
	if err != nil {
		return nil, err
	}
	if !exists.Exists {
		return nil, errors.New("error: user not found")
	}

	result, err := s.repo.GetBrands()
	if err != nil {
		return nil, err
	}

	return result, nil
}
