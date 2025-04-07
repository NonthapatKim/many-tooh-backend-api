package domain

type DeleteProductTypeByIdRequest struct {
	AccessToken   string
	ProductTypeId string
}

type DeleteProductTypeByIdResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
