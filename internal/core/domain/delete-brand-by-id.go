package domain

type DeleteBrandByIdRequest struct {
	AccessToken string
	BrandId     string
}

type DeleteBrandByIdResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
