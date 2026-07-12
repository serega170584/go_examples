package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	m  map[string]CacheItem
	mu sync.RWMutex
}

type CacheItem struct {
	value    string
	deadline time.Time
}

func NewCache() *Cache {
	return &Cache{
		m: make(map[string]CacheItem),
	}
}

func (cache *Cache) SetValue(key, value string, ttl time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cacheItem := CacheItem{
		value:    value,
		deadline: time.Now().Add(ttl),
	}

	cache.m[key] = cacheItem
}

func (cache *Cache) GetValue(key string) *string {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	if _, ok := cache.m[key]; !ok {
		return nil
	}

	if cache.m[key].deadline.Before(time.Now()) {
		return nil
	}

	mv := cache.m[key]

	return &mv.value
}

func main() {
	cache := NewCache()
	cache.SetValue("123", "456", time.Second)
	v := cache.GetValue("123")
	fmt.Println(*v)

	time.Sleep(2 * time.Second)
	v = cache.GetValue("123")
	fmt.Println(v)
}
