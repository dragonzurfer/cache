package cache

import (
	"container/list"
	"time"
)

func (c *Cache) Store(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.checkAndUpdateItemIfExists(key, value) {
		return
	}

	if c.evictList.Len() == c.maxSize {
		c.removeOldest()
	}

	c.addNewItem(key, value)
}

func (c *Cache) checkAndUpdateItemIfExists(key string, value string) bool {
	if elem, ok := c.items[key]; ok {
		c.updateExistingItem(elem, value)
		return true
	}
	return false
}

func (c *Cache) updateExistingItem(elem *list.Element, value string) {
	item := elem.Value.(*cacheItem)
	item.value = value
	item.lastAccess = time.Now()
	item.frequency++
	c.evictList.MoveToFront(elem)
}

func (c *Cache) addNewItem(key string, value string) {
	newItem := &cacheItem{
		key:        key,
		value:      value,
		createdAt:  time.Now(),
		lastAccess: time.Now(),
	}
	entry := c.evictList.PushFront(newItem)
	c.items[key] = entry
}
