package handler

import (
	"time"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) UserLogin(c *fiber.Ctx) error {
	var user domain.UserLoginRequest

	if err := c.BodyParser(&user); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	result, err := h.svc.UserLogin(user)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}
		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    result.Accesstoken,
		Expires:  time.Now().Add(6 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	return response.JSONSuccessResponse(c, fiber.Map{
		"message": "successfully login",
		"code":    200,
	})
}
