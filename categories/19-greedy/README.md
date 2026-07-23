# 19 - Greedy

> Code: [`greedy.go`](greedy.go) - Tests: [`greedy_test.go`](greedy_test.go) - run `go test ./categories/19-greedy/`

## Overview & mental model

A greedy algorithm makes the **locally optimal choice at each step and never looks back**. When valid, it's simple and fast (O(n) or O(n log n)). The catch: greed is only **correct** when the problem has the **greedy-choice property** - a series of local optima yields the global optimum. For many problems it doesn't (Coin Change with arbitrary coins), and you need DP instead.

The senior differentiator is **justifying** greed, usually via an **exchange argument**: show that any optimal solution can be transformed into the greedy one without getting worse.

## How to recognize it

- "Maximum/minimum number of ...", "can you reach ...", "minimum steps/intervals".
- Sorting the input makes an obvious best-first choice available.
- Each decision is independent enough that a local best doesn't sabotage the future.
- Contrast: if a locally worse choice can enable a better overall result, it's DP, not greedy.

## How to think / attack plan

1. Propose the greedy rule ("always pick the earliest-ending", "extend the farthest reach", "take every profitable step").
2. **Stress-test with a small counterexample.** If you find one, greedy is wrong -> switch to DP.
3. If it survives, give an **exchange argument** for why the greedy choice is safe.
4. Often the right greedy rule appears only **after sorting** by the correct key.

## How greedy differs from DP

| | Greedy | DP |
|---|---|---|
| Choice | Commit to local best, never revisit | Explore/combine all sub-choices |
| Speed | Usually O(n) / O(n log n) | Often O(n^2)+ |
| Correctness | Needs greedy-choice property | Always correct if state/transition right |

---

## Problems (easy -> hard)

### 1. Jump Game (medium)
**Direction of thinking.** Track the **farthest reachable** index. If your current index ever exceeds it, you can't proceed. No need to try specific jumps - reachability is monotonic.
**Complexity.** Time O(n), space O(1).

### 2. Jump Game II - min jumps (medium)
**Direction of thinking.** BFS by "jump level": within the current jump's reach, compute the farthest next reach; when you hit the current boundary, you must spend a jump and extend to that farthest. Greedy = always jump to whatever gets you farthest next.
**Dry run.** `[2,3,1,1,4]` -> jump to index 1 (reach 4), then to end -> 2 jumps.
**Complexity.** Time O(n), space O(1).

### 3. Gas Station (medium)
**Direction of thinking.** If total gas >= total cost, a start exists and is **unique** among candidates. Track a running tank; whenever it goes negative at station i, no start in `[start..i]` works, so reset the start to `i+1`. One pass.
**Why the reset is safe.** If you couldn't reach i from `start`, no intermediate station (which starts with even less surplus) could either.
**Complexity.** Time O(n), space O(1).

### 4. Assign Cookies (easy)
**Direction of thinking.** Sort children by greed and cookies by size; give each child the **smallest cookie that satisfies** them. Wasting a big cookie on a low-greed child would be suboptimal.
**Complexity.** Time O(n log n).

### 5. Best Time to Buy and Sell Stock II (medium)
**Direction of thinking.** With unlimited transactions, the max profit is the sum of every **positive** consecutive difference - capture each up-move. Equivalent to buying before every rise.
**Complexity.** Time O(n), space O(1).

### 6. Partition Labels (medium)
**Direction of thinking.** Each letter must be fully contained in one part, so a part must extend to the **last occurrence** of every letter it contains. Precompute last-index; sweep, extending `end`; cut when `i == end`.
**Dry run.** `"ababcbaca..."` -> first part ends where the last 'a'/'b'/'c' occurs -> size 9.
**Complexity.** Time O(n), space O(1).

### 7. Task Scheduler (medium)
**Direction of thinking.** The schedule length is dictated by the **most frequent task**: it needs `(maxCount-1)` cooldown frames of size `(n+1)`, plus a final slot for each task tied at max frequency. If other tasks fill all idle slots, the answer is just the task count.
**Dry run.** `AAABBB, n=2` -> `A B _ A B _ A B` = 8.
**Complexity.** Time O(tasks), space O(1).

### 8. Candy (hard)
**Direction of thinking.** Each child needs more than a lower-rated neighbor - on **both** sides. A single pass can't satisfy both, so do two: left-to-right enforces the left constraint, right-to-left the right, and take the max at each position.
**Dry run.** `[1,0,2]` -> left `[1,1,2]`, right merge `[2,1,2]` -> 5.
**Complexity.** Time O(n), space O(n).

---

## Common pitfalls & edge cases

- Applying greedy where DP is required (arbitrary-coin change, 0/1 knapsack) - always sanity-check with a counterexample.
- Wrong sort key - the greedy choice is only exposed by the right ordering.
- Off-by-one in "reach" tracking (Jump Game) and boundary resets (Gas Station).
- Candy: forgetting to also compare against the existing value in the second pass (`take the max`).

## Interview Q&A

- **How do you know greedy is correct?** Prove the greedy-choice property, typically via an exchange argument; if you can construct a counterexample, it isn't.
- **Greedy vs DP?** Greedy commits to a local optimum (fast, needs proof); DP explores all sub-choices (always correct with the right formulation).
- **Why is Gas Station's greedy start unique/valid?** Total feasibility guarantees a start, and the reset skips all provably-impossible starts in one pass.
- **What role does sorting play?** It surfaces the "best next" element so the greedy choice becomes obvious and provably safe.
