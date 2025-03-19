package service

import "github.com/NonthapatKim/many_tooh_backend_api/internal/core/port"

type service struct {
	repo port.Repository
}

func New(repo port.Repository) port.Service {
	return &service{repo: repo}
}
