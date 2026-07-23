package concurrency

import (
	"context"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5}
	got := WorkerPool(jobs, 3, func(x int) int { return x * x })
	want := []int{1, 4, 9, 16, 25}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %d want %d", i, got[i], want[i])
		}
	}
}

func TestFanIn(t *testing.T) {
	a := Generate(1, 2, 3)
	b := Generate(4, 5, 6)
	merged := FanIn(a, b)
	var got []int
	for v := range merged {
		got = append(got, v)
	}
	sort.Ints(got)
	want := []int{1, 2, 3, 4, 5, 6}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got %v want %v", got, want)
			break
		}
	}
}

func TestParallelMap(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	got := ParallelMap(in, 2, func(x int) int { return x + 10 })
	want := []int{11, 12, 13, 14, 15, 16}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %d want %d", i, got[i], want[i])
		}
	}
}

func TestParallelMapConcurrencyBound(t *testing.T) {
	// verify no more than `workers` run simultaneously
	var mu sync.Mutex
	current, peak := 0, 0
	ParallelMap([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, func(x int) int {
		mu.Lock()
		current++
		if current > peak {
			peak = current
		}
		mu.Unlock()
		time.Sleep(5 * time.Millisecond)
		mu.Lock()
		current--
		mu.Unlock()
		return x
	})
	if peak > 3 {
		t.Errorf("peak concurrency %d exceeded limit 3", peak)
	}
}

func TestPipeline(t *testing.T) {
	got := Pipeline([]int{1, 2, 3, 4})
	want := []int{1, 4, 9, 16}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got %v want %v", got, want)
			break
		}
	}
}

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	if c.Value() != 1000 {
		t.Errorf("got %d want 1000", c.Value())
	}
}

func TestAtomicCounter(t *testing.T) {
	c := &AtomicCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	if c.Value() != 1000 {
		t.Errorf("got %d want 1000", c.Value())
	}
}

func TestFirstResult(t *testing.T) {
	tasks := []func(context.Context) int{
		func(ctx context.Context) int {
			select {
			case <-time.After(50 * time.Millisecond):
				return 1
			case <-ctx.Done():
				return -1
			}
		},
		func(ctx context.Context) int {
			select {
			case <-time.After(10 * time.Millisecond):
				return 2 // fastest
			case <-ctx.Done():
				return -1
			}
		},
	}
	if got := FirstResult(context.Background(), tasks); got != 2 {
		t.Errorf("got %d want 2 (fastest)", got)
	}
}
