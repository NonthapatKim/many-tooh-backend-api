package handler

import (
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func nullableString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func (h *handler) AddProduct(c *fiber.Ctx) error {
	var product domain.AddProductRequest

	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized - no token",
		})
	}

	file, err := c.FormFile("product_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get image file",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to open image file",
		})
	}
	defer src.Close()

	product.AccessToken = accessToken
	product.BrandId = c.FormValue("brand_id")
	product.ProductCategoryId = c.FormValue("product_category_id")
	product.ProductTypeId = c.FormValue("product_type_id")
	product.ProductName = c.FormValue("product_name")
	product.Barcode = c.FormValue("barcode")
	product.Warning = nullableString(c.FormValue("warning"))
	product.UsageDescription = nullableString(c.FormValue("usage_description"))
	product.AmountFluoride = nullableString(c.FormValue("amount_fluoride"))
	product.Properties = nullableString(c.FormValue("properties"))
	product.ActiveIngredient = nullableString(c.FormValue("active_ingredient"))
	product.DangerousIngredient = nullableString(c.FormValue("dangerous_ingredient"))
	product.IsDangerous = c.FormValue("is_dangerous")
	product.Image = src

	result, err := h.svc.AddProduct(product)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
