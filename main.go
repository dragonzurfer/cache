// Cache
// Store(key, value)
// Retrieve(key)
// FrequencyOfAccess(key)
// CacheMisses()
// Remove(key)

// Constraints
// limited size
// eviction plan ( LRU / LFU )
// TTL ( keys ) - lazy / "active"

package main

import (
	"fmt"
	"time"

	CACHE "github.com/dragonzurfer/cache/cache"
)

func demonstrateCacheBehavior() {
	cacheSize := 5
	ttl := 5 * time.Minute
	cache := CACHE.NewCache(cacheSize, ttl)

	// Add 3 key-value pairs
	for i := 1; i <= 3; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		cache.Store(key, value)
	}

	// Access and print stats for the first 3 keys
	for i := 1; i <= 3; i++ {
		key := fmt.Sprintf("key%d", i)
		fmt.Println("Retrieved:", cache.Retrieve(key))
		fmt.Println("Frequency of Access for", key, ":", cache.FrequencyOfAccess(key))
	}
	fmt.Println("Cache Misses:", cache.CacheMisses())

	// Add 2 more key-value pairs
	for i := 4; i <= 5; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		cache.Store(key, value)
	}

	// Print stats again
	for i := 1; i <= 5; i++ {
		key := fmt.Sprintf("key%d", i)
		fmt.Println("Frequency of Access for", key, ":", cache.FrequencyOfAccess(key))
	}
	fmt.Println("Cache Misses:", cache.CacheMisses())

	// Add another key, causing eviction of key1 (LRU)
	cache.Store("key6", "value6")

	// Access the evicted key
	fmt.Println("Retrieved evicted key (key1):", cache.Retrieve("key1"))

	// Print all stats
	for i := 1; i <= 6; i++ {
		key := fmt.Sprintf("key%d", i)
		fmt.Println("Frequency of Access for", key, ":", cache.FrequencyOfAccess(key))
	}

	// Miss should be one
	fmt.Println("Cache Misses:", cache.CacheMisses())
}

func demonstrateTTLBehavior() {
	cacheSize := 5
	ttl := 10 * time.Second // Short TTL for demonstration
	cache := CACHE.NewCache(cacheSize, ttl)

	// Add key-value pairs
	for i := 1; i <= 3; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		cache.Store(key, value)
		fmt.Println("Stored:", key, value)
	}

	// Immediately retrieve keys
	for i := 1; i <= 3; i++ {
		key := fmt.Sprintf("key%d", i)
		value := cache.Retrieve(key)
		if value == "" {
			fmt.Println("Key", key, "not found (unexpected at this point)")
		} else {
			fmt.Println("Retrieved immediately:", key, value)
		}
	}

	// Wait for longer than the TTL
	fmt.Println("Waiting for", ttl+time.Second, "to allow TTL to expire...")
	time.Sleep(ttl + time.Second)

	// Attempt to retrieve the keys again
	for i := 1; i <= 3; i++ {
		key := fmt.Sprintf("key%d", i)
		value := cache.Retrieve(key)
		if value == "" {
			fmt.Println("Key", key, "not found (expected after TTL expiration)")
		} else {
			fmt.Println("Retrieved after TTL:", key, value)
		}
	}
}

func main() {
	demonstrateCacheBehavior()
	demonstrateTTLBehavior()
}
