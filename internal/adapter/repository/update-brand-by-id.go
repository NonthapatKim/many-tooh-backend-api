package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) UpdateBrandById(req domain.UpdateBrandByIdRequest) (domain.UpdateBrandByIdResponse, error) {
	query := `
		UPDATE brands SET
			name = ?
		WHERE brand_id = ?
	`
	_, err := r.db.Exec(
		query,
		req.BrandName,
		req.BrandId,
	)
	if err != nil {
		return domain.UpdateBrandByIdResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.UpdateBrandByIdResponse{}, nil
}
