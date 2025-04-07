package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) DeleteProductTypeById(req domain.DeleteProductTypeByIdRequest) (domain.DeleteProductTypeByIdResponse, error) {
	query := `DELETE FROM product_type WHERE product_type_id = ?`
	_, err := r.db.Exec(
		query,
		req.ProductTypeId,
	)
	if err != nil {
		return domain.DeleteProductTypeByIdResponse{}, fmt.Errorf("error deleting: %w", err)
	}

	return domain.DeleteProductTypeByIdResponse{}, nil
}
