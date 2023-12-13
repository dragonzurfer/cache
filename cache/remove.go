package cache

import "container/list"

// Remove removes a key from the cache
func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.removeElement(elem)
	}
}

// Helper methods
func (c *Cache) removeOldest() {
	elem := c.evictList.Back()
	if elem != nil {
		c.removeElement(elem)
	}
}

func (c *Cache) removeElement(elem *list.Element) {
	c.evictList.Remove(elem)
	item := elem.Value.(*cacheItem)
	delete(c.items, item.key)
}
