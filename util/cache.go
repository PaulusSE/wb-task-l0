package util

import (
	"errors"
	"sync"

	"github.com/PaulusSE/wbL0/db"
)

type Cache struct {
	sync.RWMutex
	items map[string]interface{}
}

func New() *Cache {
	items := make(map[string]interface{})
	cache := Cache{
		items: items,
	}
	return &cache
}

func (c *Cache) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.items[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	item, found := c.items[key]
	if !found {
		return db.Order{}, false
	}
	return item, true
}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("Key not found")
	}
	delete(c.items, key)
	return nil
}
