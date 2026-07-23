# 13 - Sorting

> Code: [`sorting.go`](sorting.go) - Tests: [`sorting_test.go`](sorting_test.go) - run `go test ./categories/13-sorting/`

## Overview & mental model

Comparison-based sorting is bounded below by **O(n log n)**. You should know the trade-offs cold:

| Algorithm | Time (avg/worst) | Space | Stable? | Notes |
|---|---|---|---|---|
| Merge sort | O(n log n) / O(n log n) | O(n) | yes | predictable; good for linked lists / external sort |
| Quicksort | O(n log n) / O(n^2) | O(log n) | no | fast in practice; pivot choice matters |
| Heapsort | O(n log n) / O(n log n) | O(1) | no | in-place, no recursion |
| Counting/Radix | O(n + k) | O(n + k) | yes | only for small-range integer keys |
| Insertion | O(n^2) / O(n^2) | O(1) | yes | great for nearly-sorted / tiny arrays |

Two senior-level insights: (1) many problems become trivial once you **sort with the right comparator**; (2) if you only need the **kth element**, **quickselect** averages O(n) - don't sort the whole array.

## How to recognize it

- "Sort by a custom rule" (largest number, by frequency) -> comparator.
- "Kth smallest/largest" one-shot -> quickselect (or a heap; see [11](../11-heap-priority-queue/)).
- Small integer keys, need linear time -> counting/radix.
- The problem gets easy if the data were ordered (intervals, greedy) -> sort first.

## How to think / attack plan

1. Do I need a full sort, or just the kth / top-k? Full -> `sort.Slice`; kth -> quickselect; top-k -> heap.
2. Is stability required (equal keys keep input order)? Merge sort / `sort.Stable`.
3. Are keys small integers? Counting/radix beats O(n log n).
4. For "arrange to optimize concatenation/adjacency", the trick is usually a clever pairwise comparator.

## Go note

Use `sort.Slice(s, less)` for a custom order (not stable) or `sort.SliceStable`. `sort.Ints`, `sort.Strings` for the common cases. `sort.Search` does binary search.

---

## Problems (easy -> hard)

### 1. Merge Sort (implement) (medium)
**Direction of thinking.** Divide the array in half, sort each recursively, then **merge** two sorted halves in linear time. Stable and worst-case O(n log n), at the cost of O(n) scratch space.
**Dry run.** `[5,2,9,1]` -> `[5,2]`,`[9,1]` -> `[2,5]`,`[1,9]` -> merge -> `[1,2,5,9]`.
**Complexity.** Time O(n log n), space O(n).

### 2. Quicksort (implement) (medium)
**Direction of thinking.** Pick a pivot, **partition** so smaller elements go left and larger right, recurse on both sides. In place (O(log n) stack). Worst case O(n^2) on bad pivots - mitigate with a middle/median or random pivot.
**Complexity.** Time O(n log n) avg / O(n^2) worst, space O(log n).

### 3. Quickselect - kth smallest (medium)
**Direction of thinking.** Like quicksort, but after partitioning you only recurse into the side containing index `k-1`. Average O(n) (n + n/2 + n/4 + ... = 2n).
**Dry run.** `[3,2,1,5,6,4], k=2`: partition until the pivot lands at index 1.
**Complexity.** Time O(n) avg / O(n^2) worst, space O(1).

### 4. Counting Sort (easy)
**Direction of thinking.** When values are small non-negative integers, tally counts and emit in order - no comparisons, O(n + maxVal). This underlies radix sort.
**Complexity.** Time O(n + k), space O(k).

### 5. Largest Number (medium) - custom comparator
**Direction of thinking.** To maximize the concatenation, sort strings so that `a` precedes `b` iff `a+b > b+a` (e.g. "9" before "34" because "934" > "349"). Handle the all-zeros edge case.
**Dry run.** `[3,30,34,5,9]` -> "9534330".
**Complexity.** Time O(n log n * L), space O(n).

### 6. H-Index (medium)
**Direction of thinking.** Sort citations descending; the h-index is the largest `i` such that the i-th paper has `>= i` citations. Scanning the sorted list finds it directly.
**Complexity.** Time O(n log n), space O(1).

### 7. Sort Characters By Frequency (medium)
**Direction of thinking.** Count frequencies, then sort distinct characters by descending count and expand. (A bucket-by-frequency approach is O(n); see [06](../06-hashing/).)
**Complexity.** Time O(n log n), space O(n).

### 8. Wiggle Sort (medium)
**Direction of thinking.** Want `a[0] <= a[1] >= a[2] <= ...`. A single greedy pass fixes each adjacent pair by swapping when the local relation is violated; because each swap only strengthens the previous relation, one pass suffices - no full sort.
**Complexity.** Time O(n), space O(1).

---

## Common pitfalls & edge cases

- Quicksort's O(n^2) on already-sorted input with a naive first-element pivot - use a better pivot.
- Forgetting stability requirements (use merge/stable sort).
- Counting sort with negative or huge-range values (needs offset or a different algorithm).
- Largest Number: the all-zeros case must return "0", not "000".
- `sort.Slice` is not stable; use `sort.SliceStable` when equal keys must preserve order.

## Interview Q&A

- **Merge vs quicksort?** Merge is stable and worst-case O(n log n) but O(n) space; quicksort is in-place and faster in practice but O(n^2) worst case and unstable.
- **Why is comparison sort bounded by O(n log n)?** There are n! orderings; a comparison tree needs log2(n!) = O(n log n) comparisons to distinguish them.
- **When can you beat O(n log n)?** With bounded integer keys via counting/radix sort - O(n).
- **Kth element without full sort?** Quickselect, O(n) average.
- **Is Go's `sort.Slice` stable?** No; use `sort.SliceStable` (it's an insertion/merge hybrid) when you need stability.
