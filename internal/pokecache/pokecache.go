package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mutex    sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) *Cache {
	new := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go new.reapLoop()
	return new
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	if time.Since(entry.createdAt) > c.interval {
		delete(c.entries, key)
		return nil, false
	}
	return entry.value, true
}

func (c *Cache) reapLoop() {
	for {
		time.Sleep(c.interval)
		c.mutex.Lock()
		for key := range c.entries {
			if time.Since(c.entries[key].createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}
