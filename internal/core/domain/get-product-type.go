package domain

type GetProuctTypeRequest struct {
	AccessToken string
}

type GetProuctTypeResponse struct {
	ProductTypeId   string `json:"product_type_id"`
	ProductTypeName string `json:"product_type_name"`
}
