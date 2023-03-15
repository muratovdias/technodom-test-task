package cache

import "sync"

type Cacher interface {
	Add(key, value string)
	Get(key string) (value string, ok bool)
	Len() int
}

type Cache struct {
	cache   map[string]string
	maxSize int
	mutex   *sync.RWMutex
}

func NewCache(maxSize int) *Cache {
	return &Cache{
		cache:   make(map[string]string, 500),
		maxSize: maxSize,
		mutex:   &sync.RWMutex{},
	}
}

func (c *Cache) Add(key, value string) {}
func (c *Cache) Get(key string) (value string, ok bool) {
	return "", false
}

func (c *Cache) Len() int {
	return 0
}
