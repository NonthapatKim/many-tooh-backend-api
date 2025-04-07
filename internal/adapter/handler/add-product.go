package handler

import (
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func nullableString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func (h *handler) AddProduct(c *fiber.Ctx) error {
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

	req := domain.AddProductRequest{
		AccessToken:         accessToken,
		BrandId:             c.FormValue("brand_id"),
		ProductCategoryId:   c.FormValue("product_category_id"),
		ProductTypeId:       c.FormValue("product_type_id"),
		ProductName:         c.FormValue("product_name"),
		Barcode:             c.FormValue("barcode"),
		Warning:             nullableString(c.FormValue("warning")),
		UsageDescription:    nullableString(c.FormValue("usage_description")),
		AmountFluoride:      nullableString(c.FormValue("amount_fluoride")),
		Properties:          nullableString(c.FormValue("properties")),
		ActiveIngredient:    nullableString(c.FormValue("active_ingredient")),
		DangerousIngredient: nullableString(c.FormValue("dangerous_ingredient")),
		IsDangerous:         c.FormValue("is_dangerous"),
		Image:               src,
	}

	result, err := h.svc.AddProduct(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
