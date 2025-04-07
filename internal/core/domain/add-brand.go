package domain

type AddBrandRequest struct {
	AccessToken string
	BrandName   string `json:"brand_name"`
}

type AddBrandResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
