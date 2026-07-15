package cache

import (
	"sync"
	"time"

	"github.com/fandrien/book-cabin/model"
)

type cacheItem struct {
	Response  *model.SearchResponse
	ExpiredAt time.Time
}

type MemoryCache struct {
	mu   sync.RWMutex
	data map[string]cacheItem
	ttl  time.Duration
}

func NewMemoryCache(ttl time.Duration) *MemoryCache {
	return &MemoryCache{
		data: make(map[string]cacheItem),
		ttl:  ttl,
	}
}

func (c *MemoryCache) Get(key string) (*model.SearchResponse, bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.data[key]
	if !ok {
		return nil, false
	}

	if time.Now().After(item.ExpiredAt) {
		delete(c.data, key)
		return nil, false
	}

	return item.Response, true
}

func (c *MemoryCache) Set(
	key string,
	value *model.SearchResponse,
) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheItem{
		Response:  value,
		ExpiredAt: time.Now().Add(c.ttl),
	}
}
