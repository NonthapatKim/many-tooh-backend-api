package domain

import "mime/multipart"

type UpdateProductByIdRequest struct {
	AccessToken         string
	ProductId           string `json:"product_id"`
	BrandId             string `json:"brand_id"`
	ProductCategoryId   string `json:"product_category_id"`
	ProductTypeId       string `json:"product_type_id"`
	Image               *multipart.File
	ProductImageUrl     *string `json:"product_image_url"`
	ProductName         string  `json:"product_name"`
	Barcode             string  `json:"bar_code"`
	Warning             *string `json:"warning"`
	UsageDescription    *string `json:"usage_description"`
	AmountFluoride      *string `json:"amount_fluoride"`
	Properties          *string `json:"properties"`
	ActiveIngredient    *string `json:"active_ingredient"`
	DangerousIngredient *string `json:"dangerous_ingredient"`
	IsDangerous         string  `json:"is_dangerous"`
}

type UpdateProductByIdResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
