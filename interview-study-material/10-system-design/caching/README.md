# Caching

Caches trade memory for speed by keeping hot data close. The two design axes interviewers test:

- **Eviction policy** — when full, *what* do you drop? (LRU = drop the least-recently-used.)
- **Expiration (TTL)** — *when* does data go stale, and *how* do you reclaim it (lazy on read vs a background sweeper)?

The recurring senior theme here is **"how do you get O(1)?"** and **"is it safe under concurrency?"**

## Programs

### [LRU cache](lru-cache/lru_cache.go)

O(1) `Get` and `Put` with least-recently-used eviction.

**The core insight (the canonical design answer):** you need two things at once — **O(1) lookup by key** and **O(1) reordering by recency**. No single structure gives both, so **combine two**:
- a **hash map** `key -> node` for O(1) find, plus
- a **doubly linked list** ordered most-recent-at-front for O(1) move/evict.

On `Get`/`Put`, move the touched node to the **front**; when full, evict the node at the **back** (least recently used) and delete its key from the map. Go's `container/list` gives O(1) `MoveToFront`, `PushFront`, `Remove`.

**Dry run** — capacity 2: `Put(1,10)`, `Put(2,20)` → list `[2,1]`. `Get(1)` → moves 1 to front `[1,2]`, returns 10. `Put(3,30)` → full, evict back (2) → `[3,1]`. `Get(2)` → `-1` (evicted). ✓

- **O(1) time** per op, **O(capacity) space.**
- **Why doubly (not singly) linked?** O(1) removal needs the node's predecessor, which only a `prev` pointer gives.
- **Follow-up:** LFU (least *frequently* used) needs frequency buckets; and this version isn't goroutine-safe (add a mutex for concurrency).

### [TTL cache with lazy expiry](ttl-cache-lazy-expiry/ttl_cache_lazy_expiry.go)

Entries carry an `expireAt`; expiry is checked **only when the key is read**.

**How to think about it:** store `value + expireAt` per key. On `Get`, if `now > expireAt`, delete and report a miss. This is **lazy deletion** — cheap and simple, expired data is purged on access.

- **Trade-off (state this):** dead entries that are **never read again** linger forever, wasting memory. Fine for hot keys, bad for large key spaces with churn — which motivates the next program.
- **O(1)** per op; not concurrency-safe as written.

### [Concurrent TTL cache with background cleanup](ttl-cache-background-cleanup/ttl_cache_background_cleanup.go)

Adds a **mutex** for thread safety and a **background goroutine** that periodically sweeps expired keys.

**How to think about it:** two upgrades over the lazy version. (1) Every map access is wrapped in `mu.Lock()/Unlock()` because Go maps are **not safe for concurrent read+write** (a data race will crash). (2) A goroutine driven by a `time.Ticker` runs `cleanup()` on an interval, reclaiming memory from keys that are never read again. A `stop` channel + `select` lets you shut the goroutine down cleanly (avoiding a **goroutine leak**).

**The pattern to remember:**

```
for {
    select {
    case <-ticker.C: cleanup()   // do periodic work
    case <-stop:     return      // graceful shutdown
    }
}
```

- **Trade-offs:** lazy = cheap but leaks memory; background sweep = reclaims memory but costs a periodic O(n) scan and a running goroutine. Real systems often do **both**.
- **Concurrency note:** always provide a `Stop()` so the ticker goroutine doesn't outlive the cache.

## Review checklist

- Why does O(1) LRU need *both* a hash map and a doubly linked list — what does each provide?
- Why doubly linked rather than singly?
- Lazy vs background TTL cleanup: state the memory/CPU trade-off.
- Why must the concurrent cache lock every map access, and why the `stop` channel?
