package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	mux      *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache:    make(map[string]cacheEntry),
		mux:      &sync.Mutex{},
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		// Collect entries ready to be deleted
		toDelete := make([]string, 0)
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				toDelete = append(toDelete, key)
			}
		}
		// Delete old entries
		for _, key := range toDelete {
			c.mux.Lock()
			delete(c.cache, key)
			c.mux.Unlock()
		}
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, ok := c.cache[key]
	c.mux.Unlock()
	return entry.val, ok
}
