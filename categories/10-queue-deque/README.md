# 10 - Queue & Deque

> Code: [`queue_deque.go`](queue_deque.go) - Tests: [`queue_deque_test.go`](queue_deque_test.go) - run `go test ./categories/10-queue-deque/`

## Overview & mental model

A **queue** is FIFO - first in, first out. It's the backbone of **BFS** (level-order exploration) and **stream/window** processing (evict the oldest). A **deque** (double-ended queue) supports push/pop at both ends; its killer application is the **monotonic deque**, which yields sliding-window max/min in **O(n)** by keeping only elements that could still become the answer and discarding dominated ones.

In Go there's no built-in deque, but a slice works: `append` to push back, `s[1:]` to pop front, `s[:len(s)-1]` to pop back. (`container/list` is available but slices are faster for interviews.)

## How to recognize it

- "Process in arrival order", "level by level", "recent N", "last k".
- **Sliding-window max/min** in linear time -> monotonic deque.
- "Implement queue with stacks" / ring buffer / circular queue design questions.
- DP where each state depends on the best of a moving window -> deque optimization.

## How to think / attack plan

1. Do I always remove the **oldest**? Plain queue. Do I need both ends? Deque.
2. For window max/min: maintain a deque of **indices** whose values are monotonic. The front is the current extreme; pop dominated values from the back, evict out-of-window indices from the front.
3. For "queue from stacks", amortize: reverse lazily only when the output stack empties.
4. For fixed capacity, a **ring buffer** (array + head + size, modular indexing) gives O(1) with no shifting.

## Core template - sliding-window maximum

```go
var dq []int  // indices, values decreasing
for i, x := range a {
    if len(dq) > 0 && dq[0] <= i-k { dq = dq[1:] }        // evict old
    for len(dq) > 0 && a[dq[len(dq)-1]] <= x { dq = dq[:len(dq)-1] }
    dq = append(dq, i)
    if i >= k-1 { out = append(out, a[dq[0]]) }
}
```

---

## Problems (easy -> hard)

### 1. Implement Queue using Stacks (easy)
**Direction of thinking.** A stack reverses order; two stacks reverse twice = original FIFO order. Push to `in`; when you need the front, if `out` is empty, pour `in` into `out`. Each element moves at most once between stacks -> **amortized O(1)**.
**Complexity.** Amortized O(1) per op.

### 2. Moving Average from Data Stream (easy)
**Thinking.** Keep a window and a running sum; on overflow, subtract the front and drop it.
**Complexity.** O(1) per `Next`.

### 3. Number of Recent Calls (easy)
**Thinking.** Pings are increasing; keep a queue of timestamps and evict anything older than `t-3000` from the front.
**Complexity.** Amortized O(1) per `Ping`.

### 4. Design Circular Queue (medium)
**Direction of thinking.** A fixed array with `head` and `size`; the tail index is `(head+size) % cap`. Modular arithmetic avoids shifting elements. Full when `size == cap`.
**Complexity.** All ops O(1), space O(k).

### 5. Sliding Window Maximum (hard) - the signature deque problem
**Direction of thinking.** Recomputing each window's max is O(nk). Keep a **decreasing** deque of indices: the front holds the window's max. When a new value arrives, pop smaller values from the back (they're dominated and can never be the max while the newcomer is in the window). Evict indices that slid out of the window from the front.
**Why it's O(n).** Each index is pushed and popped once.
**Dry run.** `[1,3,-1,-3,5,3,6,7], k=3` -> `[3,3,5,5,6,7]`.
**Complexity.** Time O(n), space O(k).

### 6. Shortest Subarray with Sum at Least K (hard)
**Direction of thinking.** With negatives, sliding window fails. Use prefix sums; you want the closest earlier prefix `P[j]` with `P[i]-P[j] >= k`. Keep an **increasing** deque of prefix indices: pop the front when it satisfies (record length), and pop the back while it's >= current prefix (larger earlier prefixes are useless).
**Complexity.** Time O(n), space O(n).

### 7. Jump Game VI (medium-hard) - deque-optimized DP
**Direction of thinking.** `dp[i] = nums[i] + max(dp[i-k..i-1])`. The inner max over a moving window is a sliding-window maximum -> monotonic deque turns O(nk) into O(n).
**Dry run.** `[1,-1,-2,4,-7,3], k=2` -> 7.
**Complexity.** Time O(n), space O(n).

---

## Common pitfalls & edge cases

- `s[1:]` to pop the front does not reclaim memory immediately; fine for interviews, but note it for long-lived queues.
- Monotonic deque: store **indices** so you can evict out-of-window elements.
- Strictness (`<=` vs `<`) when popping the back affects handling of equal values.
- Circular queue: full vs empty ambiguity - track `size` explicitly to disambiguate.

## Interview Q&A

- **Why two stacks for a queue and what's the amortized cost?** Each element is transferred at most once from `in` to `out`, so O(1) amortized despite occasional O(n) transfers.
- **Why is the monotonic-deque window max O(n)?** Every index enters and leaves the deque exactly once.
- **Queue vs deque?** Queue removes only from the front; a deque removes/adds at both ends, enabling window extrema.
- **When BFS vs DFS?** BFS (queue) for shortest path in unweighted graphs / level order; DFS (stack/recursion) for connectivity, cycles, and full exploration.
