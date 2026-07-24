package main

import (
	"fmt"
	"time"
)

type Item struct {
	value    int
	expireAt time.Time
}

type TTLCache struct {
	data map[string]Item
}

func NewTTLCache() *TTLCache {
	return &TTLCache{
		data: make(map[string]Item),
	}
}

func (c *TTLCache) Put(key string, value int, ttl time.Duration) {
	c.data[key] = Item{
		value:    value,
		expireAt: time.Now().Add(ttl),
	}
}

func (c *TTLCache) Get(key string) (int, bool) {
	item, ok := c.data[key]
	if !ok {
		return 0, false
	}

	if time.Now().After(item.expireAt) {
		delete(c.data, key)
		return 0, false
	}

	return item.value, true
}
func main() {
	cache := NewTTLCache()

	cache.Put("a", 100, 2*time.Second)

	val, ok := cache.Get("a")
	fmt.Println(val, ok) // 100 true

	time.Sleep(3 * time.Second)

	val, ok = cache.Get("a")
	fmt.Println(val, ok) // 0 false
}
