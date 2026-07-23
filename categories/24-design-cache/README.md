# 24 - Design & Cache

> Code: [`cache.go`](cache.go) - Tests: [`cache_test.go`](cache_test.go) - run `go test ./categories/24-design-cache/`

## Overview & mental model

"Design" problems ask you to build a data structure meeting specific complexity guarantees (usually **O(1)** for every operation). The universal trick: a hash map gives O(1) lookup but no ordering, so you **combine it with a second structure** that maintains the ordering you need:

| Requirement | Combine map with |
|---|---|
| Recency (LRU) | Doubly linked list (front = newest) |
| Frequency (LFU) | Per-frequency lists + a `minFreq` pointer |
| O(1) random element | A slice (swap-with-last to delete) |
| Value-at-time | Sorted list of (timestamp, value) + binary search |
| Rate limiting | Token bucket / sliding-window counters |

State the **invariant** that keeps the two structures consistent - that's what interviewers probe.

## How to recognize it

- "Design a cache / LRU / LFU", "O(1) get and put".
- "Insert, delete, and getRandom in O(1)".
- "Get the value at or before time t", "rate limiter", "implement a hashmap".

## How to think / attack plan

1. List each operation and its **required complexity** - that dictates the structures.
2. What ordering/auxiliary info does eviction or retrieval need? Add the structure that provides it in O(1)/O(log n).
3. Define the **sync invariant**: after every op, the map and the ordering structure agree.
4. Handle capacity 0, updates to existing keys, and eviction tie-breaks explicitly.

---

## Problems

### 1. LRU Cache (medium/hard) - the classic
**Direction of thinking.** Need O(1) get/put with eviction of the **least recently used**. A doubly linked list orders nodes by recency (move to front on access; evict from the back); a hash map maps keys to list nodes for O(1) splicing.
**Invariant.** `items[key]` points to the exact list node holding that key.
**Dry run.** cap 2: put(1),put(2),get(1) [order 1,2 -> 1 front], put(3) evicts 2.
**Complexity.** O(1) per op, space O(capacity).

### 2. LFU Cache (hard)
**Direction of thinking.** Evict the **least frequently used**, breaking ties by **least recently used**. Keep each key's frequency, one doubly linked list per frequency (recency within a frequency), and a `minFreq` pointer. On access, move the node from its freq list to freq+1; on insert-when-full, evict the back of the `minFreq` list.
**Invariant.** `minFreq` is the smallest non-empty frequency; each freq list is ordered by recency.
**Complexity.** O(1) per op, space O(capacity).

### 3. Insert Delete GetRandom O(1) (medium)
**Direction of thinking.** `getRandom` needs an array (O(1) index); `delete` in O(1) from an array is impossible in general - unless you **swap the target with the last element** and shrink. A map tracks each value's index.
**Invariant.** `idx[val]` is the position of `val` in the slice.
**Complexity.** O(1) per op, space O(n).

### 4. Time-Based Key-Value Store (medium)
**Direction of thinking.** Timestamps for a key arrive **increasing**, so store `(timestamp, value)` appended per key and **binary search** for the largest timestamp <= query.
**Complexity.** Set O(1), Get O(log n).

### 5. Token-Bucket Rate Limiter (design)
**Direction of thinking.** Tokens refill continuously at a fixed rate up to a capacity; each request consumes one. On each request, add `elapsed * rate` tokens (capped), then allow iff >= 1 token remains. Injecting the clock (`nowMs`) makes it deterministic/testable. Alternatives: fixed-window counter (simple, bursty at edges) and sliding-window log (accurate, more memory).
**Complexity.** O(1) per request.

### 6. Design HashMap (medium)
**Direction of thinking.** Buckets + chaining: `hash(key)=key % numBuckets` (a prime reduces clustering); each bucket is a list of key-value pairs scanned linearly. Demonstrates how maps work under the hood (load factor, collisions, resize).
**Complexity.** O(1) average, O(n) worst per op.

---

## Common pitfalls & edge cases

- LRU/LFU: forgetting to update the map when moving/removing nodes -> dangling pointers.
- LFU: not advancing `minFreq` correctly when a frequency list empties.
- RandomizedSet: not updating the moved element's index after swap-with-last.
- Capacity 0 caches; updating an existing key must not trigger eviction.
- Rate limiter: not capping tokens at capacity (allows unbounded bursts).

## Interview Q&A

- **Why a doubly (not singly) linked list for LRU?** O(1) removal of an arbitrary node requires access to its predecessor - a doubly linked list gives that directly.
- **How does LFU break ties?** By recency within the same frequency (LRU inside each frequency bucket).
- **How to delete from an array in O(1)?** Swap the element with the last and shrink; order isn't preserved, which is fine for a set.
- **Fixed vs sliding window vs token bucket rate limiting?** Fixed window is simplest but bursts at boundaries; sliding-window log is precise but memory-heavy; token bucket smooths bursts with O(1) state.
- **What happens when a hashmap's load factor gets high?** Collisions grow; real maps **resize/rehash** to keep operations near O(1).
