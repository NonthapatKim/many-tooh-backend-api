package repository

import "github.com/chanitt/go-hexagonal-template/internal/core/port"

type repository struct {
}

func New() port.Repository {
	return &repository{}
}
