package service

import "github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"

func (s *service) UserLogout() (domain.UserLogoutResponse, error) {
	return domain.UserLogoutResponse{
		Code:    200,
		Message: "successfully log out",
	}, nil
}
