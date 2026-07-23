// Package cache contains data-structure DESIGN problems - the kind senior
// interviews favor because they combine two structures to hit O(1) guarantees.
//
// The recurring idea: a hash map gives O(1) lookup, but maps have no order; pair
// the map with an ordering structure (doubly linked list for recency, frequency
// buckets for LFU, a slice for random access) so BOTH the lookup and the ordering
// operation are O(1). Always state the invariant that keeps the two in sync.
package cache

import (
	"container/list"
	"math/rand"
	"sort"
)

// --- 1. LRU Cache: hash map + doubly linked list ---
// The list keeps usage order (front = most recent). The map gives O(1) access to
// each node so we can move/evict in O(1).

type lruEntry struct {
	key, value int
}

type LRUCache struct {
	capacity int
	ll       *list.List            // front = most recently used
	items    map[int]*list.Element // key -> list node
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		ll:       list.New(),
		items:    make(map[int]*list.Element),
	}
}

func (c *LRUCache) Get(key int) int {
	if el, ok := c.items[key]; ok {
		c.ll.MoveToFront(el) // mark most recently used
		return el.Value.(*lruEntry).value
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if el, ok := c.items[key]; ok {
		el.Value.(*lruEntry).value = value
		c.ll.MoveToFront(el)
		return
	}
	if c.ll.Len() == c.capacity {
		back := c.ll.Back() // least recently used
		if back != nil {
			c.ll.Remove(back)
			delete(c.items, back.Value.(*lruEntry).key)
		}
	}
	el := c.ll.PushFront(&lruEntry{key, value})
	c.items[key] = el
}

// --- 2. LFU Cache: hash maps + per-frequency doubly linked lists ---
// Evict the least frequently used; break ties by least recently used. Track each
// key's frequency and keep one list per frequency, plus the current minFreq.

type lfuEntry struct {
	key, value, freq int
}

type LFUCache struct {
	capacity int
	minFreq  int
	items    map[int]*list.Element // key -> node
	freqs    map[int]*list.List    // frequency -> list of nodes (front = most recent)
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		freqs:    make(map[int]*list.List),
	}
}

func (c *LFUCache) bump(el *list.Element) {
	e := el.Value.(*lfuEntry)
	c.freqs[e.freq].Remove(el)
	if c.freqs[e.freq].Len() == 0 {
		delete(c.freqs, e.freq)
		if c.minFreq == e.freq {
			c.minFreq++
		}
	}
	e.freq++
	if c.freqs[e.freq] == nil {
		c.freqs[e.freq] = list.New()
	}
	c.items[e.key] = c.freqs[e.freq].PushFront(e)
}

func (c *LFUCache) Get(key int) int {
	if el, ok := c.items[key]; ok {
		v := el.Value.(*lfuEntry).value
		c.bump(el)
		return v
	}
	return -1
}

func (c *LFUCache) Put(key, value int) {
	if c.capacity == 0 {
		return
	}
	if el, ok := c.items[key]; ok {
		el.Value.(*lfuEntry).value = value
		c.bump(el)
		return
	}
	if len(c.items) == c.capacity {
		lru := c.freqs[c.minFreq].Back() // least frequent, least recent
		c.freqs[c.minFreq].Remove(lru)
		delete(c.items, lru.Value.(*lfuEntry).key)
	}
	e := &lfuEntry{key, value, 1}
	if c.freqs[1] == nil {
		c.freqs[1] = list.New()
	}
	c.items[key] = c.freqs[1].PushFront(e)
	c.minFreq = 1
}

// --- 3. RandomizedSet: insert/remove/getRandom all O(1) ---
// A slice gives O(1) random access; a map gives O(1) location. To remove in O(1),
// swap the target with the last element, then shrink.

type RandomizedSet struct {
	vals []int
	idx  map[int]int // value -> index in vals
}

func NewRandomizedSet() *RandomizedSet {
	return &RandomizedSet{idx: make(map[int]int)}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.idx[val]; ok {
		return false
	}
	s.idx[val] = len(s.vals)
	s.vals = append(s.vals, val)
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	i, ok := s.idx[val]
	if !ok {
		return false
	}
	last := len(s.vals) - 1
	s.vals[i] = s.vals[last] // move last into the hole
	s.idx[s.vals[i]] = i
	s.vals = s.vals[:last]
	delete(s.idx, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	return s.vals[rand.Intn(len(s.vals))]
}

// --- 4. Time-Based Key-Value Store ---
// Set(key, value, timestamp); Get(key, t) returns the value with the largest
// timestamp <= t. Timestamps are increasing per key, so binary-search them.

type tvPair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	store map[string][]tvPair
}

func NewTimeMap() *TimeMap { return &TimeMap{store: make(map[string][]tvPair)} }

func (m *TimeMap) Set(key, value string, timestamp int) {
	m.store[key] = append(m.store[key], tvPair{timestamp, value})
}

func (m *TimeMap) Get(key string, timestamp int) string {
	pairs := m.store[key]
	// largest index with timestamp <= given
	i := sort.Search(len(pairs), func(i int) bool { return pairs[i].timestamp > timestamp })
	if i == 0 {
		return ""
	}
	return pairs[i-1].value
}

// --- 5. Token-Bucket Rate Limiter ---
// Tokens refill at a fixed rate up to a capacity; each request consumes one.
// Time is injected (nowMs) so it is deterministic and testable.

type RateLimiter struct {
	capacity    float64
	refillPerMs float64 // tokens added per millisecond
	tokens      float64
	lastMs      int64
}

func NewRateLimiter(capacity int, refillPerSecond float64) *RateLimiter {
	return &RateLimiter{
		capacity:    float64(capacity),
		refillPerMs: refillPerSecond / 1000.0,
		tokens:      float64(capacity),
		lastMs:      0,
	}
}

// Allow reports whether a request at time nowMs may proceed, consuming a token.
func (r *RateLimiter) Allow(nowMs int64) bool {
	elapsed := float64(nowMs - r.lastMs)
	r.tokens = minF(r.capacity, r.tokens+elapsed*r.refillPerMs)
	r.lastMs = nowMs
	if r.tokens >= 1 {
		r.tokens--
		return true
	}
	return false
}

func minF(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// --- 6. Design HashMap (open addressing via buckets with chaining) ---

type MyHashMap struct {
	buckets [][]kv
	size    int
}

type kv struct{ key, value int }

func NewMyHashMap() *MyHashMap {
	const buckets = 769 // a prime reduces collisions
	return &MyHashMap{buckets: make([][]kv, buckets)}
}

func (m *MyHashMap) hash(key int) int { return key % len(m.buckets) }

func (m *MyHashMap) Put(key, value int) {
	b := m.hash(key)
	for i := range m.buckets[b] {
		if m.buckets[b][i].key == key {
			m.buckets[b][i].value = value
			return
		}
	}
	m.buckets[b] = append(m.buckets[b], kv{key, value})
}

func (m *MyHashMap) Get(key int) int {
	b := m.hash(key)
	for _, e := range m.buckets[b] {
		if e.key == key {
			return e.value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	b := m.hash(key)
	for i, e := range m.buckets[b] {
		if e.key == key {
			m.buckets[b] = append(m.buckets[b][:i], m.buckets[b][i+1:]...)
			return
		}
	}
}
