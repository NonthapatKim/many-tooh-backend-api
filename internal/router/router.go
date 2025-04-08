package router

import (
	"fmt"
	"os"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/adapter/handler"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/adapter/handler/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	app *fiber.App
}

const serviceBaseURL = "/api"

func NewRouter(h handler.Handler) (*Router, error) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, https://many-tooh.web.app/",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	basePath := app.Group(serviceBaseURL)
	basePathV1 := basePath.Group("/v1").Use(middleware.LoggerMiddleware())

	basePathV1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("สวัสดี !")
	})

	brand := basePathV1.Group("/brands")
	{
		brand.Get("/", h.GetBrands)
		brand.Post("/add", h.AddBrand)
		brand.Put("/:brandId", h.UpdateBrandById)
		brand.Delete("/:brandId", h.DeleteBrandById)
	}

	product := basePathV1.Group("/products")
	{
		product.Get("/", h.GetProducts)
		product.Post("/add", h.AddProduct)
		product.Patch("/:productId", h.UpdateProductById)
		product.Delete("/:productId", h.DeleteProductById)
	}

	category := basePathV1.Group("/categories")
	{
		category.Get("/", h.GetProductCategories)
		category.Post("/add", h.AddProductCategory)
		category.Put("/:categoryId", h.UpdateProductCategoryById)
		category.Delete("/:categoryId", h.DeleteProductCategoryById)
	}

	categoryType := basePathV1.Group("/types")
	{
		categoryType.Get("/", h.GetProductType)
		categoryType.Post("/add", h.AddProductType)
		categoryType.Put("/:productTypeId", h.UpdateProductTypeById)
		categoryType.Delete("/:productTypeId", h.DeleteProductTypeById)
	}

	user := basePathV1.Group("/users")
	{
		user.Post("/create", h.CreateStaffUser)
		user.Post("/login", h.UserLogin)
		user.Post("/logout", h.UserLogOut)

		user.Get("/authenticate", h.UserAuthenticate)
	}

	return &Router{app: app}, nil
}

func (r *Router) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port", port)
	return r.app.Listen(":" + port)
}
