// Package concurrency contains Go concurrency interview problems.
//
// Go's motto: "Don't communicate by sharing memory; share memory by
// communicating." Goroutines are cheap threads; channels pass data and
// synchronize. The core patterns: WORKER POOL (bounded parallelism over a job
// queue), FAN-IN / FAN-OUT (merge/split channels), PIPELINE (stages connected by
// channels), a SEMAPHORE (buffered channel) to cap concurrency, and MUTEX/ATOMIC
// for shared state. Always ask: who closes the channel, and can this deadlock or
// race? Run tests with `go test -race`.
package concurrency

import (
	"context"
	"sync"
	"sync/atomic"
)

// WorkerPool processes jobs with a fixed number of workers and returns results
// in the SAME ORDER as the input. Workers pull indices from a channel; each
// writes to its own result slot, so there's no data race and no lock needed.
func WorkerPool(jobs []int, numWorkers int, work func(int) int) []int {
	n := len(jobs)
	results := make([]int, n)
	idxCh := make(chan int)
	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range idxCh { // ranges until idxCh is closed
				results[i] = work(jobs[i])
			}
		}()
	}
	for i := 0; i < n; i++ {
		idxCh <- i
	}
	close(idxCh) // signal workers no more jobs; range loops exit
	wg.Wait()
	return results
}

// FanIn merges several input channels into one output channel. A goroutine per
// input forwards values; a closer goroutine waits for all of them, then closes
// the output (the standard "who closes" answer for fan-in).
func FanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// ParallelMap applies f to every element with at most `workers` running at once,
// using a buffered channel as a SEMAPHORE. Results preserve input order.
func ParallelMap(input []int, workers int, f func(int) int) []int {
	out := make([]int, len(input))
	sem := make(chan struct{}, workers) // capacity = max concurrency
	var wg sync.WaitGroup
	for i, v := range input {
		wg.Add(1)
		sem <- struct{}{} // acquire a slot (blocks if `workers` are busy)
		go func(i, v int) {
			defer wg.Done()
			defer func() { <-sem }() // release the slot
			out[i] = f(v)
		}(i, v)
	}
	wg.Wait()
	return out
}

// --- Pipeline: stages connected by channels, each stage a goroutine ---

// Generate emits the given numbers on a channel, then closes it.
func Generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// Square reads a channel and emits squares; it closes its output when its input
// closes (each stage owns and closes the channel it produces).
func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// Pipeline wires Generate -> Square and collects the results.
func Pipeline(nums []int) []int {
	var res []int
	for v := range Square(Generate(nums...)) {
		res = append(res, v)
	}
	return res
}

// SafeCounter is a concurrency-safe counter guarded by a mutex.
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// AtomicCounter is a lock-free counter using atomic operations - cheaper than a
// mutex for a single integer.
type AtomicCounter struct {
	count int64
}

func (c *AtomicCounter) Inc()       { atomic.AddInt64(&c.count, 1) }
func (c *AtomicCounter) Value() int { return int(atomic.LoadInt64(&c.count)) }

// FirstResult runs all tasks concurrently and returns the first result, then
// cancels the rest via context. The buffered channel prevents the slower
// goroutines from leaking (they can always send and exit).
func FirstResult(ctx context.Context, tasks []func(context.Context) int) int {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // cancel losers when we return

	resCh := make(chan int, len(tasks)) // buffered so no goroutine blocks forever
	for _, task := range tasks {
		go func(t func(context.Context) int) {
			resCh <- t(ctx)
		}(task)
	}
	return <-resCh
}
