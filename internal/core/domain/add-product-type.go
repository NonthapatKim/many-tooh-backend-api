package domain

type AddProductTypeRequest struct {
	AccessToken     string
	ProductTypeName string `json:"product_type_name"`
}

type AddProductTypeResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
