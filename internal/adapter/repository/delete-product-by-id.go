package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) DeleteProductById(req domain.DeleteProductByIdRequest) (domain.DeleteProductByIdResponse, error) {
	query := `DELETE FROM products WHERE product_id = ?`
	_, err := r.db.Exec(
		query,
		req.ProductId,
	)
	if err != nil {
		return domain.DeleteProductByIdResponse{}, fmt.Errorf("error deleting: %w", err)
	}

	return domain.DeleteProductByIdResponse{}, nil
}
