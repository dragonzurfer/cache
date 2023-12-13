package cache

import "time"

// startEvictionRoutine runs a background routine for active eviction
func (c *Cache) startEvictionRoutine() {
	ticker := time.NewTicker(time.Minute) // Run every minute, adjust as needed
	for range ticker.C {
		c.evictExpiredItems()
	}
}

// evictExpiredItems evicts items that have exceeded their TTL
func (c *Cache) evictExpiredItems() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for e := c.evictList.Back(); e != nil; e = e.Prev() {
		if c.isExpired(e) {
			c.removeElement(e)
		}
	}
}
