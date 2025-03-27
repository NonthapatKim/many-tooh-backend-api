package domain

import "mime/multipart"

type AddProductRequest struct {
	AccessToken         string
	BrandId             string `json:"brand_id"`
	ProductCategoryId   string `json:"product_category_id"`
	ProductTypeId       string `json:"product_type_id"`
	Image               multipart.File
	ImageUrl            string
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

type AddProductResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
