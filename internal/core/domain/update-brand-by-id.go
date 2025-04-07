package domain

type UpdateBrandByIdRequest struct {
	AccessToken string
	BrandId     string
	BrandName   string `json:"brand_name"`
}

type UpdateBrandByIdResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
