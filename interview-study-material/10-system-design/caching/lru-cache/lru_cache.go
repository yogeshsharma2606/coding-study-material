package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key   int
	value int
}

func NewLRU(cap int) *LRUCache {
	return &LRUCache{
		capacity: cap,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (lru *LRUCache) Get(key int) int {
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem) // mark as recently used
		return elem.Value.(entry).value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	// If key exists → update & move to front
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		elem.Value = entry{key, value}
		return
	}
	// If capacity reached → remove LRU
	if lru.list.Len() == lru.capacity {
		back := lru.list.Back()
		if back != nil {
			lru.list.Remove(back)
			delete(lru.cache, back.Value.(entry).key)
		}
	}
	// Insert new
	elem := lru.list.PushFront(entry{key, value})
	lru.cache[key] = elem
}

func main() {
	lru := NewLRU(2)

	lru.Put(1, 10)
	lru.Put(2, 20)

	fmt.Println(lru.Get(1)) // 10

	lru.Put(3, 30) // Evicts key 2

	fmt.Println(lru.Get(2)) // -1
	fmt.Println(lru.Get(3)) // 30
}
