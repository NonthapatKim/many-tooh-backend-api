package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (s *service) UpdateProductById(req domain.UpdateProductByIdRequest) (domain.UpdateProductByIdResponse, error) {
	var checkList []domain.CheckExistsRequest

	if req.AccessToken == "" {
		return domain.UpdateProductByIdResponse{}, errors.New("token is required")
	}

	if req.ProductId == "" {
		return domain.UpdateProductByIdResponse{}, errors.New("product_id is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UpdateProductByIdResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.UpdateProductByIdResponse{}, errors.New("invalid token: missing user_id")
	}

	checkList = []domain.CheckExistsRequest{
		{
			Table:  "users",
			Column: "user_id",
			Id:     &userId,
		},
		{
			Table:  "products",
			Column: "product_id",
			Id:     &req.ProductId,
		},
	}

	for _, check := range checkList {
		exists, err := s.repo.CheckExists(check)
		if err != nil {
			return domain.UpdateProductByIdResponse{}, err
		}
		if !exists.Exists {
			return domain.UpdateProductByIdResponse{}, fmt.Errorf("error: %s not found in %s", check.Column, check.Table)
		}
	}

	if req.Image != nil {
		CLOUDINARY_URL := os.Getenv("CLOUDINARY_URL")
		cld, err := cloudinary.NewFromURL(CLOUDINARY_URL)
		if err != nil {
			return domain.UpdateProductByIdResponse{}, errors.New("failed to initialize Cloudinary")
		}

		if req.ProductImageUrl != nil {
			_, err := cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
				PublicID: *req.ProductImageUrl,
			})
			if err != nil {
				return domain.UpdateProductByIdResponse{}, errors.New("failed to delete old image from Cloudinary")
			}
		}

		imageData, err := io.ReadAll(*req.Image)
		if err != nil {
			return domain.UpdateProductByIdResponse{}, errors.New("failed to read image file")
		}

		uploadRes, err := cld.Upload.Upload(context.Background(), bytes.NewReader(imageData), uploader.UploadParams{})
		if err != nil {
			return domain.UpdateProductByIdResponse{}, errors.New("failed to upload image to Cloudinary")
		}

		ImageUrl := uploadRes.SecureURL
		req.ProductImageUrl = &ImageUrl
	}

	_, err = s.repo.UpdateProductById(req)
	if err != nil {
		return domain.UpdateProductByIdResponse{}, err
	}

	response := domain.UpdateProductByIdResponse{
		Code:    200,
		Message: "successfully updated product",
	}

	return response, nil
}
