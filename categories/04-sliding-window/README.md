# 04 - Sliding Window

> Code: [`sliding_window.go`](sliding_window.go) - Tests: [`sliding_window_test.go`](sliding_window_test.go) - run `go test ./categories/04-sliding-window/`

## Overview & mental model

A sliding window is a contiguous range `[left, right]` you slide across an array/string, incrementally updating an aggregate (sum, count, frequency map) as elements enter on the right and leave on the left. It reuses overlap between adjacent subarrays, turning O(n^2)/O(n*k) brute force into O(n).

Two shapes:
- **Fixed window** - size `k` is given; add the new element, drop the one `k` behind.
- **Variable window** - grow `right` to satisfy a goal, then shrink `left` while a constraint is (or isn't) met, tracking the best.

## How to recognize it

- "**Contiguous** subarray/substring" plus an optimization (longest/shortest/max/min) or a validity condition.
- "At most k distinct", "sum >= target", "no repeating chars", "containing all of t".
- Signals it is NOT sliding window: subsequence (non-contiguous), or the array has negatives and you need "sum == k" over subarrays (then use prefix sum + hashmap, because shrinking is not monotonic).

## How to think / attack plan

1. Is the window **fixed** or **variable**? Fixed -> simple add/drop. Variable -> the "grow then shrink" template.
2. What state summarizes the window? A running sum, a `char->count` map, or a `maxFreq`.
3. What makes the window **valid/invalid**, and is validity **monotonic** (once valid, staying/ shrinking keeps it decidable)? Monotonicity is what lets left only move forward -> O(n).
4. Record the answer at the right moment (when valid for "longest", after satisfying for "shortest").

## Core template - variable window

```go
left := 0
for right := 0; right < len(a); right++ {
    add(a[right])
    for windowInvalid() {   // or: while valid, for "shortest"
        remove(a[left]); left++
    }
    best = update(best, right-left+1)
}
```

---

## Problems (easy -> hard)

### 1. Max Sum Subarray of Size K (easy, fixed)
**Thinking.** Recomputing each window's sum is O(n*k). Instead slide: `sum += nums[r] - nums[r-k]`. O(n).
**Dry run.** `[2,1,5,1,3,2], k=3`: sums 8,7,9,6 -> best 9.
**Complexity.** Time O(n), space O(1).

### 2. Longest Substring Without Repeating Characters (medium)
**Direction of thinking.** Grow right; if the incoming char was seen **inside** the current window, jump `left` to just past its last position. Store last index per char.
**Why the jump.** Moving left one-by-one also works but the jump is cleaner and still O(n).
**Dry run.** `"abcabcbb"`: window grows to `abc` (3); at second `a`, left jumps; max stays 3.
**Complexity.** Time O(n), space O(min(n, alphabet)).

### 3. Minimum Size Subarray Sum (medium)
**Problem.** Shortest subarray with sum >= target.
**Direction of thinking.** Grow right to reach the target, then **shrink from the left** while still >= target, recording the shortest. Requires non-negative numbers (so shrinking monotonically reduces the sum).
**Dry run.** `t=7, [2,3,1,2,4,3]`: window `[4,3]` sums 7 at length 2.
**Complexity.** Time O(n), space O(1).

### 4. Longest Substring with At Most K Distinct (medium)
**Direction of thinking.** Keep a `char->count` map; the number of distinct chars is `len(map)`. When it exceeds k, shrink left, deleting a char when its count hits 0.
**Complexity.** Time O(n), space O(k).

### 5. Longest Repeating Character Replacement (medium)
**Problem.** Longest substring that becomes all-one-character after replacing at most k chars.
**Direction of thinking.** A window is feasible when `windowLen - maxFreq <= k` (the non-dominant chars are the ones you'd replace). Track `maxFreq`; when infeasible, slide left by one. (We don't decrease `maxFreq` on shrink - it can only cause the window to not grow, which is fine and keeps it O(n).)
**Dry run.** `"AABABBA", k=1`: best window length 4.
**Complexity.** Time O(n), space O(1) (26 counts).

### 6. Find All Anagrams in a String (medium, fixed)
**Direction of thinking.** An anagram of `p` is any window of `len(p)` whose letter-frequency array equals p's. Slide a fixed window and compare `[26]int` arrays (comparable in Go).
**Dry run.** `s="cbaebabacd", p="abc"` -> matches at indices 0 and 6.
**Complexity.** Time O(n), space O(1).

### 7. Max Consecutive Ones III (medium)
**Problem.** Longest run of 1s if you may flip at most k zeros.
**Direction of thinking.** A window is valid while it contains <= k zeros. Grow right; when zeros exceed k, shrink left. The window length is the answer.
**Complexity.** Time O(n), space O(1).

### 8. Minimum Window Substring (hard)
**Problem.** Smallest substring of `s` containing every char of `t` (with multiplicity).
**Direction of thinking.** Track how many distinct required chars are fully satisfied (`formed`) vs total required. Expand right until `formed == required`, then **contract left** to minimize while still satisfied. The tricky part is the `formed` bookkeeping - only decrement when a required char drops *below* its needed count.
**Dry run.** `s="ADOBECODEBANC", t="ABC"` -> "BANC".
**Complexity.** Time O(n), space O(alphabet).

---

## Common pitfalls & edge cases

- Using sliding window on arrays with **negatives** for "sum == k" - shrinking isn't monotonic; use prefix sum + hashmap instead.
- Forgetting to `delete` a key when its count hits 0 (breaks the "distinct = len(map)" invariant).
- In "shortest" problems, record the answer **inside** the shrink loop; in "longest", record after making the window valid.
- Fixed-window comparisons: `[26]int` arrays are comparable with `==` in Go - use that.

## Interview Q&A

- **Sliding window vs prefix sum?** Window needs a monotonic validity (usually non-negative values); prefix sum + hashmap handles arbitrary values and exact-sum subarray counts.
- **How is it O(n) if there's an inner loop?** `left` only moves forward across the whole run, so total left-moves <= n; amortized O(n).
- **When fixed vs variable?** Fixed when the size is given; variable when you optimize length under a constraint.
