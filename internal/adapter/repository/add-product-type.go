package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) AddProductType(req domain.AddProductTypeRequest) (domain.AddProductTypeResponse, error) {
	query := `
		INSERT INTO product_type (
			name
		) VALUES (?)
	`
	_, err := r.db.Exec(
		query,
		req.ProductTypeName,
	)
	if err != nil {
		return domain.AddProductTypeResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.AddProductTypeResponse{}, nil
}
