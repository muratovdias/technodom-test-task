package service

import (
	"github.com/muratovdias/technodom-test-task/internal/app/cache"
	"github.com/muratovdias/technodom-test-task/internal/app/repository"
)

type Service struct {
	Admin
}

func NewService(repo *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Admin: NewAdminService(repo.Admin, cache),
	}
}
