package domain

type UpdateProductTypeByIdRequest struct {
	AccessToken     string
	ProductTypeId   string
	ProductTypeName string `json:"product_type_name"`
}

type UpdateProductTypeByIdResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
