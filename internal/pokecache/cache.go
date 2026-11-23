package pokecache

import "sync"

type Cache struct {
	entries map[string][]byte
	mu      *sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		entries: make(map[string][]byte),
		mu:      &sync.Mutex{},
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return val, true
}
