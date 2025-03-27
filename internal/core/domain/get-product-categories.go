package domain

type GetProductCategoriesRequest struct {
	AccessToken string
}

type GetProductCategoriesResponse struct {
	ProductCategoryId   string `json:"product_category_id"`
	ProductCategoryName string `json:"product_category_name"`
}
