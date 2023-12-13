package cache

type CacheInterface interface {
	Store(key string, value string)
	Retrieve(key string) string
	FrequencyOfAccess(key string) int
	CacheMisses() int
	Remove(key string)
}
