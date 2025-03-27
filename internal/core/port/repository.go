package port

import "github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"

type Repository interface {
	// Brands
	GetBrands() ([]domain.GetBrandsResponse, error)

	// Product
	AddProduct(req domain.AddProductRequest) (domain.AddProductResponse, error)
	DeleteProductById(req domain.DeleteProductByIdRequest) (domain.DeleteProductByIdResponse, error)
	GetProducts() ([]domain.GetProductsResponse, error)
	GetProductCategories() ([]domain.GetProductCategoriesResponse, error)
	GetProductType() ([]domain.GetProuctTypeResponse, error)
	UpdateProductById(req domain.UpdateProductByIdRequest) (domain.UpdateProductByIdResponse, error)

	// User
	CreateStaffUser(req domain.CreateStaffUserRequest) (domain.CreateStaffUserResponse, error)
	UserLogin(req domain.UserLoginRequest) (domain.UserLoginResult, error)

	// Other
	CheckExists(req domain.CheckExistsRequest) (domain.CheckExistsResponse, error)
}
