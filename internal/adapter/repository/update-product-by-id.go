package repository

import (
	"fmt"
	"time"

	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"
)

func (r *repository) UpdateProductById(req domain.UpdateProductByIdRequest) (domain.UpdateProductByIdResponse, error) {
	query := `
		UPDATE products SET
			brand_id = ?,
			product_category_id = ?,
			product_type_id = ?,
			image_url = ?,
			name = ?,
			barcode = ?,
			warning = ?,
			usage_description = ?,
			amount_fluoride = ?,
			properties = ?,
			active_ingredient = ?,
			dangerous_ingredient = ?,
			is_dangerous = ?,
			updated_at = ?
		WHERE product_id = ?
	`
	_, err := r.db.Exec(
		query,
		req.BrandId,
		req.ProductCategoryId,
		req.ProductTypeId,
		req.ProductImageUrl,
		req.ProductName,
		req.Barcode,
		req.Warning,
		req.UsageDescription,
		req.AmountFluoride,
		req.Properties,
		req.ActiveIngredient,
		req.DangerousIngredient,
		req.IsDangerous,
		time.Now(),
		req.ProductId,
	)
	if err != nil {
		return domain.UpdateProductByIdResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.UpdateProductByIdResponse{}, nil
}
