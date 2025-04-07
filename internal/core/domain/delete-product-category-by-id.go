package domain

type DeleteProductCategoryByIdRequest struct {
	AccessToken string
	CategoryId  string
}

type DeleteProductCategoryByIdResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
