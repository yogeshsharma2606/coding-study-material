package cache

import "testing"

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if c.Get(1) != 1 {
		t.Error("get 1")
	}
	c.Put(3, 3) // evicts 2
	if c.Get(2) != -1 {
		t.Error("2 should be evicted")
	}
	c.Put(4, 4) // evicts 1
	if c.Get(1) != -1 {
		t.Error("1 should be evicted")
	}
	if c.Get(3) != 3 || c.Get(4) != 4 {
		t.Error("3 and 4 should remain")
	}
}

func TestLFUCache(t *testing.T) {
	c := NewLFUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if c.Get(1) != 1 {
		t.Error("get 1")
	}
	c.Put(3, 3) // evicts key 2 (freq 1, least recent)
	if c.Get(2) != -1 {
		t.Error("2 should be evicted")
	}
	if c.Get(3) != 3 {
		t.Error("get 3")
	}
	c.Put(4, 4) // evicts key 1 (freq 2) vs 3 (freq 2): 1 is LRU
	if c.Get(1) != -1 {
		t.Error("1 should be evicted")
	}
	if c.Get(3) != 3 || c.Get(4) != 4 {
		t.Error("3 and 4 should remain")
	}
}

func TestRandomizedSet(t *testing.T) {
	s := NewRandomizedSet()
	if !s.Insert(1) {
		t.Error("insert 1")
	}
	if s.Insert(1) {
		t.Error("duplicate insert should fail")
	}
	if s.Remove(2) {
		t.Error("remove absent should fail")
	}
	s.Insert(2)
	s.Insert(3)
	if !s.Remove(1) {
		t.Error("remove 1")
	}
	for i := 0; i < 20; i++ {
		r := s.GetRandom()
		if r != 2 && r != 3 {
			t.Errorf("random gave %d", r)
		}
	}
}

func TestTimeMap(t *testing.T) {
	m := NewTimeMap()
	m.Set("foo", "bar", 1)
	if m.Get("foo", 1) != "bar" {
		t.Error("t=1")
	}
	if m.Get("foo", 3) != "bar" {
		t.Error("t=3 should still be bar")
	}
	m.Set("foo", "bar2", 4)
	if m.Get("foo", 4) != "bar2" {
		t.Error("t=4")
	}
	if m.Get("foo", 5) != "bar2" {
		t.Error("t=5")
	}
	if m.Get("foo", 0) != "" {
		t.Error("t=0 should be empty")
	}
}

func TestRateLimiter(t *testing.T) {
	// capacity 3, 1 token per second
	r := NewRateLimiter(3, 1)
	// 3 immediate requests at t=0 allowed, 4th denied
	for i := 0; i < 3; i++ {
		if !r.Allow(0) {
			t.Errorf("request %d should be allowed", i)
		}
	}
	if r.Allow(0) {
		t.Error("4th request should be denied")
	}
	// after 1s, one token refilled
	if !r.Allow(1000) {
		t.Error("should be allowed after refill")
	}
	if r.Allow(1000) {
		t.Error("no more tokens")
	}
}

func TestMyHashMap(t *testing.T) {
	m := NewMyHashMap()
	m.Put(1, 1)
	m.Put(2, 2)
	if m.Get(1) != 1 {
		t.Error("get 1")
	}
	if m.Get(3) != -1 {
		t.Error("3 absent")
	}
	m.Put(2, 1) // overwrite
	if m.Get(2) != 1 {
		t.Error("overwrite 2")
	}
	m.Remove(2)
	if m.Get(2) != -1 {
		t.Error("2 removed")
	}
	// collision case: 1 and 1+769 hash to the same bucket
	m.Put(1+769, 99)
	if m.Get(1) != 1 || m.Get(1+769) != 99 {
		t.Error("collision handling failed")
	}
}
