package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) AddProductCategory(req domain.AddProductCategoryRequest) (domain.AddProductCategoryResponse, error) {
	query := `
		INSERT INTO product_categories (
			name
		) VALUES (?)
	`
	_, err := r.db.Exec(
		query,
		req.CategoryName,
	)
	if err != nil {
		return domain.AddProductCategoryResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.AddProductCategoryResponse{}, nil
}
