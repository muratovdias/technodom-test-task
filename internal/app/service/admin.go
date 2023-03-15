package service

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/muratovdias/technodom-test-task/internal/app/cache"
	"github.com/muratovdias/technodom-test-task/internal/app/models"
	"github.com/muratovdias/technodom-test-task/internal/app/repository"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("this link already exists")
)

type Admin interface {
	GetLinks(offset int) (*[]models.Link, error)
	GetLinkByID(id int) (models.Link, error)
	UpdateLink(id int, newActiveLink string) error
	DeleteLink(id int) error
	CreateLink(newLink string) error
}

type AdminService struct {
	repo  repository.Admin
	cache cache.Cacher
}

func NewAdminService(repo repository.Admin, cache *cache.Cache) *AdminService {
	return &AdminService{
		repo:  repo,
		cache: cache,
	}
}

func (a *AdminService) GetLinks(page int) (*[]models.Link, error) {
	if page != 1 {
		page = (page - 1) * 25
	} else {
		page -= 1
	}
	links, err := a.repo.GetLinks(page)
	if err != nil {
		log.Println("service: " + err.Error())
		return nil, err
	}
	return links, nil
}

func (a *AdminService) GetLinkByID(id int) (models.Link, error) {
	link, err := a.repo.GetLinkByID(id)
	if err != nil {
		log.Println("service: " + err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return link, ErrNotFound
		}
		return link, err
	}
	return link, nil
}

func (a *AdminService) UpdateLink(id int, newActiveLink string) error {
	if err := a.repo.UpdateLink(id, newActiveLink); err != nil {
		log.Println("service: " + err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}
	return nil
}

func (a *AdminService) DeleteLink(id int) error {
	if err := a.repo.DeleteLink(id); err != nil {
		log.Println("service: " + err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}
	return nil
}

func (a *AdminService) CreateLink(newLink string) error {
	if err := a.repo.CreateLink(newLink); err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return ErrAlreadyExists
		}
		return err
	}
	return nil
}
