package domain

type DeleteProductByIdRequest struct {
	AccessToken string
	ProductId   string
}

type DeleteProductByIdResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
