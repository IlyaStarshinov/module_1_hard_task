package cache

import (
	"sync"
)

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Не меняйте названия структуры и название метода создания экземпляра Cache, иначе не будут проходить тесты

type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewCache создаёт и возвращает новый экземпляр Cache.
func NewCache() Interface {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[k] = v
}

func (c *Cache) Get(k string) (v string, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok = c.data[k]
	return v, ok

}
