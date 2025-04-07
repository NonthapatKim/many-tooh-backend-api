package handler

import (
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	// Brand
	AddBrand(c *fiber.Ctx) error
	DeleteBrandById(c *fiber.Ctx) error
	GetBrands(c *fiber.Ctx) error
	UpdateBrandById(c *fiber.Ctx) error

	// Product
	AddProduct(c *fiber.Ctx) error
	DeleteProductById(c *fiber.Ctx) error
	GetProducts(c *fiber.Ctx) error
	UpdateProductById(c *fiber.Ctx) error

	// ProductCategory
	AddProductCategory(c *fiber.Ctx) error
	GetProductCategories(c *fiber.Ctx) error
	UpdateProductCategoryById(c *fiber.Ctx) error
	DeleteProductCategoryById(c *fiber.Ctx) error

	// ProductType
	AddProductType(c *fiber.Ctx) error
	DeleteProductTypeById(c *fiber.Ctx) error
	GetProductType(c *fiber.Ctx) error
	UpdateProductTypeById(c *fiber.Ctx) error

	// User
	CreateStaffUser(c *fiber.Ctx) error
	UserAuthenticate(c *fiber.Ctx) error
	UserLogin(c *fiber.Ctx) error
	UserLogOut(c *fiber.Ctx) error
}

type handler struct {
	svc port.Service
}

func New(svc port.Service) Handler {
	return &handler{svc: svc}
}
