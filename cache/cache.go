package cache

import (
	"container/list"
	"sync"
)

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Не меняйте названия структуры и название метода создания экземпляра Cache, иначе не будут проходить тесты

type Cache struct {
	capacity int // Максимальное количество элементов
	data     map[string]*list.Element
	list     *list.List // Двусвязный список для реализации вытеснения LRU
	mu       sync.RWMutex
}

type cacheEntry struct {
	key   string
	value string
}

// NewCache создаёт и возвращает новый экземпляр Cache.
func NewCache(capacity int) Interface {
	return &Cache{
		capacity: capacity,
		data:     make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Если ключ уже существует - обновляем значение и перемещаем в начало
	if elem, ok := c.data[k]; ok {
		c.list.MoveToFront(elem)
		elem.Value.(*cacheEntry).value = v
		return
	}

	// Добавляем новый элемент
	elem := c.list.PushFront(&cacheEntry{k, v})
	c.data[k] = elem

	// Если превысили capacity - удаляем самый старый элемент
	if c.list.Len() > c.capacity {
		oldest := c.list.Back()
		if oldest != nil {
			delete(c.data, oldest.Value.(*cacheEntry).key)
			c.list.Remove(oldest)
		}
	}
}

func (c *Cache) Get(k string) (v string, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if elem, ok := c.data[k]; ok {
		// Перемещаем элемент в начало (последний использованный)
		c.list.MoveToFront(elem)
		return elem.Value.(*cacheEntry).value, true
	}

	return "", false

}
