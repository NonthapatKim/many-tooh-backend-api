package handler

import (
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) AddProductType(c *fiber.Ctx) error {
	var productType domain.AddProductTypeRequest

	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized - no token",
		})
	}

	if err := c.BodyParser(&productType); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.AddProductTypeRequest{
		AccessToken:     accessToken,
		ProductTypeName: productType.ProductTypeName,
	}

	result, err := h.svc.AddProductType(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
