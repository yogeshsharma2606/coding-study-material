# 11 - Heap / Priority Queue

> Code: [`heap.go`](heap.go) - Tests: [`heap_test.go`](heap_test.go) - run `go test ./categories/11-heap-priority-queue/`

## Overview & mental model

A binary heap is a complete tree stored in an array where each parent is <= (min-heap) or >= (max-heap) its children. It gives **O(log n)** push and pop-extreme, and **O(1)** peek at the min/max. Reach for a heap whenever you repeatedly need the current extreme: **top-k**, **k-way merge**, **running median**, **scheduling by priority**.

The defining trick for **top-k**: to keep the k *largest* elements, use a **min-heap of size k** - when it exceeds k, pop the smallest. The root is then the kth largest. (Symmetrically, k smallest -> max-heap of size k.)

## Go implementation note

Go's `container/heap` isn't a data structure - it's a set of algorithms over any type implementing `heap.Interface` (`Len/Less/Swap` + `Push/Pop`). This package wraps it in a generic `PQ[T]` taking a `less` function, so each problem just supplies its ordering.

```go
h := NewPQ(func(a, b int) bool { return a < b }) // min-heap
h.push(x); top := h.pop()
```

## How to recognize it

- "k largest / smallest / closest / most frequent", "kth ...".
- "Merge k sorted ...", "median of a stream", "schedule by priority/deadline".
- Any repeated "give me the current best/worst and remove it".

## How to think / attack plan

1. Do I need the extreme **repeatedly** (not just once)? Sorting once is fine for a single query; a heap wins when the set changes over time or you only need k.
2. Top-k -> **size-k heap of the opposite polarity**. This is O(n log k), better than O(n log n) full sort.
3. Two-heaps pattern for medians / balancing halves.
4. Could a **bucket sort** (bounded keys) or **quickselect** (O(n) average, one-shot kth) beat the heap? Mention the trade-off.

---

## Problems (easy -> hard)

### 1. Kth Largest Element in an Array (medium)
**Direction of thinking.** A full sort is O(n log n). A **min-heap of size k** keeps only the k largest, giving O(n log k). (Quickselect averages O(n) but is O(n^2) worst case - good to mention.)
**Dry run.** `[3,2,1,5,6,4], k=2`: heap keeps {5,6}; root 5 is the 2nd largest.
**Complexity.** Time O(n log k), space O(k).

### 2. Kth Largest in a Stream (easy, design)
**Thinking.** Same size-k min-heap, persisted; `Add` pushes and trims, returning the root.
**Complexity.** O(log k) per add.

### 3. Last Stone Weight (easy)
**Thinking.** Repeatedly smash the two heaviest -> a **max-heap**; push back the difference.
**Complexity.** Time O(n log n).

### 4. K Closest Points to Origin (medium)
**Direction of thinking.** "k closest" -> keep a **max-heap of size k** by squared distance (avoid sqrt); pop the farthest when it overflows.
**Complexity.** Time O(n log k), space O(k).

### 5. Merge k Sorted Lists (hard)
**Direction of thinking.** The next output is always the smallest among the k current heads. A **min-heap of heads** gives that in O(log k); pop it, append, push its successor. Total O(N log k) for N nodes - better than merging pairwise naively.
**Dry run.** heads {1,1,2} -> emit 1, push 4; and so on.
**Complexity.** Time O(N log k), space O(k).

### 6. Find Median from Data Stream (hard) - two heaps
**Direction of thinking.** Keep the lower half in a **max-heap** and the upper half in a **min-heap**, balanced so their sizes differ by at most 1. The median is a root (or the average of both roots). Insert by pushing to one heap, moving its extreme to the other, then rebalancing.
**Dry run.** Add 1,2 -> lo{1}, hi{2} -> median 1.5; add 3 -> lo{2,1}, hi{3} -> median 2.
**Complexity.** O(log n) per add, O(1) per query.

### 7. Top K Frequent Elements (medium)
**Direction of thinking.** Count, then keep a size-k min-heap by frequency. (Bucket sort by frequency is O(n) - see [06 - Hashing](../06-hashing/); the heap is the general-purpose answer.)
**Complexity.** Time O(n log k), space O(n).

### 8. Reorganize String (medium)
**Direction of thinking.** To avoid adjacent duplicates, always place the **most frequent remaining** character that differs from the one just placed. A **max-heap by count** does this; hold the just-used char aside for one step so it can't be picked consecutively. Impossible iff some char count exceeds `(n+1)/2`.
**Complexity.** Time O(n log 26) = O(n), space O(1).

---

## Common pitfalls & edge cases

- Using a max-heap for "k largest" (that's O(n log n)); the size-k **min**-heap is the efficient choice.
- Forgetting to rebalance the two median heaps -> wrong median.
- Comparing distances with sqrt (unnecessary and imprecise) - compare squared distances.
- `container/heap` requires calling `heap.Push/heap.Pop` (which call your `Push/Pop` and re-heapify) - not appending directly.

## Interview Q&A

- **Heap vs sorting for top-k?** Heap is O(n log k) and streaming-friendly; a full sort is O(n log n) and needs all data up front.
- **Heap vs quickselect for kth element?** Quickselect averages O(n) for a one-shot query but is O(n^2) worst case and mutates the array; a heap is O(n log k) and stable in behavior.
- **Why two heaps for a median?** Splitting at the median lets each side's extreme sit right at the boundary, so the median is always a root.
- **Build-heap complexity?** `heap.Init` is O(n) (not O(n log n)) via bottom-up sift-down.
