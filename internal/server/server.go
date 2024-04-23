package server

import (
	"fmt"
	"log"

	"github.com/chanitt/go-hexagonal-template/internal/adapter/handler"
	"github.com/chanitt/go-hexagonal-template/internal/adapter/repository"
	"github.com/chanitt/go-hexagonal-template/internal/core/service"
	"github.com/chanitt/go-hexagonal-template/internal/router"
	"github.com/joho/godotenv"
)

func RunServer() error {
	fmt.Println("Starting server...")
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// init repository
	repo := repository.New()

	// init service
	svc := service.New(repo)

	// init handler
	hdl := handler.New(svc)

	router, err := router.NewRouter(hdl)
	if err != nil {
		fmt.Println("Error initializing router", err)
		return err
	}

	return router.Start()
}
