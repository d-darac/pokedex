package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type ReadWrite interface {
	Add(key string, val []byte)
	Get(key string) ([]byte, bool)
}

func NewCache(interval time.Duration) PokeCache {
	cache := PokeCache{
		data: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

func (c *PokeCache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *PokeCache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.data[key]
	return entry.val, exists

}

func (c *PokeCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}

}

func (c *PokeCache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.data {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.data, key)
		}
	}
}
