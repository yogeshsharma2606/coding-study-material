# 12 - Binary Search

> Code: [`binary_search.go`](binary_search.go) - Tests: [`binary_search_test.go`](binary_search_test.go) - run `go test ./categories/12-binary-search/`

## Overview & mental model

Binary search repeatedly halves the search space, giving **O(log n)**. It applies whenever the space is **monotonic**: either the array is sorted, or a boolean predicate `feasible(x)` transitions once from false to true (`F F F T T T`). Then you binary-search for the boundary.

The senior-level generalization is **binary search on the answer**: when you can't search an array directly but you *can* cheaply test "is candidate value `v` good enough?", binary-search over the value range. This solves optimization problems (min speed, min capacity, smallest largest sum) in `O(n log(range))`.

## How to recognize it

- Sorted (or rotated-sorted) array + "find/insert/first/last".
- "Minimize the maximum" / "maximize the minimum" / "smallest X such that feasible".
- Huge value range (up to 1e9+) with a fast feasibility check -> answer search.
- "Find peak", "find min in rotated" - local monotonicity.

## How to think / attack plan

1. Is there monotonicity? If sorted, direct search. If not, is there a predicate that flips once? That's your search key.
2. Pick an invariant and stick to it. Two robust choices:
   - **Half-open** `[lo, hi)` with `hi = len`, loop `lo < hi`, `hi = mid` / `lo = mid+1` -> naturally computes a **lower bound**.
   - **Closed** `[lo, hi]`, loop `lo <= hi`, `hi = mid-1` / `lo = mid+1` -> classic exact match.
3. Always use `mid = lo + (hi-lo)/2` to avoid overflow.
4. For "answer search", define `feasible(v)`, find the range `[lo, hi]`, and search for the boundary.

## Core template - lower bound (first index with a[i] >= target)

```go
lo, hi := 0, len(a)
for lo < hi {
    mid := lo + (hi-lo)/2
    if a[mid] < target { lo = mid + 1 } else { hi = mid }
}
return lo
```

---

## Problems (easy -> hard)

### 1. Binary Search (easy)
**Thinking.** The canonical exact-match search. Master one template to avoid off-by-one bugs.
**Complexity.** Time O(log n), space O(1).

### 2. Search Insert Position (easy)
**Thinking.** This is exactly a **lower bound**: first index with `a[i] >= target`.
**Complexity.** O(log n).

### 3. Find First and Last Position (medium)
**Direction of thinking.** Two boundary searches: leftmost index `>= target` and (leftmost index `> target`) - 1. If the left boundary isn't `target`, it's absent.
**Dry run.** `[5,7,7,8,8,10], t=8`: left=3, rightBoundary(9)-1=4 -> [3,4].
**Complexity.** O(log n).

### 4. Search in Rotated Sorted Array (medium)
**Direction of thinking.** After a rotation, at least one half of `[lo, mid, hi]` is still sorted. Detect the sorted half (`a[lo] <= a[mid]`), check whether the target lies within it, and recurse into the correct side.
**Dry run.** `[4,5,6,7,0,1,2], t=0`: left half sorted, 0 not in [4,7] -> go right -> find at 4.
**Complexity.** O(log n).

### 5. Find Minimum in Rotated Sorted Array (medium)
**Direction of thinking.** Compare `a[mid]` to `a[hi]`: if `a[mid] > a[hi]`, the min is strictly right; else it's at `mid` or left. Converges to the pivot.
**Complexity.** O(log n).

### 6. Find Peak Element (medium)
**Direction of thinking.** Treat out-of-bounds as -inf. If `a[mid] < a[mid+1]`, a peak must exist to the right (you're on an ascending slope); go there. Otherwise a peak is at `mid` or left.
**Why a peak always exists in the chosen half.** Walking uphill in a bounded array must reach a peak.
**Complexity.** O(log n).

### 7. Koko Eating Bananas (medium) - answer search
**Direction of thinking.** You can't binary-search the piles, but you can binary-search the **eating speed**. `feasible(speed)` = total hours <= h, which is **monotonic** (faster is never worse). Search speed in `[1, max pile]`.
**Dry run.** `piles=[3,6,7,11], h=8` -> min speed 4.
**Complexity.** Time O(n log(maxPile)), space O(1).

### 8. Capacity to Ship Packages Within D Days (hard) - answer search
**Direction of thinking.** Binary-search the **capacity**. Lower bound = the heaviest package (must fit); upper bound = sum of all (one day). `feasible(cap)` = days needed <= D.
**Complexity.** Time O(n log(sum)), space O(1).

### 9. Sqrt(x) (easy) - answer search
**Thinking.** Find the largest `m` with `m*m <= x`. Binary search `[1, x/2]`; return the floor.
**Complexity.** O(log x).

---

## Common pitfalls & edge cases

- Off-by-one from mixing invariants - commit to `[lo, hi)` or `[lo, hi]`, not both.
- Integer overflow in `mid` (use `lo + (hi-lo)/2`) and in `mid*mid` for large x.
- Infinite loops when the update doesn't shrink the range (e.g. `lo = mid` without `+1`).
- Answer search: getting the `lo`/`hi` bounds wrong or a non-monotonic feasibility function.

## Interview Q&A

- **How do you know binary search applies?** The space is sorted or a predicate is monotonic (`F...F T...T`).
- **What is "binary search on the answer"?** Binary-searching over the value range using a feasibility test, when direct array search doesn't apply.
- **Why `lo + (hi-lo)/2`?** Avoids `lo+hi` overflow in fixed-width integers.
- **Lower vs upper bound?** Lower bound = first `>= target`; upper bound = first `> target`. Their difference is the count of `target`.
