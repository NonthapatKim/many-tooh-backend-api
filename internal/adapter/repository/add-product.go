package repository

import (
	"fmt"

	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"
)

func (r *repository) AddProduct(req domain.AddProductRequest) (domain.AddProductResponse, error) {
	query := `
		INSERT INTO products (
			brand_id,
			product_category_id,
			product_type_id,
			image_url,
			name,
			barcode,
			warning,
			usage_description,
			amount_fluoride,
			properties,
			active_ingredient,
			dangerous_ingredient,
			is_dangerous
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		req.BrandId,
		req.ProductCategoryId,
		req.ProductTypeId,
		req.ImageUrl,
		req.ProductName,
		req.Barcode,
		req.Warning,
		req.UsageDescription,
		req.AmountFluoride,
		req.Properties,
		req.ActiveIngredient,
		req.DangerousIngredient,
		req.IsDangerous,
	)
	if err != nil {
		return domain.AddProductResponse{}, fmt.Errorf("error creating: %w", err)
	}

	return domain.AddProductResponse{}, nil
}
