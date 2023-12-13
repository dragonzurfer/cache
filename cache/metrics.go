package cache

// FrequencyOfAccess returns the number of times a key has been accessed
func (c *Cache) FrequencyOfAccess(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		return elem.Value.(*cacheItem).frequency
	}
	return 0
}

// CacheMisses returns the number of cache misses
func (c *Cache) CacheMisses() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.cacheMisses
}
