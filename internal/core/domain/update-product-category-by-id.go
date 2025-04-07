package domain

type UpdateProductCategoryByIdRequest struct {
	AccessToken  string
	CategoryId   string
	CategoryName string `json:"product_category_name"`
}

type UpdateProductCategoryByIdResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
