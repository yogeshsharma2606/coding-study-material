# 09 - Monotonic Stack

> Code: [`monotonic_stack.go`](monotonic_stack.go) - Tests: [`monotonic_stack_test.go`](monotonic_stack_test.go) - run `go test ./categories/09-monotonic-stack/`

## Overview & mental model

A **monotonic stack** keeps its contents sorted (strictly increasing or decreasing) as you scan. Before pushing the current element, you **pop every element that breaks the order**. The insight: the element that causes a pop is exactly the popped element's **next greater / next smaller** neighbor. Because each element is pushed and popped at most once, an entire family of problems that look O(n^2) run in **O(n)**.

Store **indices** (not values) so you can compute distances/widths.

## How to recognize it

- "**Next/previous greater or smaller** element", "how many days until warmer", "span".
- Histogram/skyline areas, "largest rectangle", trapping rain water.
- "Sum/count over all subarrays of their min/max" (contribution counting).
- "Smallest number after removing k digits" (greedy monotonic construction).

## How to think / attack plan

1. For each element, what neighbor am I looking for? Next greater? Previous smaller? That sets the stack direction:
   - **Next greater** -> pop while `top < current` (decreasing stack).
   - **Next smaller** -> pop while `top > current` (increasing stack).
2. Do I need the value or the **distance/width**? If width, push indices.
3. For "all subarrays" sums, think **contribution**: how many subarrays does each element dominate as the min/max? Use previous-smaller and next-smaller boundaries.
4. Handle duplicates by making one boundary strict and the other non-strict to avoid double counting.

## Core template - next greater to the right

```go
res := make([]int, n)         // default -1
var st []int                   // indices, values decreasing
for i, x := range a {
    for len(st) > 0 && a[st[len(st)-1]] < x {
        res[st[len(st)-1]] = x  // (or i - idx for distance)
        st = st[:len(st)-1]
    }
    st = append(st, i)
}
```

---

## Problems (easy -> hard)

### 1. Next Greater Element I (easy)
**Thinking.** Precompute next-greater for every value in `nums2` with one decreasing-stack pass into a map, then answer `nums1` queries by lookup.
**Complexity.** Time O(n+m), space O(n).

### 2. Next Greater Element II - circular (medium)
**Direction of thinking.** Wrap-around means an element's next-greater might be earlier in the array. Simulate two passes by iterating `2n` times with `i%n`; only push indices during the first pass.
**Dry run.** `[1,2,1]` -> `[2,-1,2]` (the last 1 wraps to the leading 2).
**Complexity.** Time O(n), space O(n).

### 3. Daily Temperatures (medium)
**Direction of thinking.** For each day, the answer is the distance to the next warmer day = next-greater by index. Decreasing stack of indices; on a warmer day, pop and record `i - j`.
**Complexity.** Time O(n), space O(n).

### 4. Largest Rectangle in Histogram (hard)
**Direction of thinking.** A bar's maximal rectangle extends left/right until it hits a shorter bar. Keep an **increasing** stack; when a shorter bar arrives, pop and treat the popped bar's height as the rectangle height, with width bounded by the new index and the new stack top. A trailing sentinel `0` flushes everything.
**Dry run.** `[2,1,5,6,2,3]` -> best 10 (heights 5,6 over width 2).
**Complexity.** Time O(n), space O(n).

### 5. Trapping Rain Water - stack version (hard)
**Direction of thinking.** A decreasing stack; when a taller bar appears, the popped "valley" is bounded by the new bar and the one now on top. Add water layer by layer: `width * (min(leftWall, rightWall) - valley)`.
**Complexity.** Time O(n), space O(n). (The two-pointer version in [03](../03-two-pointers/) is O(1) space.)

### 6. Sum of Subarray Minimums (medium-hard)
**Direction of thinking.** Instead of enumerating subarrays, count each element's **contribution**: `arr[i]` is the minimum of `left[i] * right[i]` subarrays, where `left`/`right` are distances to the previous/next smaller elements. Use `>=` on one side and `>` on the other so duplicates aren't counted twice.
**Dry run.** `[3,1,2,4]` -> 17.
**Complexity.** Time O(n), space O(n).

### 7. Remove K Digits (medium)
**Direction of thinking.** To minimize the number, greedily remove a digit whenever it is larger than the next one (a "descent"). A non-decreasing stack does exactly this; if budget remains, trim from the end; strip leading zeros.
**Dry run.** `"1432219", k=3` -> pop 4,3,2 -> "1219".
**Complexity.** Time O(n), space O(n).

### 8. Online Stock Span (medium, design)
**Direction of thinking.** The span is how many consecutive prior prices were <= today. A decreasing stack of `(price, span)`; collapse popped spans into the current one for O(1) amortized per query.
**Complexity.** Amortized O(1) per `Next`, space O(n).

---

## Common pitfalls & edge cases

- Storing values when you need indices for distance/width.
- Wrong strictness (`<` vs `<=`) causes duplicate handling bugs (Sum of Subarray Minimums).
- Forgetting the sentinel that flushes the stack in histogram problems.
- Off-by-one in width: `i - st.top - 1` after popping.

## Interview Q&A

- **Why is it O(n) despite the inner while loop?** Each index is pushed once and popped once -> total operations <= 2n.
- **Increasing vs decreasing stack - how to choose?** Decreasing finds next-greater; increasing finds next-smaller. Decide from what neighbor you need.
- **Why store indices?** To compute distances/widths and to index back into the array.
- **How does the contribution trick avoid double counting?** Make the previous-smaller boundary strict and the next-smaller non-strict (or vice versa) so equal values are attributed to exactly one element.
