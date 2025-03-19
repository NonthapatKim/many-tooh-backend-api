package router

import (
	"fmt"
	"os"

	"github.com/NonthapatKim/many_tooth_backend_api/internal/adapter/handler"
	"github.com/NonthapatKim/many_tooth_backend_api/internal/adapter/handler/middleware"
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
		AllowOrigins: "*", // For Develop Only
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	basePath := app.Group(serviceBaseURL)
	basePathV1 := basePath.Group("/v1").Use(middleware.LoggerMiddleware())

	basePathV1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("สวัสดี !")
	})

	return &Router{app: app}, nil
}

func (r *Router) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4300"
	}

	fmt.Println("Listening on port", port)
	return r.app.Listen(":" + port)
}
