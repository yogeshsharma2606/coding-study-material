# Go Concurrency

Concurrency is the senior-interview differentiator. Go's model is small but you must wield it correctly:

- **Goroutine** — a cheap independent unit of execution (`go f()`).
- **Channel** — a typed pipe for *communicating* between goroutines. Go's mantra: **"share memory by communicating"** (pass data over channels) rather than "communicate by sharing memory" (locks around shared state).
- **`sync.WaitGroup`** — wait for a set of goroutines to finish (`Add`/`Done`/`Wait`).
- **`sync.Mutex`** — guard shared state when channels don't fit.

**The four bugs interviewers probe for:** *deadlock* (everyone waiting), *race condition* (unsynchronized shared writes — run with `go test -race`), *goroutine leak* (a goroutine blocked forever on a channel), and *channel misuse* (sending on a closed channel panics; only the sender should `close`).

## Programs

### [Odd/even coordination](odd-even-coordination/odd_even_coordination.go)

Two goroutines print numbers 1..n strictly alternating odd, even, odd, even… in order.

**How to think about it:** you want interleaving, so use channels as **turn tokens**. Each goroutine blocks on receiving its token, does its work, then hands the token to the other. This is a **ping-pong / signaling** pattern — the channel isn't carrying data, it's carrying *permission to proceed*. `main` kicks it off by sending the first token to odd; the last even number closes both channels to end cleanly.

- **Teaching point:** unbuffered channels are a **rendezvous** — a send blocks until a receive is ready, which is exactly the handshake that forces alternation.
- **Gotcha:** ordering of who closes the channel matters; sending on a closed channel panics.

### [Producer-consumer](producer-consumer/producer_consumer.go)

One producer feeds a channel; multiple consumers (workers) drain it — and a new consumer is added **dynamically** mid-stream.

**How to think about it:** this is the **worker-pool / fan-out** pattern. Consumers each `for val := range dataChan` — a `range` over a channel automatically ends when the channel is **closed and drained**, which is how workers know to shut down. The `WaitGroup` lets `main` block until every consumer has exited. The producer is the sole owner of `close(dataChan)` — closing signals "no more work."

- **Why `range` + `close` is elegant:** no sentinel values, no manual "are we done?" flag — closing the channel broadcasts completion to all workers at once.
- **Backpressure:** an unbuffered channel means the producer blocks until a consumer is ready — natural flow control.

### [Parallel slice sum](parallel-slice-sum/parallel_slice_sum.go)

Sum a large slice by splitting it across `runtime.NumCPU()` goroutines, each summing a chunk, then combining.

**How to think about it:** this is **map-reduce / fan-out-fan-in**. Partition the data (map), each goroutine computes a partial sum independently (no shared writes → no locks needed), and push results into a **buffered** channel. `main` fans them back in by reading exactly `n` partials and adding them.

**The two details that matter:** (1) the **last chunk takes the remainder** (`end = len(nums)`) so integer division doesn't drop trailing elements; (2) pass the slice as a **function argument** `func(part []int)` to avoid the classic loop-variable capture bug.

- **Teaching point:** parallelism only helps CPU-bound work large enough to outweigh goroutine overhead; for a tiny slice the sequential sum wins.

### [Deadlock-free money transfer](deadlock-free-transfer/deadlock_free_transfer.go)

Concurrent transfers between accounts, each needing **two** locks, without deadlocking.

**The classic deadlock and its fix (this is the whole lesson):** if goroutine A locks account 1 then waits for 2, while goroutine B locks 2 then waits for 1, both wait forever — a **circular wait**. The fix is **stable lock ordering**: *always acquire locks in a consistent global order* (here, by ascending account `ID`). With a total order, a cycle is impossible, so deadlock is impossible.

- **Why it works:** deadlock needs all four Coffman conditions; imposing a lock order breaks the "circular wait" condition.
- **Verify with `go run -race`** to confirm no data races on `balance`.
- **Note:** `defer Unlock()` guarantees release even if the body panics.

## Review checklist

- "Share memory by communicating" — when do you reach for channels vs a mutex?
- Why does `for range ch` + `close(ch)` cleanly shut down a worker pool, and who is allowed to close?
- Explain the circular-wait deadlock and how stable lock ordering prevents it.
- In fan-out-fan-in, why is no lock needed on the partial sums, and why pass the chunk as an argument?
- Which of these would `go test -race` flag, and what are the four classic concurrency bugs?
