package main

import (
	"sync"
	"time"
)

type CacheItem struct {
	Result     int
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

func (c *Cache) Set(key string, value int, duration time.Duration) {
	c.mu.Lock()
	c.items[key] = CacheItem{Result: value, Expiration: time.Now().Add(duration)}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (int, bool) {
	c.mu.RLock()
	item, exists := c.items[key]
	if !exists || item.Expiration.Before(time.Now()) {
		c.mu.RUnlock()
		return 0, false
	}
	c.mu.RUnlock()
	return item.Result, true
}
