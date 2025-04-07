package handler

import (
	"mime/multipart"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) UpdateProductById(c *fiber.Ctx) error {
	productId := c.Params("productId")

	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized - no token",
		})
	}

	imageFile, _ := c.FormFile("product_image")
	var src multipart.File

	if imageFile != nil {
		var err error
		src, err = imageFile.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to open image file",
			})
		}
		defer src.Close()
	} else {
		src = nil
	}

	req := domain.UpdateProductByIdRequest{
		AccessToken:         accessToken,
		ProductId:           productId,
		BrandId:             c.FormValue("brand_id"),
		ProductCategoryId:   c.FormValue("product_category_id"),
		ProductTypeId:       c.FormValue("product_type_id"),
		ProductImageUrl:     nullableString(c.FormValue("product_image_url")),
		ProductName:         c.FormValue("product_name"),
		Barcode:             c.FormValue("barcode"),
		Warning:             nullableString(c.FormValue("warning")),
		UsageDescription:    nullableString(c.FormValue("usage_description")),
		AmountFluoride:      nullableString(c.FormValue("amount_fluoride")),
		Properties:          nullableString(c.FormValue("properties")),
		ActiveIngredient:    nullableString(c.FormValue("active_ingredient")),
		DangerousIngredient: nullableString(c.FormValue("dangerous_ingredient")),
		IsDangerous:         c.FormValue("is_dangerous"),
		Image:               &src,
	}

	result, err := h.svc.UpdateProductById(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
