package service

import (
	"log"

	"github.com/muratovdias/technodom-test-task/internal/app/cache"
	"github.com/muratovdias/technodom-test-task/internal/app/repository"
)

type Client interface {
	Redirect(link string) (int, error)
}

type ClientService struct {
	repo  repository.Clinet
	cache *cache.Cache
}

func NewClientService(repo repository.Clinet, cache *cache.Cache) *ClientService {
	return &ClientService{
		repo:  repo,
		cache: cache,
	}
}

func (c *ClientService) Redirect(link string) (int, error) {
	_, ok := c.cache.Get(link)
	if ok {
		return 301, nil
	} else {
		res, err := c.repo.Redirect(link)
		if err != nil {
			log.Println(err.Error())
			return 0, err
		}
		log.Println("link got from db")
		c.cache.Add(res.HistoryLink, res.ActiveLink)
		return 200, nil
	}
}
