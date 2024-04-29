package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]interface{}
	mutex    sync.RWMutex
	ttl      map[string]*time.Time
	ttlMutex sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
		ttl:  make(map[string]*time.Time),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value

	// Set TTL
	expiration := time.Now().Add(ttl)
	c.ttlMutex.Lock()
	defer c.ttlMutex.Unlock()
	c.ttl[key] = &expiration
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, found := c.data[key]
	if !found {
		return nil, false
	}

	// Check TTL
	c.ttlMutex.Lock()
	defer c.ttlMutex.Unlock()
	expiration, exists := c.ttl[key]
	if exists && expiration.Before(time.Now()) {
		delete(c.data, key)
		delete(c.ttl, key)
		return nil, false
	}

	return value, true
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)

	// Delete TTL
	c.ttlMutex.Lock()
	defer c.ttlMutex.Unlock()
	delete(c.ttl, key)
}
