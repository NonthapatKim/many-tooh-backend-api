package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
)

func (s *service) UserAuthenticate(req domain.UserAuthenticateRequest) (domain.UserAuthenticateResponse, error) {
	if req.AccessToken == "" {
		return domain.UserAuthenticateResponse{}, errors.New("token is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UserAuthenticateResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.UserAuthenticateResponse{}, errors.New("invalid token claims: missing staff_id")
	}

	role, ok := claims["role"].(string)
	if !ok || role == "" {
		return domain.UserAuthenticateResponse{}, errors.New("invalid token claims: missing role")
	}

	username, ok := claims["username"].(string)
	if !ok || username == "" {
		return domain.UserAuthenticateResponse{}, errors.New("invalid token claims: missing username")
	}

	response := domain.UserAuthenticateResponse{
		UserId:   userId,
		Role:     role,
		Username: username,
	}

	return response, nil
}
