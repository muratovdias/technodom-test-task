package cache

import (
	"log"
	"sync"
	"time"

	"github.com/muratovdias/technodom-test-task/internal/app/models"
)

type Cacher interface {
	Add(key, value string)
	Get(key string) (value cacheValue, ok bool)
	Len() int
}

type cacheValue struct {
	ActiveLink string
	expire     time.Time
}

type Cache struct {
	cache   map[string]cacheValue
	maxSize int
	mutex   *sync.RWMutex
}

func NewCache(maxSize int) *Cache {
	return &Cache{
		cache:   make(map[string]cacheValue, 500),
		maxSize: maxSize,
		mutex:   &sync.RWMutex{},
	}
}

func (c *Cache) FillCache(links []models.Link, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("cache fill")
	for _, link := range links {
		value := cacheValue{
			ActiveLink: link.ActiveLink,
			expire:     time.Now().Add(5 * time.Minute),
		}
		c.cache[link.HistoryLink] = value
	}
}

func (c *Cache) Add(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if len(c.cache) < 1001 {
		newValue := cacheValue{
			ActiveLink: value,
			expire:     time.Now().Add(5 * time.Minute),
		}
		if _, ok := c.cache[key]; !ok {
			c.cache[key] = newValue
			log.Println("link added to cache")
		}
	} else {
		log.Println("cache is full")
	}
}

func (c *Cache) Get(key string) (cacheValue, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, ok := c.cache[key]
	if !ok {
		return cacheValue{}, false
	}
	if time.Now().After(value.expire) {
		delete(c.cache, key)
		return cacheValue{}, false
	}
	log.Println("link got from cache")
	return value, ok
}

func (c *Cache) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.cache)
}
