package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache	map[string]cacheEntry
	mu		*sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(reapInterval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go c.reapLoop(reapInterval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cutoffTime := time.Now().Add(-interval)
		for key, value := range c.cache {
			if value.createdAt.Before(cutoffTime) {
				c.mu.Lock()
				delete(c.cache, key)
				c.mu.Unlock()
			}
		}
	}
}