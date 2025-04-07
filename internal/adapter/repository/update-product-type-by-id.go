package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
)

func (r *repository) UpdateProductTypeById(req domain.UpdateProductTypeByIdRequest) (domain.UpdateProductTypeByIdResponse, error) {
	query := `
		UPDATE product_type SET
			name = ?
		WHERE product_type_id = ?
	`
	_, err := r.db.Exec(
		query,
		req.ProductTypeName,
		req.ProductTypeId,
	)
	if err != nil {
		return domain.UpdateProductTypeByIdResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.UpdateProductTypeByIdResponse{}, nil
}
