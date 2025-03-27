package domain

type GetBrandsRequest struct {
	AccessToken string
}

type GetBrandsResponse struct {
	BrandId   string `json:"brand_id"`
	BrandName string `json:"brand_name"`
}
