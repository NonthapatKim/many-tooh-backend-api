package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) CreateStaffUser(req domain.CreateStaffUserRequest) (domain.CreateStaffUserResponse, error) {
	query := `
		INSERT INTO users (
			email,
			username,
			password,
			user_type
		) VALUES (?, ?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		req.Email,
		req.Username,
		req.Password,
		"admin",
	)
	if err != nil {
		return domain.CreateStaffUserResponse{}, fmt.Errorf("error creating user: %w", err)
	}

	return domain.CreateStaffUserResponse{}, nil
}
