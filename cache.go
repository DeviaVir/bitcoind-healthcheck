package main

import (
	"sync"
	"time"
)

type CacheItem struct {
	Result     float64
	Expiration time.Time
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]CacheItem
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
	}
}

func (c *Cache) Set(key string, value float64, duration time.Duration) {
	c.mu.Lock()
	c.items[key] = CacheItem{Result: value, Expiration: time.Now().Add(duration)}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (float64, bool) {
	c.mu.RLock()
	item, exists := c.items[key]
	if !exists || item.Expiration.Before(time.Now()) {
		c.mu.RUnlock()
		return 0, false
	}
	c.mu.RUnlock()
	return item.Result, true
}
