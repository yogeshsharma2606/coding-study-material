# 06 - Hashing

> Code: [`hashing.go`](hashing.go) - Tests: [`hashing_test.go`](hashing_test.go) - run `go test ./categories/06-hashing/`

## Overview & mental model

A hash map/set gives **amortized O(1)** insert, lookup, and delete. The core interview move: whenever a brute-force solution repeatedly **searches** the data ("for each x, is there a y such that..."), replace the inner search with a hash lookup, collapsing O(n^2) into O(n) at the cost of O(n) space.

Three shapes:
- **Set** (`map[T]struct{}`) - membership / dedup.
- **Counter** (`map[T]int`) - frequencies.
- **Signature map** - key on a derived value (pair-sum, sorted letters, remainder) to group or relate items.

## How to recognize it

- "Contains duplicate", "first unique", "intersection", "anagram groups", "top k frequent".
- Any "find two/four things that combine to X" where sorting isn't required.
- You need O(1) membership tests or frequency counts.

## How to think / attack plan

1. What am I looking up repeatedly? Make that the **key**.
2. Do I need existence (set), a count (counter), or a grouping (signature)?
3. Can I split the problem so the map stores one half and I query with the other (meet-in-the-middle, e.g. 4Sum II)?
4. Remember hashing costs O(n) space and gives amortized (not worst-case) O(1).

## Go note

Use `map[T]struct{}` for sets (`struct{}` is zero-width). Iterating a map is **randomly ordered** - never rely on order; sort the output if determinism matters.

---

## Problems (easy -> hard)

### 1. Contains Duplicate (easy)
**Thinking.** Insert into a set; a repeat means duplicate. (Sorting also works in O(n log n)/O(1) space; hashing is O(n)/O(n).)
**Complexity.** Time O(n), space O(n).

### 2. First Unique Character (easy)
**Thinking.** Two passes: count all, then find the first with count 1. A `[26]int` suffices for lowercase.
**Complexity.** Time O(n), space O(1).

### 3. Intersection of Two Arrays (easy)
**Thinking.** Put one array in a set, scan the other, emit-and-delete to dedupe.
**Complexity.** Time O(n+m), space O(n).

### 4. Top K Frequent Elements (medium)
**Direction of thinking.** Count frequencies, then you need the k largest by count. A heap gives O(n log k). But frequencies are bounded by n, so **bucket sort by frequency** (index = frequency) gives O(n). Walk buckets from high to low.
**Why buckets beat a heap here.** The key (frequency) is a small integer range, so counting-sort-style bucketing is linear.
**Dry run.** `[1,1,1,2,2,3], k=2`: freq {1:3,2:2,3:1}; buckets[3]=[1], buckets[2]=[2] -> [1,2].
**Complexity.** Time O(n), space O(n).

### 5. Longest Consecutive Sequence (medium)
**Direction of thinking.** Sorting gives O(n log n). To hit **O(n)**: put all in a set, and only start counting a run from a value `x` whose predecessor `x-1` is **absent** (a true run start). Each element is visited at most twice total.
**Why the "no predecessor" check matters.** Without it you'd re-walk the same run from every member -> O(n^2).
**Dry run.** `[100,4,200,1,3,2]`: run starts at 1 -> 1,2,3,4 -> length 4.
**Complexity.** Time O(n), space O(n).

### 6. Valid Sudoku (medium)
**Direction of thinking.** A placement is valid iff no digit repeats in its row, column, or 3x3 box. Maintain 9 sets each; the box index is `(r/3)*3 + c/3`. One pass over 81 cells.
**Complexity.** Time O(1) (fixed 9x9), space O(1).

### 7. 4Sum II - count tuples (medium)
**Direction of thinking.** Four arrays, count tuples summing to 0. O(n^4) is hopeless. **Split in half**: hash all `a[i]+b[j]` sums with counts (O(n^2)), then for each `c[k]+d[l]` add the count of its negation. Meet-in-the-middle.
**Complexity.** Time O(n^2), space O(n^2).

### 8. Majority Element (medium)
**Direction of thinking.** A hash counter is O(n) space. But since the majority appears > n/2 times, **Boyer-Moore voting** gets O(1) space: keep a candidate and a counter; matching votes increment, others decrement; on 0 pick a new candidate. The majority survives all cancellations.
**Dry run.** `[2,2,1,1,1,2,2]` -> candidate ends as 2.
**Complexity.** Time O(n), space O(1).

---

## Common pitfalls & edge cases

- Relying on Go map iteration order (it's randomized) - sort outputs for deterministic results/tests.
- Using a map where a fixed array suffices (small known alphabet) - the array is faster and O(1) space.
- Forgetting hashing is amortized O(1); adversarial inputs can degrade it (rarely relevant in interviews, worth mentioning).
- Not deduping outputs when the problem asks for unique results.

## Interview Q&A

- **Set vs map in Go?** Use `map[T]struct{}` for a set - no value overhead.
- **When does hashing beat sorting?** When you need O(n) or original order/indices; sorting wins when you also need order or want O(1) space.
- **Worst-case hash complexity?** O(n) per op under collisions; average/amortized O(1). Go randomizes hash seeds to resist collision attacks.
- **How to get O(n) for top-k?** Bucket by frequency instead of a comparison sort/heap.
