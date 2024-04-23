package service

import "github.com/chanitt/go-hexagonal-template/internal/core/port"

type service struct {
	repo port.Repository
}

func New(repo port.Repository) port.Service {
	return &service{repo: repo}
}
