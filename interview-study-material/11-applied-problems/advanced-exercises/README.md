# Applied Interview Exercises

These are **harder, multi-concept problems** — the kind of "onsite hard" question that combines two or three patterns. The skill being tested is **decomposition**: recognizing which known tool solves each sub-part, then wiring them together.

## Programs

### [Airport gate assignment](airport-gate-assignment/airport_gate_assignment.go)

An event-driven scheduler: flights arrive over time; assign each the lowest-numbered free gate, hold it for a duration, then release it. Late arrivals wait in a queue.

**How to think about it — this is a two-heap + queue simulation:**
- **Min-heap of free gate numbers** → so "give me the smallest available gate" is O(log n).
- **Min-heap of occupied gates keyed by `releaseTime`** → so "which gate frees next?" is O(log n). Before assigning, pop everything whose `releaseTime <= now` back into the free heap.
- **Waiting queue (FIFO)** → flights that arrive when all gates are busy wait their turn; if none are free, fast-forward `currentTime` to the next release.

**Why this data-structure combo:** you constantly need two different "minimums" — smallest gate id *and* earliest release time — and a heap gives each in O(log n). A plain scan would be O(n) per event. Go's `container/heap` requires you to implement the `heap.Interface` (`Len/Less/Swap/Push/Pop`) — study that boilerplate; it recurs in every heap problem.

- **Complexity:** O(E log G) for E events and G gates.
- **Recognition cue:** *"assign the smallest/earliest available resource"* + *"scheduling over time"* → heap(s) + event loop.

### [Prime string splits](prime-string-splits/prime_string_splits.go)

Count the ways to split a digit string into contiguous pieces where **every piece is a prime number** (no leading zeros, primes up to 10^6).

**How to think about it — sieve + 1-D DP:**
- **Sieve of Eratosthenes** precomputes primality up to 10^6 once, so each "is this piece prime?" check is O(1). (The sieve itself is O(N log log N).)
- **DP:** let `dp[i]` = number of valid ways to split the first `i` characters. `dp[0] = 1` (empty prefix, one way). For each end `i`, look back at each start `j` and, if `s[j..i)` is a valid prime, add `dp[j]` to `dp[i]`. Answer is `dp[n]`.

**The two clever bounds:** (1) since primes ≤ 10^6 have at most 7 digits, the inner loop only looks back **≤ 7** characters — turning a potential O(n²) into O(7n) = O(n). (2) Skip pieces with a **leading zero**. Results are taken mod 1e9+7 because the count explodes.

**Dry run** — `"11375"` → answer `3` (e.g. `11|3|7|5`, `113|7|5`, `11|37|5`, etc. — the valid prime partitions).

- **Recognition cue:** *"count the number of ways to partition/split"* → DP where `dp[i]` aggregates over valid last-pieces. The prime check is the problem-specific predicate.

### [Wind-blown leaves](wind-blown-leaves/wind_blown_leaves.go)

A grid holds counts of leaves; a wind string (`U/D/L/R`) blows every leaf one cell per step; leaves pushed off the edge are lost. Count how many remain.

**How to think about it — grid simulation:** for each wind character, build a **fresh grid** and move every cell's count to its neighbor in the wind direction; if the target is off-grid, those leaves are dropped. Repeat for each wind, then sum.

**Why a new grid each step (the subtle bug this avoids):** if you moved leaves in-place, a leaf shifted into a not-yet-processed cell could be moved **again in the same step** — double-counting. Writing to a separate `newGrid` guarantees each leaf moves exactly once per wind.

- **Complexity:** O(W · H · L) for grid W×H and L winds.
- **Recognition cue:** *"simulate movement on a grid step by step"* → grid + direction deltas; watch for in-place mutation hazards and out-of-bounds.

## Review checklist

- Gate assignment: why two heaps and a queue, and what "minimum" does each heap serve?
- Can you implement Go's `heap.Interface` (the 5 methods) from memory?
- Prime splits: why does the inner DP loop only look back 7 characters, and why `dp[0]=1`?
- Wind leaves: why allocate a fresh grid each step instead of mutating in place?
- For each, name the underlying pattern(s) you'd shout out in an interview.
