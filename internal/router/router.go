package router

import (
	"fmt"

	"github.com/chanitt/go-hexagonal-template/internal/adapter/handler"
)

type Router struct {
}

const serviceBaseURL = "/api"

func NewRouter(h handler.Handler) (*Router, error) {

	return &Router{}, nil
}

func (r *Router) Start() error {
	fmt.Println("Listening on port 8080")
	// return r.Listen(":8080")
	return nil
}
