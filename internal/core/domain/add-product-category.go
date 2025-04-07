package domain

type AddProductCategoryRequest struct {
	AccessToken  string
	CategoryName string `json:"product_category_name"`
}

type AddProductCategoryResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
