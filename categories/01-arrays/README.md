# 01 - Arrays

> Code: [`arrays.go`](arrays.go) - Tests: [`arrays_test.go`](arrays_test.go) - run `go test ./categories/01-arrays/`

## Overview & mental model

An array is a block of contiguous memory, so **random access by index is O(1)** but **insertion/deletion in the middle is O(n)** (everything shifts). Almost every "optimal" array solution is one of three ideas:

1. **Single scan with an invariant** - carry one or two running values (a min-so-far, a running sum) and update the answer as you go. Turns O(n^2) into O(n).
2. **Two indices** - a slow/fast pair (partitioning in place) or a converging pair on sorted data. Removes the inner loop.
3. **Trade space for time** - a hash map to remember what you've seen, so you don't re-scan.

## How to recognize it

- Input is a flat list of numbers and the answer is a single value, an index, or a rearrangement.
- The words **"in-place"** (target O(1) extra space), **"sorted"** (enables two pointers / binary search), or **"contiguous subarray"** (Kadane / sliding window / prefix sum).

## How to think / attack plan

1. Can I sort? Does sortedness unlock two pointers or binary search? (But sorting costs O(n log n) and destroys original indices.)
2. Is the answer over **subarrays** (contiguous) or **subsequences**? Contiguous -> running sum / window. 
3. What single running quantity, if I tracked it during one pass, would let me compute the answer? That quantity is your invariant.
4. If brute force re-searches the array, replace the search with a **hash map**.

## Core technique - the in-place two-pointer partition

```go
slow := 0
for fast := 0; fast < len(nums); fast++ {
    if keep(nums[fast]) {
        nums[slow], nums[fast] = nums[fast], nums[slow]
        slow++
    }
}
// nums[:slow] now holds the kept elements, in order
```

---

## Problems (easy -> hard)

### 1. Two Sum (easy)

**Problem.** Given `nums` and `target`, return indices of the two numbers summing to `target`. Exactly one solution; can't use the same element twice.

**Direction of thinking.** Brute force checks every pair: O(n^2). The repeated work is "for each x, search the rest for target-x." Replace that inner search with an O(1) hash lookup. Do it in one pass by checking the map *before* inserting, which also guarantees we don't reuse an element.

**Why this approach.** Hashing converts the search from O(n) to O(1), the canonical space-for-time trade.

```go
seen := make(map[int]int)
for i, x := range nums {
    if j, ok := seen[target-x]; ok { return []int{j, i} }
    seen[x] = i
}
```

**Dry run.** `nums=[2,7,11,15], target=9`

| i | x | need=9-x | seen before? | action |
|---|---|---|---|---|
| 0 | 2 | 7 | no | put 2->0 |
| 1 | 7 | 2 | yes (0) | return [0,1] |

**Complexity.** Time O(n), space O(n).

---

### 2. Best Time to Buy and Sell Stock (easy)

**Problem.** One buy then one later sell; maximize profit. Return 0 if none.

**Direction of thinking.** Profit selling today = `today - (cheapest day so far)`. So keep the min price seen and the best profit; a single scan suffices. Brute force pairs would be O(n^2).

**Why.** The best sell day only depends on the min *before* it, which is one running variable.

**Dry run.** `[7,1,5,3,6,4]`: min tracks 7->1; best profit updates at 5 (4), then 6 (5). Answer **5**.

**Complexity.** Time O(n), space O(1).

---

### 3. Maximum Subarray / Kadane (medium)

**Problem.** Find the contiguous subarray with the largest sum.

**Direction of thinking.** Define `cur` = best sum of a subarray *ending exactly at i*. At each i you either extend (`cur + x`) or start a new subarray (`x`). Take the max; track the global best. This is the seed of dynamic programming (`dp[i]` depends on `dp[i-1]`).

**Why.** A negative running prefix can only hurt future sums, so dropping it (restarting) is provably optimal.

**Dry run.** `[-2,1,-3,4,-1,2,1,-5,4]`

| x | -2 | 1 | -3 | 4 | -1 | 2 | 1 | -5 | 4 |
|---|---|---|---|---|---|---|---|---|---|
| cur | -2 | 1 | -2 | 4 | 3 | 5 | 6 | 1 | 5 |
| best | -2 | 1 | 1 | 4 | 4 | 5 | **6** | 6 | 6 |

Answer **6** (subarray `[4,-1,2,1]`).

**Complexity.** Time O(n), space O(1).

---

### 4. Maximum Product Subarray (medium)

**Problem.** Largest product of a contiguous subarray.

**Direction of thinking.** Sums are monotonic under extension, but **products flip sign** with a negative. A very negative running product can become the maximum after multiplying by another negative. So track both `curMax` and `curMin`; when `x < 0`, swap them before updating.

**Why.** The optimal answer ending at i is `max(x, curMax*x, curMin*x)` - you must remember the most-negative candidate too.

**Dry run.** `[-2,3,-4]`: at 3 -> max 3,min -6 (wait, -2 first). Start max=min=-2. x=3: max=3,min=-6. x=-4: swap -> max=-6,min=3; max=max(-4,24)=24. Answer **24**.

**Complexity.** Time O(n), space O(1).

---

### 5. Move Zeroes (easy)

**Problem.** Move all 0s to the end in-place, keep the order of non-zeros.

**Direction of thinking.** This is the two-pointer partition: `slow` marks where the next non-zero goes; `fast` scans. Swap non-zeros forward. The zeros naturally bubble to the back.

**Why.** In-place and stable in one pass; no extra array.

**Dry run.** `[0,1,0,3,12]` -> swaps put 1,3,12 up front -> `[1,3,12,0,0]`.

**Complexity.** Time O(n), space O(1).

---

### 6. Product of Array Except Self (medium)

**Problem.** `out[i]` = product of all elements except `nums[i]`, **without division**, in O(n).

**Direction of thinking.** `out[i] = (product of everything left of i) * (product of everything right of i)`. Compute left-products in a forward sweep, then multiply by right-products in a backward sweep using one running variable.

**Why.** Division is banned (and breaks on zeros). Two directional prefix products avoid it and reuse the output array for O(1) extra space.

**Dry run.** `[1,2,3,4]` -> left: `[1,1,2,6]`; multiply by right (`right`=1,4,12,24): `[24,12,8,6]`.

**Complexity.** Time O(n), space O(1) extra (output aside).

---

### 7. Rotate Array (medium)

**Problem.** Rotate right by `k` in-place.

**Direction of thinking.** Naive: shift one step k times -> O(n*k). Better: rotating right by k means the last k elements come to the front. The **reverse trick**: reverse the whole array, then reverse the first k and the remaining n-k.

**Why.** Three reversals = O(n) time, O(1) space, and no temp array.

**Dry run.** `[1..7], k=3`: reverse all -> `[7,6,5,4,3,2,1]`; reverse first 3 -> `[5,6,7,4,3,2,1]`; reverse rest -> `[5,6,7,1,2,3,4]`.

**Complexity.** Time O(n), space O(1). Remember `k %= n`.

---

### 8. Sort Colors / Dutch National Flag (medium)

**Problem.** Sort an array of only 0,1,2 in one pass, in-place.

**Direction of thinking.** Three regions: `[0,low)`=0s, `[low,mid)`=1s, `(high,end]`=2s, and `[mid,high]` unknown. Scan `mid`: a 0 swaps down to `low`; a 2 swaps up to `high` (and you *don't* advance mid, because the swapped-in value is unexamined); a 1 stays.

**Why.** A counting sort needs two passes; this classic partition does it in one.

**Dry run.** `[2,0,2,1,1,0]` -> `[0,0,1,1,2,2]`.

**Complexity.** Time O(n), space O(1).

---

### 9. Merge Sorted Array In-Place (medium)

**Problem.** `nums1` (size m+n, last n slots empty) and `nums2` (size n) are sorted; merge into `nums1`.

**Direction of thinking.** Merging front-to-back would overwrite unread `nums1` values. **Fill from the back**, largest first, so you always write into already-vacated slots.

**Why.** Back-filling gives O(m+n) time with O(1) space and no temporary buffer.

**Dry run.** `nums1=[1,2,3,_,_,_], nums2=[2,5,6]` -> write 6,5,3,2,2,1 from the back -> `[1,2,2,3,5,6]`.

**Complexity.** Time O(m+n), space O(1).

---

## Common pitfalls & edge cases

- Forgetting `k %= n` in rotation (panics or no-ops when k >= n).
- Empty arrays and single-element arrays - guard early (e.g., `MaxProfit`).
- Kadane initialized to 0 fails on all-negative inputs; initialize to `nums[0]`.
- In Dutch flag, advancing `mid` after a 2-swap skips an unexamined element.
- Integer overflow in product problems for large inputs (use `int64` if constrained).

## Interview Q&A

- **Why is a hash map O(1)?** Amortized average; worst case O(n) with adversarial collisions. Good enough for interviews, mention it.
- **Subarray vs subsequence?** Subarray is contiguous (window/Kadane); subsequence keeps order but allows gaps (often DP).
- **When NOT to sort?** When you need original indices (Two Sum) or must stay O(n).
- **How to make an in-place algorithm stable?** The slow/fast swap preserves relative order of kept elements.
