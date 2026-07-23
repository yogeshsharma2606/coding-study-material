# 18 - Dynamic Programming

> Code: [`dp.go`](dp.go) - Tests: [`dp_test.go`](dp_test.go) - run `go test ./categories/18-dynamic-programming/`

## Overview & mental model

DP solves problems that have **overlapping subproblems** (the same smaller problem recurs) and **optimal substructure** (the optimal answer is built from optimal sub-answers). You cache each subproblem's answer so it's computed once, turning exponential recursion into polynomial time.

The reliable recipe:

1. **Define the state** - what do the indices/parameters of `dp` mean? (e.g. "`dp[i]` = best answer using the first i items").
2. **Write the transition** - how does a state depend on smaller states?
3. **Base cases** - the smallest states' values.
4. **Order of evaluation** - top-down (memoized recursion) or bottom-up (fill a table).
5. **Optimize space** - if `dp[i]` only needs `dp[i-1]`/`dp[i-2]`, keep rolling variables.

## How to recognize it

- "Number of ways", "min/max cost", "can you reach/achieve", "longest/shortest ... subsequence".
- A greedy choice would be wrong because a locally worse choice can be globally better.
- The brute-force recursion re-solves the same inputs (draw the recursion tree - repeats mean DP).

## How to think / attack plan

1. Start with the **brute-force recursion** ("try every choice"). Its parameters become your DP **state**.
2. Add **memoization** to kill repeated work (top-down). This alone often suffices in interviews.
3. If asked, convert to **bottom-up** by iterating states in dependency order, then **reduce space**.
4. Classic state shapes: 1D over an index, 2D over two strings/indices, "knapsack" over items x capacity.

## Common DP families (recognize the shape)

| Family | State | Example |
|---|---|---|
| Linear / Fibonacci | dp[i] from dp[i-1], dp[i-2] | Climb Stairs, House Robber |
| Unbounded knapsack | dp[amount] | Coin Change |
| 0/1 knapsack | dp[capacity], iterate items | Partition Equal Subset |
| Two sequences (grid) | dp[i][j] | LCS, Edit Distance |
| Subsequence | dp[i] = best ending at i | LIS |
| Grid paths | dp[i][j] | Unique Paths |

---

## Problems (easy -> hard)

### 1. Climbing Stairs (easy)
**Direction of thinking.** To reach step n you came from n-1 or n-2, so `ways(n)=ways(n-1)+ways(n-2)` - Fibonacci. Two rolling variables give O(1) space.
**Complexity.** Time O(n), space O(1).

### 2. House Robber (medium)
**Direction of thinking.** At house i, either skip it (`dp[i-1]`) or rob it and skip the neighbor (`dp[i-2]+nums[i]`). Take the max. The "can't be adjacent" constraint is the whole recurrence.
**Complexity.** Time O(n), space O(1).

### 3. Coin Change (medium) - unbounded knapsack
**Direction of thinking.** `dp[a]` = fewest coins for amount a; try each coin as the last one: `dp[a]=1+min(dp[a-c])`. Coins reusable -> iterate amounts ascending.
**Dry run.** coins {1,2,5}, amount 11 -> 5+5+1 = 3 coins.
**Complexity.** Time O(amount * #coins), space O(amount).

### 4. Longest Increasing Subsequence (medium)
**Direction of thinking.** O(n^2): `dp[i]` = LIS ending at i. **O(n log n)** (shown in code): keep `tails[k]` = smallest possible tail of an increasing subsequence of length k+1; binary-search where each number extends/replaces. The length of `tails` is the answer.
**Dry run.** `[10,9,2,5,3,7,101,18]` -> 4 (`2,3,7,101`).
**Complexity.** Time O(n log n), space O(n).

### 5. Longest Common Subsequence (medium) - two-sequence DP
**Direction of thinking.** `dp[i][j]` on prefixes: if last chars match, `dp[i-1][j-1]+1`; else the best of dropping one char from either string. The canonical 2D grid DP.
**Complexity.** Time O(mn), space O(mn) (reducible to O(min(m,n))).

### 6. 0/1 Knapsack (medium)
**Direction of thinking.** Each item is taken or not. 1D `dp[c]` over capacity, iterated **downward** so each item contributes at most once (upward iteration would allow reuse = unbounded knapsack).
**Why downward.** Prevents using an item's updated value within the same item's loop.
**Complexity.** Time O(n * capacity), space O(capacity).

### 7. Unique Paths (medium) - grid DP
**Direction of thinking.** Paths to a cell = paths from above + paths from the left. One row rolled left-to-right suffices.
**Complexity.** Time O(mn), space O(n).

### 8. Word Break (medium)
**Direction of thinking.** `dp[i]` = can `s[:i]` be segmented. For each i, look back to a split `j` where `dp[j]` is true and `s[j:i]` is a word. Cap the look-back by the longest dictionary word.
**Complexity.** Time O(n * maxWordLen), space O(n).

### 9. Edit Distance (hard)
**Direction of thinking.** `dp[i][j]` = edits to turn `a[:i]` into `b[:j]`. If chars match, no cost (diagonal); else 1 + min(replace=diag, delete=up, insert=left). Base cases are empty-to-prefix (all inserts/deletes).
**Dry run.** "horse"->"ros" = 3.
**Complexity.** Time O(mn), space O(mn) (reducible to O(n)).

### 10. Partition Equal Subset Sum (medium)
**Direction of thinking.** Two equal halves exist iff a subset sums to total/2 - a boolean 0/1 knapsack. Odd total is immediately impossible.
**Complexity.** Time O(n * sum), space O(sum).

---

## Common pitfalls & edge cases

- Wrong iteration direction in 1D knapsack (down for 0/1, up for unbounded).
- Off-by-one between "index into the array" and "prefix length" in 2D DP (the `+1` padding).
- Using DP where **greedy** suffices (slower) or greedy where DP is required (wrong).
- Forgetting base cases (`dp[0]`) or initializing min-DP to a small value instead of infinity.

## Interview Q&A

- **Top-down vs bottom-up?** Same complexity; top-down (memoized recursion) is easier to derive from brute force, bottom-up avoids recursion overhead and enables space reduction.
- **How do you find the DP state?** It's the set of parameters that fully describe a subproblem in your brute-force recursion.
- **When does greedy fail and DP is needed?** When a locally optimal choice can prevent the global optimum (e.g. Coin Change with arbitrary denominations).
- **How to reconstruct the actual solution (not just its value)?** Store choices/parents and backtrack through the table, or recompute from the DP.
