package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) AddBrand(req domain.AddBrandRequest) (domain.AddBrandResponse, error) {
	query := `
		INSERT INTO brands (
			name
		) VALUES (?)
	`
	_, err := r.db.Exec(
		query,
		req.BrandName,
	)
	if err != nil {
		return domain.AddBrandResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.AddBrandResponse{}, nil
}
