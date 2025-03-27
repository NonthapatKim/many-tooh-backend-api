package handler

import (
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	// Brand
	GetBrands(c *fiber.Ctx) error

	// Product
	AddProduct(c *fiber.Ctx) error
	DeleteProductById(c *fiber.Ctx) error
	GetProducts(c *fiber.Ctx) error
	GetProductCategories(c *fiber.Ctx) error
	GetProductType(c *fiber.Ctx) error
	UpdateProductById(c *fiber.Ctx) error

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
