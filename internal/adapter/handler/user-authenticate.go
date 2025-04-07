package handler

import (
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) UserAuthenticate(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")

	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized - no token",
		})
	}

	req := domain.UserAuthenticateRequest{
		AccessToken: accessToken,
	}

	result, err := h.svc.UserAuthenticate(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
