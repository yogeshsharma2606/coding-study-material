# 23 - Matrix / Grid

> Code: [`matrix.go`](matrix.go) - Tests: [`matrix_test.go`](matrix_test.go) - run `go test ./categories/23-matrix-grid/`

## Overview & mental model

A matrix is a 2D array indexed `[row][col]`. Recurring techniques:

- **Rotation** = **transpose then reverse** (or reverse then transpose), no extra matrix.
- **Spiral / layer** traversal: shrink four boundaries (top/bottom/left/right) as you peel edges.
- **In-place flags**: reuse the first row/column (or a spare bit) as O(1) marker storage.
- **Grid as a graph**: cells connect to neighbors -> DFS/BFS for flood fill, islands, shortest path.
- **Sorted structure**: a staircase walk from a corner searches sorted rows+columns in O(m+n).

## How to recognize it

- Input is a 2D slice; "rotate", "spiral", "diagonal", "set zeroes", "search sorted matrix".
- "Regions", "flood", "spread" on a grid -> graph traversal (see [16 - Graphs](../16-graphs/)).
- "In place" / "O(1) extra space" -> encode state in the matrix itself.

## How to think / attack plan

1. Is this really a **graph** problem on a grid? If it's about connectivity/paths, treat neighbors as edges.
2. For transformations, look for an algebraic decomposition (rotate = transpose + reverse).
3. For "in place with O(1) space", find spare capacity: unused rows/cols, or extra bits per cell.
4. Sorted rows/columns unlock binary search or the staircase walk.
5. Mind the boundaries and be careful with `[row][col]` vs `[x][y]` conventions.

---

## Problems (easy -> hard)

### 1. Rotate Image (medium)
**Direction of thinking.** Rotating 90 clockwise sends column j to row j reversed. Decompose: **transpose** (mirror over the main diagonal), then **reverse each row**. Both are in-place.
**Dry run.** `[[1,2,3],[4,5,6],[7,8,9]]` -> transpose -> reverse rows -> `[[7,4,1],[8,5,2],[9,6,3]]`.
**Complexity.** Time O(n^2), space O(1).

### 2. Spiral Matrix (medium)
**Direction of thinking.** Maintain four boundaries; traverse top row L->R, right column T->B, bottom R->L, left B->T, shrinking each boundary after use. Guard the last two edges to avoid re-traversing a single remaining row/column.
**Complexity.** Time O(mn), space O(1) extra.

### 3. Set Matrix Zeroes (medium)
**Direction of thinking.** Naive uses O(m+n) marker arrays. To get O(1), use the **first row and column** as those markers; track their own zero-state in two booleans, then apply, then fix the first row/col last.
**Complexity.** Time O(mn), space O(1).

### 4. Transpose Matrix (easy)
**Thinking.** `out[c][r] = m[r][c]`. For non-square matrices you need a new matrix (dimensions swap).
**Complexity.** Time O(mn), space O(mn).

### 5. Search a 2D Matrix (medium)
**Direction of thinking.** Rows are sorted and each row starts after the previous ends, so the matrix is one sorted sequence. Binary search over `[0, m*n)` mapping `mid -> [mid/cols][mid%cols]`.
**Complexity.** Time O(log(mn)), space O(1).

### 6. Search a 2D Matrix II (medium)
**Direction of thinking.** Rows and columns are each sorted but rows don't chain. Start at the **top-right** corner: the current value is the largest in its row and smallest in its column, so `> target` -> move left, `< target` -> move down. Each step drops a row or column.
**Complexity.** Time O(m+n), space O(1).

### 7. Flood Fill (easy)
**Thinking.** DFS/BFS from the seed, recoloring all 4-connected cells of the original color. Guard against infinite recursion when new==old color.
**Complexity.** Time O(mn), space O(mn) recursion worst case.

### 8. Game of Life (medium)
**Direction of thinking.** All cells update **simultaneously** from the old state, so you can't overwrite in place naively. Encode the new state in a **second bit** while neighbor counts still read the old bit; then shift every cell right by 1. O(1) extra space.
**Complexity.** Time O(mn), space O(1).

### 9. Diagonal Traverse (medium)
**Direction of thinking.** Cells on diagonal d share `r+c = d`. Alternate direction per diagonal (up-right on even d, down-left on odd), clamping the start to the valid corner.
**Complexity.** Time O(mn), space O(1) extra.

---

## Common pitfalls & edge cases

- Non-square matrices for rotate/transpose (dimensions change - can't do in place).
- Spiral: re-traversing the middle row/column when only one remains - guard with the boundary checks.
- Game of Life: updating in place without encoding both states corrupts neighbor counts.
- Off-by-one and row/col swaps - keep the `[row][col]` convention consistent.
- Empty matrix / empty rows.

## Interview Q&A

- **How to rotate a matrix in place?** Transpose then reverse each row (clockwise), or reverse rows then transpose (counter-clockwise).
- **How to get O(1) space for Set Zeroes?** Use the first row/column as marker storage, tracking their own state separately.
- **Why start at the top-right for the sorted-matrix search?** It's the unique corner where one direction increases and the other decreases, giving a clean elimination each step.
- **How do you update Game of Life simultaneously in place?** Store the next state in an extra bit so the current bit still reflects the old generation during counting.
