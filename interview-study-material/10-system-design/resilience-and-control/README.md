# Resilience and Traffic Control

Distributed systems fail partially and get overloaded. Two patterns protect them:

- **Circuit breaker** — stop hammering a **failing dependency**; fail fast and give it time to recover.
- **Rate limiter** — cap how much traffic you **accept**, protecting yourself (and downstreams) from overload.

Both are stateful and shared across goroutines, so **thread safety** is part of the answer.

## Programs

### [Circuit breaker](circuit-breaker/circuit_breaker.go)

A three-state machine wrapping a fallible call, modeled on an electrical breaker.

**The state machine (memorize the transitions):**
- **Closed** (normal): calls pass through. Count failures; when they hit `failureThreshold`, **trip to Open**.
- **Open** (tripped): reject calls **immediately** (fail fast, no waiting on a dead service). After `retryTimeout` elapses, move to **Half-Open**.
- **Half-Open** (probing): allow a **trial** call. Success → **Closed** (reset failures); failure → back to **Open**.

**How to think about it:** the goal is to **stop wasting time and resources** on a service that's clearly down, while still periodically checking if it's healed. Without it, a slow/failing dependency causes threads to pile up (cascading failure). The `time.Since(lastFailureTime) > retryTimeout` check is what triggers the recovery probe.

- All state reads/writes are under a `sync.Mutex` because many goroutines share one breaker.
- **Real-world extras to mention:** rolling failure windows, half-open concurrency limits, per-endpoint breakers.

### [Rate limiter](rate-limiter/rate_limiter.go)

**Token-bucket** limiter built from a buffered channel.

**The insight (why a buffered channel *is* a token bucket):** a buffered channel of capacity `N` is literally a bucket holding up to `N` tokens. `Allow()` does a **non-blocking receive** (`select` with `default`): got a token → allowed; empty → rate-limited (return `429`). A background goroutine ticks on `refillInterval` and does a **non-blocking send** to add a token, with `default` dropping it when the bucket is full (enforcing the cap).

**How to think about it:** tokens = permission. The bucket allows short **bursts** (up to `N` at once) but the refill rate bounds the **sustained** average. The two `select { case ...: default: }` blocks are the whole trick — non-blocking send/receive means neither `Allow()` nor the refiller ever blocks.

**Dry run** — cap 3, refill every 300ms, requests every 200ms: first 3 requests consume the initial tokens (allowed), then requests outpace refills so some get `429`, and a token reappears roughly every 300ms.

- **Token bucket vs alternatives:** *leaky bucket* smooths output to a constant rate (no bursts); *fixed/sliding window counters* count requests per time window. Token bucket is the go-to for "allow bursts but cap the average."

## Review checklist

- Draw the circuit-breaker states and name every transition trigger.
- Why does a circuit breaker prevent *cascading* failure?
- Why is a buffered channel a natural token bucket, and what do the `select`/`default` blocks accomplish?
- Token bucket vs leaky bucket vs sliding window — when would you pick each?
- Where is the shared state, and how is it made goroutine-safe in each?
