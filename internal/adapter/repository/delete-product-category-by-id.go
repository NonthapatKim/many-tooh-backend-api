package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) DeleteProductCategoryById(req domain.DeleteProductCategoryByIdRequest) (domain.DeleteProductCategoryByIdResponse, error) {
	query := `DELETE FROM product_categories WHERE category_id = ?`
	_, err := r.db.Exec(
		query,
		req.CategoryId,
	)
	if err != nil {
		return domain.DeleteProductCategoryByIdResponse{}, fmt.Errorf("error deleting: %w", err)
	}

	return domain.DeleteProductCategoryByIdResponse{}, nil
}
