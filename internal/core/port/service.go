package port

import "github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"

type Service interface {
	// Brands
	AddBrand(req domain.AddBrandRequest) (domain.AddBrandResponse, error)
	DeleteBrandById(req domain.DeleteBrandByIdRequest) (domain.DeleteBrandByIdResponse, error)
	GetBrands(req domain.GetBrandsRequest) ([]domain.GetBrandsResponse, error)
	UpdateBrandById(req domain.UpdateBrandByIdRequest) (domain.UpdateBrandByIdResponse, error)

	// Product
	AddProduct(req domain.AddProductRequest) (domain.AddProductResponse, error)
	DeleteProductById(req domain.DeleteProductByIdRequest) (domain.DeleteProductByIdResponse, error)
	GetProducts(req domain.GetProductsRequest) ([]domain.GetProductsResponse, error)
	UpdateProductById(req domain.UpdateProductByIdRequest) (domain.UpdateProductByIdResponse, error)

	// ProductCategory
	AddProductCategory(req domain.AddProductCategoryRequest) (domain.AddProductCategoryResponse, error)
	DeleteProductCategoryById(req domain.DeleteProductCategoryByIdRequest) (domain.DeleteProductCategoryByIdResponse, error)
	GetProductCategories(req domain.GetProductCategoriesRequest) ([]domain.GetProductCategoriesResponse, error)
	UpdateProductCategoryById(req domain.UpdateProductCategoryByIdRequest) (domain.UpdateProductCategoryByIdResponse, error)

	// ProductType
	AddProductType(req domain.AddProductTypeRequest) (domain.AddProductTypeResponse, error)
	DeleteProductTypeById(req domain.DeleteProductTypeByIdRequest) (domain.DeleteProductTypeByIdResponse, error)
	GetProductType(req domain.GetProuctTypeRequest) ([]domain.GetProuctTypeResponse, error)
	UpdateProductTypeById(req domain.UpdateProductTypeByIdRequest) (domain.UpdateProductTypeByIdResponse, error)

	// User
	CreateStaffUser(req domain.CreateStaffUserRequest) (domain.CreateStaffUserResponse, error)
	UserAuthenticate(req domain.UserAuthenticateRequest) (domain.UserAuthenticateResponse, error)
	UserLogin(req domain.UserLoginRequest) (domain.UserLoginResponse, error)
	UserLogout() (domain.UserLogoutResponse, error)
}
