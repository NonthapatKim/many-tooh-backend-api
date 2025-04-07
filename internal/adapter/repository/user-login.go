package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
)

func (r *repository) UserLogin(req domain.UserLoginRequest) (domain.UserLoginResult, error) {
	var result domain.UserLoginResult

	fmt.Println(req.Username)

	err := r.db.QueryRow(
		`SELECT 
			user_id,
			email,
			username,
			password,
			user_type,
			created_at,
			updated_at
		FROM users
		WHERE username = ?`,
		req.Username,
	).Scan(
		&result.UserId,
		&result.Email,
		&result.Username,
		&result.Password,
		&result.UserType,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, errors.New("user not found")
		}
		return result, fmt.Errorf("failed to query user: %w", err)
	}

	if bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password)) != nil {
		return result, errors.New("incorrect password")
	}

	result.Password = ""

	return result, nil
}
