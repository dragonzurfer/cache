package cache

import (
	"container/list"
	"time"
)

// Retrieve gets the value of a key from the cache
func (c *Cache) Retrieve(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		if c.isExpired(elem) {
			c.removeElement(elem)
			c.cacheMisses++
			return ""
		}
		c.updateAccessDetails(elem)
		return elem.Value.(*cacheItem).value
	}

	c.cacheMisses++
	return ""
}

func (c *Cache) isExpired(elem *list.Element) bool {
	item := elem.Value.(*cacheItem)
	return time.Since(item.createdAt) > c.ttl
}

func (c *Cache) updateAccessDetails(elem *list.Element) {
	item := elem.Value.(*cacheItem)
	item.frequency++
	item.lastAccess = time.Now()
	c.evictList.MoveToFront(elem)
}
