package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data 			map[string]cacheEntry
	mu  			*sync.Mutex
}

type cacheEntry struct {
	createdAt 	time.Time
	val 	[]byte
}

func NewCache(interval time.Duration) Cache{
	c := Cache{
		data: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, content []byte){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: content,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.data[key]
	if !ok {return nil, ok}
	return val.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, val := range c.data {
			if val.createdAt.Before(time.Now().Add(-interval)){
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}