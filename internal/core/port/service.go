package port

import "github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"

type Service interface {
	// User
	CreateStaffUser(req domain.CreateStaffUserRequest) (domain.CreateStaffUserResponse, error)
}
