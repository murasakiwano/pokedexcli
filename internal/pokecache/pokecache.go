package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mux  *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data: map[string]cacheEntry{},
		mux:  &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	createdAt := time.Now().UTC()
	entry := cacheEntry{createdAt: createdAt, val: val}
	c.data[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, found := c.data[key]

	return entry.val, found
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, since time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-since)) {
			delete(c.data, k)
		}
	}
}
