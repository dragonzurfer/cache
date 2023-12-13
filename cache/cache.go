package cache

import (
	"container/list"
	"sync"
	"time"
)

// Cache struct
type Cache struct {
	maxSize     int
	ttl         time.Duration
	evictList   *list.List
	items       map[string]*list.Element
	cacheMisses int
	mu          sync.Mutex
}

type cacheItem struct {
	key        string
	value      string
	createdAt  time.Time
	lastAccess time.Time
	frequency  int
}

// NewCache creates a new Cache
func NewCache(maxSize int, ttl time.Duration) *Cache {
	cache := &Cache{
		maxSize:     maxSize,
		ttl:         ttl,
		evictList:   list.New(),
		items:       make(map[string]*list.Element),
		cacheMisses: 0,
	}

	// Start background eviction routine
	go cache.startEvictionRoutine()

	return cache
}
