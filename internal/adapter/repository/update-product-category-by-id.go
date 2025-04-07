package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) UpdateProductCategoryById(req domain.UpdateProductCategoryByIdRequest) (domain.UpdateProductCategoryByIdResponse, error) {
	query := `
		UPDATE product_categories SET
			name = ?
		WHERE category_id = ?
	`
	_, err := r.db.Exec(
		query,
		req.CategoryName,
		req.CategoryId,
	)
	if err != nil {
		return domain.UpdateProductCategoryByIdResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.UpdateProductCategoryByIdResponse{}, nil
}
