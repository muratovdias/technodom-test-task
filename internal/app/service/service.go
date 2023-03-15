package service

import (
	"github.com/muratovdias/technodom-test-task/internal/app/cache"
	"github.com/muratovdias/technodom-test-task/internal/app/repository"
)

type Service struct {
	Admin
	Client
}

func NewService(repo *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Admin:  NewAdminService(repo.Admin, cache),
		Client: NewClientService(repo.Clinet, cache),
	}
}
