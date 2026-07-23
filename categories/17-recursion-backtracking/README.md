# 17 - Recursion & Backtracking

> Code: [`backtracking.go`](backtracking.go) - Tests: [`backtracking_test.go`](backtracking_test.go) - run `go test ./categories/17-recursion-backtracking/`

## Overview & mental model

Backtracking is a disciplined brute force over a **decision tree**. At each step you make a **choice**, **recurse** to explore its consequences, then **undo** the choice and try the next. The universal skeleton is:

```
choose  ->  explore (recurse)  ->  unchoose (backtrack)
```

The art is **pruning**: cut off branches that can't lead to a valid/optimal answer as early as possible, so you explore far fewer than the theoretical maximum.

## How to recognize it

- "Generate **all** ...", "list every ...", "count the number of ways to ...", "find **any** valid ...".
- Combinations, permutations, subsets, partitions.
- Constraint satisfaction: N-Queens, Sudoku, word search in a grid.
- The answer space is exponential (n <= ~20), and you must enumerate/validate configurations.

## How to think / attack plan

1. What is a single **decision** at each step, and what are the choices?
2. What is the **base case** (a complete, valid solution)?
3. What **state** must I carry (current path, used set, remaining target)? Remember to record a **copy** of mutable state when saving a solution.
4. What **pruning** applies (sorted candidates, constraint checks, "not enough elements left")?
5. To avoid duplicates with repeated inputs: sort, and skip `nums[i] == nums[i-1]` at the same tree depth.

## Core template

```go
func backtrack(state) {
    if isComplete(state) { record(copy(state)); return }
    for _, choice := range choices(state) {
        if !valid(choice) { continue }   // prune
        apply(choice)
        backtrack(state)
        undo(choice)
    }
}
```

---

## Problems (easy -> hard)

### 1. Subsets (medium)
**Direction of thinking.** For each element, two choices: in or out. Record the current path at every node (each node is a subset). Using a `start` index prevents revisiting earlier elements (avoids permuted duplicates).
**Complexity.** Time O(n * 2^n), space O(n) recursion.

### 2. Permutations (medium)
**Direction of thinking.** Each position picks an **unused** element; a `used[]` array tracks availability. There are n! leaves.
**Complexity.** Time O(n * n!), space O(n).

### 3. Combinations of k from n (medium)
**Direction of thinking.** Like subsets but stop at size k. Prune when the remaining numbers can't fill the rest: `i <= n-(k-len(cur))+1`.
**Complexity.** Time O(k * C(n,k)).

### 4. Combination Sum (medium)
**Direction of thinking.** Candidates can be reused, so recurse with the **same index** `i` (not `i+1`). Sort first so you can `break` once a candidate exceeds the remaining target.
**Dry run.** target 7 from [2,3,6,7] -> [2,2,3] and [7].
**Complexity.** Time exponential in the target/candidates; pruning helps a lot.

### 5. Generate Parentheses (medium)
**Direction of thinking.** Build the string char by char with two counters. **Constraint pruning**: add `(` only while `open < n`; add `)` only while `close < open`. This only ever builds valid strings - no post-filtering. The count is the Catalan number.
**Complexity.** Time O(4^n / sqrt(n)) (Catalan).

### 6. Letter Combinations of a Phone Number (medium)
**Direction of thinking.** Each digit maps to letters; recurse digit by digit, appending each letter. A product of choices.
**Complexity.** Time O(4^n) worst (7,9 have 4 letters).

### 7. Palindrome Partitioning (medium)
**Direction of thinking.** Try every prefix `s[start:end]`; if it's a palindrome, fix it and recurse on the rest. Backtrack to try longer prefixes.
**Complexity.** Time O(n * 2^n) worst.

### 8. Word Search (medium) - grid backtracking
**Direction of thinking.** DFS from each cell matching `word[idx]`, moving to 4 neighbors, marking cells visited **in place** and restoring on the way out (so other paths can reuse them). Return on the first full match.
**Complexity.** Time O(rows*cols*4^L), space O(L).

### 9. N-Queens (hard)
**Direction of thinking.** Place one queen per **row** (rows can't clash by construction). Track threatened **columns** and both **diagonals** (`r+c` and `r-c`) in boolean arrays for O(1) validity. Recurse row by row; record boards at row n.
**Why the diagonal indices.** All cells on a `/` diagonal share `r+c`; on a `\` diagonal share `r-c` (offset by n to stay non-negative).
**Complexity.** Time O(n!) with heavy pruning.

---

## Common pitfalls & edge cases

- Saving a **reference** to the mutable path instead of a copy - all results end up identical. Always copy on record.
- Forgetting to undo state (leaving `used[i]=true`) corrupts sibling branches.
- Duplicate results with repeated inputs - sort and skip same-value siblings at the same depth.
- Missing pruning turns a feasible problem into a timeout.

## Interview Q&A

- **Backtracking vs plain recursion?** Backtracking explicitly undoes choices to reuse the same state buffer while exploring alternatives.
- **How to prune effectively?** Sort to enable early `break`, check partial constraints before recursing, and bound the remaining possibilities.
- **How to avoid duplicate combinations/permutations?** Use a `start` index (combinations) or skip equal siblings after sorting (with duplicates).
- **Time complexity of enumerating subsets/permutations?** 2^n subsets, n! permutations - inherently exponential; that's expected for "generate all".
