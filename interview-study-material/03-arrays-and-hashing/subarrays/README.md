# Subarray Patterns

A **subarray** is a *contiguous* slice of an array. Almost every subarray problem is a fight to get from the brute-force **O(n²)** ("try all start/end pairs") down to **O(n)** by removing repeated work. The three tools that do it:

1. **Fixed-size sliding window** — window length is given (size k). Add the new element, drop the old one.
2. **Prefix sums + hash map** — for "sum equals k" over arbitrary-length ranges. `sum(i..j) = prefix[j] − prefix[i−1]`.
3. **Kadane** — for "maximum-sum subarray"; a tiny DP that decides at each step "extend or restart".

**How to pick:** *fixed length* → fixed window. *"sum = k", "count subarrays with..."* → prefix sum + map. *"maximum/minimum sum"* → Kadane. *variable-length window with a constraint* → the two-pointer sliding window (see category 04).

## Programs

### [Subarray with given sum — brute force](subarray-with-given-sum/subarray_with_given_sum.go)

Find a contiguous subarray summing to a target by trying every start `i` and extending end `j`, accumulating the sum.

**How to think about it:** this is the **baseline you should always be able to state first**. Fixing `i` and growing `j` while keeping a running `sum` avoids recomputing each window from scratch, so it's O(n²), not O(n³).

- **O(n²) time, O(1) space.** Use it to explain *why* you then optimize with prefix sums.

### [Longest subarray with sum K — brute force](longest-sum-k-brute-force/longest_sum_k_brute_force.go)

Same double loop, but tracks the **length** of the longest range whose sum equals `k`.

**How to think about it:** identical structure to the above — the only change is that when `sum == k` you record `j - i + 1` and keep the max. This works for negatives too (unlike a naive window), which is exactly why the optimized version needs prefix sums rather than a sliding window.

- **O(n²) time, O(1) space.**

### [Longest subarray with sum K — optimized (prefix + map)](longest-sum-k-prefix-map/longest_sum_k_prefix_map.go)

O(n) solution using a hash map of prefix sum → **earliest** index.

**The key insight (memorize this):** if the running sum up to `i` is `S`, then a subarray ending at `i` has sum `k` **iff** some earlier prefix equalled `S − k`. So at each index: check whether `S == k` (whole prefix works) or whether `S − k` was seen before; if so, the length is `i − firstIndex(S−k)`. Store each prefix sum's **first** occurrence only, because the earliest start gives the **longest** subarray.

The file contains an exhaustive boxed **walkthrough** on `[10,5,2,7,1,9], k=15` — read it once and prefix-sum-with-map will stick. Answer there is length 4 (`[5,2,7,1]`).

- **O(n) time, O(n) space.** Handles negatives (a sliding window can't).
- **Why store earliest index?** For *longest* subarray you want the farthest-back matching prefix. (For *counting* subarrays you'd store frequencies instead.)

### [Maximum sum of a size-K subarray](maximum-sum-fixed-window/maximum_sum_fixed_window.go)

Largest sum among all windows of exactly length `k`, via a **fixed sliding window**.

**How to think about it:** compute the first window's sum once. To move the window right by one, you don't re-add k elements — you **add the entering element and subtract the leaving one** (`windowSum += arr[i] - arr[i-k]`). That's the whole trick: reuse the previous window's work.

**Dry run** — `[2,1,5,1,3,2], k=3`: first window `2+1+5=8`; slide → `8+1−2=7`; → `7+3−1=9`; → `9+2−5=6`. Max = **9**.

- **O(n) time, O(1) space.** Returns −1 when the array is shorter than `k`.

### [Maximum subarray — Kadane](maximum-subarray-kadane/maximum_subarray_kadane.go)

Maximum-sum contiguous subarray (values may be negative), also returning the actual range.

**The insight that makes Kadane unforgettable:** at each element ask one question — *"is the running sum helping me or dragging me down?"* If the running sum has gone **negative**, it can only hurt what comes next, so **throw it away and start fresh** at the current element. Otherwise **extend**. Track the best sum ever seen (and the start/end to recover the subarray).

**Dry run** — `[-2,1,-3,4,-1,2,1,-5,4]`:

```
cur=-2 best=-2
1:  cur<0 -> restart cur=1   best=1
-3: cur=1+(-3)=-2            best=1
4:  cur<0 -> restart cur=4   best=4
-1: cur=3                    best=4
2:  cur=5                    best=5
1:  cur=6                    best=6   <- max
-5: cur=1                    best=6
4:  cur=5                    best=6
```

Answer **6**, subarray `[4,-1,2,1]`.

- **O(n) time, O(1) space.** It's really 1-D DP: `best_ending_here = max(x, best_ending_here + x)`.
- **Edge case:** all-negative arrays — this version seeds with `arr[0]`, so it correctly returns the least-negative single element.

## Review checklist

- Given a problem, can you instantly pick window vs prefix-sum-map vs Kadane?
- Why does the "sum = k" optimization need prefix sums (and handle negatives) while max-fixed-window uses a simple rolling sum?
- Why store the *earliest* index of each prefix sum for the longest-subarray variant?
- State Kadane's recurrence in one line and explain the "restart when negative" rule.
