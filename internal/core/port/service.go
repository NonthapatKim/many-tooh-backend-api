package port

import "github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"

type Service interface {
	// Brands
	GetBrands(req domain.GetBrandsRequest) ([]domain.GetBrandsResponse, error)

	// Product
	AddProduct(req domain.AddProductRequest) (domain.AddProductResponse, error)
	DeleteProductById(req domain.DeleteProductByIdRequest) (domain.DeleteProductByIdResponse, error)
	GetProducts(req domain.GetProductsRequest) ([]domain.GetProductsResponse, error)
	GetProductCategories(req domain.GetProductCategoriesRequest) ([]domain.GetProductCategoriesResponse, error)
	GetProductType(req domain.GetProuctTypeRequest) ([]domain.GetProuctTypeResponse, error)
	UpdateProductById(req domain.UpdateProductByIdRequest) (domain.UpdateProductByIdResponse, error)

	// User
	CreateStaffUser(req domain.CreateStaffUserRequest) (domain.CreateStaffUserResponse, error)
	UserAuthenticate(req domain.UserAuthenticateRequest) (domain.UserAuthenticateResponse, error)
	UserLogin(req domain.UserLoginRequest) (domain.UserLoginResponse, error)
	UserLogout() (domain.UserLogoutResponse, error)
}
