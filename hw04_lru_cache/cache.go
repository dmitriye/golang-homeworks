package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.Mutex

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()

	item, exists := c.items[key]

	if exists {
		item.Value.(*cacheItem).value = value
		c.queue.MoveToFront(item)
	} else {
		if c.queue.Len() == c.capacity {
			lastItem := c.queue.Back()
			delete(c.items, lastItem.Value.(*cacheItem).key)
			c.queue.Remove(lastItem)
		}

		item = c.queue.PushFront(&cacheItem{key: key, value: value})
		c.items[key] = item
	}

	return exists
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	if item, exists := c.items[key]; exists {
		c.queue.MoveToFront(item)
		return item.Value.(*cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.Lock()
	defer c.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
