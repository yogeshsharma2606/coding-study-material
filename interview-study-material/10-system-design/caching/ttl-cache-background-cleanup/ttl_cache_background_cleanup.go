package main

import (
	"fmt"
	"sync"
	"time"
)

type ItemA struct {
	value    int
	expireAt time.Time
}

type TTLCacheA struct {
	data map[string]ItemA
	mu   sync.Mutex
	stop chan struct{}
}

func NewTTLCacheA(cleanupInterval time.Duration) *TTLCacheA {
	c := &TTLCacheA{
		data: make(map[string]ItemA),
		stop: make(chan struct{}),
	}

	go c.startCleanup(cleanupInterval)
	return c
}

func (c *TTLCacheA) PutA(key string, value int, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = ItemA{
		value:    value,
		expireAt: time.Now().Add(ttl),
	}
}

func (c *TTLCacheA) GetA(key string) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

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

func (c *TTLCacheA) startCleanup(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanup()
		case <-c.stop:
			fmt.Println("channel closed")
			return
		}
	}
}

func (c *TTLCacheA) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, item := range c.data {
		if now.After(item.expireAt) {
			delete(c.data, key)
		}
	}
}

func (c *TTLCacheA) Stop() {
	close(c.stop)
}

func main() {
	cache := NewTTLCacheA(1 * time.Second)

	cache.PutA("a", 100, 2*time.Second)

	time.Sleep(3 * time.Second)

	val, ok := cache.GetA("a")
	fmt.Println(val, ok) // 0 false

	cache.Stop()
}
