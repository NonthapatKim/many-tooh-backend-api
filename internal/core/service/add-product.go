package service

import (
	"context"
	"errors"
	"os"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/function"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (s *service) AddProduct(req domain.AddProductRequest) (domain.AddProductResponse, error) {
	if req.AccessToken == "" {
		return domain.AddProductResponse{}, errors.New("token is required")
	}

	claims, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.AddProductResponse{}, err
	}

	userId, ok := claims["user_id"].(string)
	if !ok || userId == "" {
		return domain.AddProductResponse{}, errors.New("invalid token: missing user_id")
	}

	reqUserExists := domain.CheckExistsRequest{
		Table:  "users",
		Column: "user_id",
		Id:     &userId,
	}

	exists, err := s.repo.CheckExists(reqUserExists)
	if err != nil {
		return domain.AddProductResponse{}, err
	}
	if !exists.Exists {
		return domain.AddProductResponse{}, errors.New("error: user not found")
	}

	CLOUDINARY_URL := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(CLOUDINARY_URL)
	if err != nil {
		return domain.AddProductResponse{}, errors.New("failed to initialize Cloudinary")
	}

	uploadRes, err := cld.Upload.Upload(context.Background(), req.Image, uploader.UploadParams{})
	if err != nil {
		return domain.AddProductResponse{}, errors.New("failed to upload image to Cloudinary")
	}

	req.ImageUrl = uploadRes.SecureURL

	_, err = s.repo.AddProduct(req)
	if err != nil {
		return domain.AddProductResponse{}, err
	}

	response := domain.AddProductResponse{
		Code:    200,
		Message: "successfully created product",
	}

	return response, nil
}
