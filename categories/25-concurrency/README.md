# 25 - Concurrency (Go)

> Code: [`concurrency.go`](concurrency.go) - Tests: [`concurrency_test.go`](concurrency_test.go) - run `go test -race ./categories/25-concurrency/`

## Overview & mental model

Go's concurrency model: **goroutines** are cheap, independently-scheduled functions (`go f()`); **channels** pass values between them and synchronize. The guiding principle:

> Don't communicate by sharing memory; share memory by communicating.

The patterns that cover most interviews:

- **Worker pool** - a fixed set of goroutines draining a job channel (bounded parallelism).
- **Fan-out / fan-in** - split work across goroutines, then merge their outputs into one channel.
- **Pipeline** - stages connected by channels; each stage owns and closes the channel it produces.
- **Semaphore** - a buffered channel used to cap concurrency.
- **Mutex / atomic** - protect shared mutable state when channels aren't the natural fit.
- **Context** - propagate cancellation/deadlines across goroutines.

Two questions to ask about any concurrent design: **who closes the channel?** and **can this deadlock or race?**

## How to recognize it

- "Process N tasks in parallel with a limit", "merge results from multiple sources".
- "Rate limit / bound concurrency", "cancel remaining work once one finishes".
- "Make this counter/map thread-safe", "avoid the data race".

## How to think / attack plan

1. Model the dataflow: sources -> processing -> sink. Channels connect stages.
2. Decide ownership: the producer of a channel closes it; consumers `range` until closed.
3. Bound concurrency with a worker pool or a buffered-channel semaphore.
4. Prevent leaks: ensure every goroutine can finish (buffered result channels, context cancellation).
5. For shared state, prefer channels; use a mutex/atomic when the state is a simple counter/map. **Always test with `-race`.**

## Core idioms

```go
// WaitGroup: wait for a set of goroutines
var wg sync.WaitGroup
wg.Add(1); go func(){ defer wg.Done(); /* work */ }(); wg.Wait()

// Semaphore via buffered channel (cap = max concurrency)
sem := make(chan struct{}, n)
sem <- struct{}{}       // acquire
/* work */
<-sem                    // release

// select: wait on multiple channels / cancellation
select {
case v := <-ch:      use(v)
case <-ctx.Done():   return ctx.Err()
}
```

---

## Problems

### 1. Worker Pool (bounded parallelism)
**Direction of thinking.** Spawning one goroutine per job is unbounded; instead spawn `numWorkers` goroutines that pull job indices from a channel. Each worker writes to its **own result slot** (distinct index), so there's no shared write and no lock. Close the job channel to signal completion; a `WaitGroup` waits for all workers.
**Pitfalls.** Forgetting to close the job channel -> workers block forever (deadlock).
**Complexity.** Work is O(n * cost) wall-clock O(n/numWorkers * cost).

### 2. Fan-In (merge channels)
**Direction of thinking.** One forwarding goroutine per input channel copies values to a shared output. A separate goroutine `wg.Wait()`s for all forwarders, then **closes the output** - this is the canonical answer to "who closes the merged channel?".
**Pitfalls.** Closing the output from a forwarder -> panic on send to a closed channel.

### 3. Parallel Map with a Semaphore
**Direction of thinking.** Cap concurrency with a buffered channel of capacity `workers`: acquire a slot before launching, release when done. Results keep input order because each goroutine writes to `out[i]`.
**Pitfalls.** Acquiring inside the goroutine instead of before launching would let unbounded goroutines spawn.

### 4. Pipeline
**Direction of thinking.** Chain stages `Generate -> Square -> collect`, each a goroutine emitting on a channel it **owns and closes** when its input drains. Backpressure is automatic (unbuffered channels block until the next stage reads).
**Pitfalls.** A stage not closing its output channel stalls the downstream `range`.

### 5. Thread-Safe Counter (mutex vs atomic)
**Direction of thinking.** `count++` is a read-modify-write - a **data race** under concurrent access. Guard with a `sync.Mutex`, or for a single integer use `sync/atomic` (lock-free, cheaper). Demonstrate you know both and their trade-offs.
**Verify.** `go test -race` catches the race if you forget the lock.

### 6. First Result / cancellation with context
**Direction of thinking.** Launch all tasks with a shared cancellable `context`; the first to send on a **buffered** result channel wins, and `cancel()` stops the rest. The buffer ensures losing goroutines can still send and exit (no leak).
**Pitfalls.** An unbuffered result channel would leak the slower goroutines (they block on send forever).

---

## Common pitfalls & edge cases

- **Deadlock**: sending on a channel with no receiver, or forgetting to close a channel a `range` waits on.
- **Data race**: concurrent read/write of shared memory without a mutex/atomic - use `go test -race`.
- **Goroutine leak**: a goroutine blocked forever on a channel; give it a way out (buffer, context, close).
- **Closing a channel twice** or sending after close -> panic. The producer owns the close.
- **WaitGroup misuse**: calling `Add` after the goroutine may have started, or forgetting `Done`.

## Interview Q&A

- **Goroutine vs OS thread?** Goroutines are user-space, multiplexed onto threads by the Go runtime; they start with a tiny stack (~2KB) and are far cheaper, so you can have millions.
- **Buffered vs unbuffered channel?** Unbuffered synchronizes sender and receiver (rendezvous); buffered decouples them up to its capacity.
- **When mutex vs channel?** Channels for transferring ownership/dataflow; a mutex/atomic for protecting a small piece of shared state (a counter, a map).
- **How do you cap concurrency?** A worker pool or a buffered-channel semaphore.
- **How do you cancel work?** Pass a `context.Context` and `select` on `ctx.Done()` in each goroutine.
- **What does `go test -race` do?** Instruments memory access to detect data races at runtime - run it on any concurrent code.
