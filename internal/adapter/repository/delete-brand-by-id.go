package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) DeleteBrandById(req domain.DeleteBrandByIdRequest) (domain.DeleteBrandByIdResponse, error) {
	query := `DELETE FROM brands WHERE brand_id = ?`
	_, err := r.db.Exec(
		query,
		req.BrandId,
	)
	if err != nil {
		return domain.DeleteBrandByIdResponse{}, fmt.Errorf("error deleting: %w", err)
	}

	return domain.DeleteBrandByIdResponse{}, nil
}
