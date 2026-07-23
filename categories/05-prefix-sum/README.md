# 05 - Prefix Sum

> Code: [`prefix_sum.go`](prefix_sum.go) - Tests: [`prefix_sum_test.go`](prefix_sum_test.go) - run `go test ./categories/05-prefix-sum/`

## Overview & mental model

Define `P[i] = a[0] + a[1] + ... + a[i-1]` (with `P[0] = 0`). Then the sum of any range `[l, r]` is `P[r+1] - P[l]` in **O(1)** after **O(n)** preprocessing. That alone answers range-sum queries.

The powerful idea is the **prefix + hash map** combo: many "count/length of subarrays with property X" problems become "count pairs of prefixes with a relationship." Because a subarray sum is a *difference of two prefixes*, you scan once and use a hash map keyed on prefix values (or their remainders) to count matches in O(n).

Crucially, prefix sums work with **negative numbers**, where sliding window fails.

## How to recognize it

- "Sum of subarray/range", repeated range queries -> prefix array.
- "How many subarrays sum to k", "subarray sum divisible by k", "equal number of 0s and 1s", "longest subarray with sum k" -> prefix + hash map.
- 2D version: "sum of a submatrix" -> 2D prefix sums.

## How to think / attack plan

1. Rewrite the subarray property as a relation between two prefixes. Sum `[l..r] = k` -> `P[r+1] - P[l] = k` -> `P[l] = P[r+1] - k`.
2. Decide what to store in the map: a **count** (for "how many") or the **first index** (for "longest").
3. Seed the map with the empty prefix: `{0: 1}` for counts, `{0: -1}` for lengths.
4. For divisibility, key on `sum mod k` (normalize negatives to `[0,k)`).

---

## Problems (easy -> hard)

### 1. Range Sum Query - Immutable (easy)
**Thinking.** Precompute `prefix`; answer each query as `prefix[r+1]-prefix[l]`.
**Complexity.** Build O(n), each query O(1), space O(n).

### 2. Subarray Sum Equals K (medium) - the canonical one
**Direction of thinking.** Brute force sums every subarray: O(n^2). Instead, as you accumulate `running`, the number of subarrays ending here with sum k equals the number of earlier prefixes equal to `running-k`. Keep a **count** map of prefix values.
**Why seed `{0:1}`.** A subarray starting at index 0 corresponds to the empty prefix having value 0.
**Dry run.** `[1,1,1], k=2`: running 1 (need -1: 0), 2 (need 0: 1 hit), 3 (need 1: 1 hit) -> 2.
**Complexity.** Time O(n), space O(n).

### 3. Contiguous Array - equal 0s and 1s (medium)
**Direction of thinking.** Map `0 -> -1`. Then "equal 0s and 1s" means the running sum **returns to a previous value**. Store the **first index** of each running value; the length is `i - firstIndex`.
**Dry run.** `[0,1,0]` -> sums -1,0,-1; -1 first at 0, again at 2 -> length 2.
**Complexity.** Time O(n), space O(n).

### 4. Find Pivot Index (easy)
**Thinking.** Precompute total. Scan keeping `left`; pivot when `left == total - left - nums[i]`. No map needed.
**Complexity.** Time O(n), space O(1).

### 5. Subarray Sums Divisible by K (medium)
**Direction of thinking.** Two prefixes with the **same remainder mod k** bound a subarray whose sum is divisible by k. Count remainders. Normalize negative remainders with `((x%k)+k)%k` (Go's `%` keeps the sign of the dividend).
**Dry run.** `[4,5,0,-2,-3,1], k=5` -> 7.
**Complexity.** Time O(n), space O(k).

### 6. Continuous Subarray Sum (medium)
**Problem.** Is there a subarray of length >= 2 with sum a multiple of k?
**Direction of thinking.** Same remainder at two indices >= 2 apart -> divisible subarray of length >= 2. Store the **earliest** index per remainder and check the gap.
**Complexity.** Time O(n), space O(k).

### 7. Maximum Size Subarray Sum Equals k (medium)
**Direction of thinking.** Like #2 but we want the **longest**, so store the **first index** of each prefix. If `running-k` was seen, candidate length is `i - firstIndex[running-k]`. Handles negatives (sliding window can't).
**Dry run.** `[1,-1,5,-2,3], k=3` -> length 4 (`[1,-1,5,-2]`).
**Complexity.** Time O(n), space O(n).

### 8. Range Sum Query 2D - Immutable (medium)
**Direction of thinking.** Build a 2D prefix table `P[i][j]` = sum of the rectangle from origin to `(i-1,j-1)`. A submatrix sum is **inclusion-exclusion**: `P[r2+1][c2+1] - P[r1][c2+1] - P[r2+1][c1] + P[r1][c1]`.
**Why add back the corner.** The two subtractions remove the top-left rectangle twice; add it once.
**Complexity.** Build O(mn), each query O(1), space O(mn).

---

## Common pitfalls & edge cases

- Forgetting to seed the map (`{0:1}` for counts, `{0:-1}` for indices) - misses subarrays starting at index 0.
- Go's `%` can be negative - normalize for divisibility problems.
- Using a **count** map when you need the **first index** (longest), or vice versa.
- 2D: off-by-one between the padded prefix grid (size m+1 x n+1) and the original.

## Interview Q&A

- **Prefix sum vs sliding window?** Window needs monotonic validity (usually non-negative values); prefix + hashmap handles negatives and exact-sum counting.
- **Why does "same prefix value" imply a zero-sum (or equal 0/1) subarray?** Equal prefixes means their difference (the subarray between them) sums to zero.
- **How to support updates (mutable range sums)?** Use a Fenwick/Binary Indexed Tree or segment tree for O(log n) updates and queries.
