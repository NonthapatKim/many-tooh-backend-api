package handler

import (
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain"
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetProductCategories(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized - no token",
		})
	}

	req := domain.GetProductCategoriesRequest{
		AccessToken: accessToken,
	}

	result, err := h.svc.GetProductCategories(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
