# 14 - Trees

> Code: [`trees.go`](trees.go) - Tests: [`trees_test.go`](trees_test.go) - run `go test ./categories/14-trees/`

## Overview & mental model

A binary tree is recursion incarnate: the answer for a node is built from the answers for its subtrees. The single most useful question is:

> **What do I need FROM my children, and what do I RETURN to my parent?**

Traversal order encodes the strategy:
- **Preorder** (root, left, right) - top-down; copying, serialization, prefix work.
- **Inorder** (left, root, right) - for a **BST this yields sorted order**.
- **Postorder** (left, right, root) - bottom-up aggregates: height, diameter, "is balanced", subtree sums.
- **BFS / level-order** (queue) - shortest depth, level grouping, right-side view.

## How to recognize it

- Input is a `*TreeNode`; "depth", "path", "ancestor", "level", "validate", "serialize".
- "Sorted"/"kth" + BST -> inorder.
- "Shortest depth"/"level"/"width" -> BFS.
- "Aggregate from leaves up" (height, diameter, balance) -> postorder DFS.

## How to think / attack plan

1. Choose DFS (recursion) or BFS (queue) based on whether you need depth-first structure or level order.
2. For DFS, decide the traversal order by what info flows where (down = arguments; up = return value).
3. BST? Exploit the ordering - inorder is sorted; comparisons prune half the tree.
4. Watch base cases (`nil`) and single-node trees.

## Core template - "return to parent" DFS

```go
func solve(n *TreeNode) retType {
    if n == nil { return base }
    l := solve(n.Left)
    r := solve(n.Right)
    // combine l, r, n.Val; optionally update a global answer
    return combined
}
```

---

## Problems (easy -> hard)

### 1. Maximum Depth (easy)
**Thinking.** Postorder: `1 + max(depth(left), depth(right))`. The cleanest possible "combine children" example.
**Complexity.** Time O(n), space O(h) recursion (h = height).

### 2. Inorder Traversal (easy)
**Direction of thinking.** Recursion is one line; the **iterative** version (explicit stack) shows how traversal really works and avoids stack-overflow on skewed trees. Push all left children, visit, then go right.
**Complexity.** Time O(n), space O(h).

### 3. Level Order Traversal (medium)
**Direction of thinking.** BFS with a queue. Snapshot the queue size at the start of each level to group nodes by depth.
**Complexity.** Time O(n), space O(width).

### 4. Validate BST (medium)
**Direction of thinking.** A common wrong answer checks only `left < root < right` locally. The real invariant is that each node lies within an inherited `(low, high)` range from all ancestors. Pass the tightening bounds down.
**Dry run.** Node 6 under root 5's left subtree violates the upper bound `< 5` -> invalid.
**Complexity.** Time O(n), space O(h).

### 5. Lowest Common Ancestor (medium)
**Direction of thinking.** Postorder search: return a node upward if its subtree contains p or q. If both children return non-nil, the current node is where the paths split -> it's the LCA.
**Complexity.** Time O(n), space O(h).

### 6. Diameter of Binary Tree (medium)
**Direction of thinking.** The longest path through a node is `leftHeight + rightHeight`. Compute heights bottom-up and update a global max at each node - the return value (height) differs from the tracked answer (diameter).
**Complexity.** Time O(n), space O(h).

### 7. Invert Binary Tree (easy)
**Thinking.** Swap children recursively - the famous "can you invert a binary tree?" question.
**Complexity.** Time O(n), space O(h).

### 8. Kth Smallest in a BST (medium)
**Direction of thinking.** Inorder visits BST nodes in ascending order; stop at the kth. No full traversal needed.
**Complexity.** Time O(h + k), space O(h).

### 9. Binary Tree Right Side View (medium)
**Direction of thinking.** BFS and take the last node of each level (or DFS visiting right first, recording the first node seen per depth).
**Complexity.** Time O(n), space O(width).

### 10. Serialize and Deserialize (hard)
**Direction of thinking.** Preorder with explicit **nil markers** (`#`) uniquely encodes structure; deserialize by consuming tokens in the same preorder, recursing left then right. This is how you persist/transmit a tree.
**Complexity.** Time O(n), space O(n).

---

## Common pitfalls & edge cases

- Validating a BST with only local parent-child comparisons (miss deep violations) - use inherited bounds or inorder-monotonicity.
- Confusing the DFS **return value** with a **global** answer (diameter, max path sum).
- Deep/skewed trees blow the recursion stack - mention iterative traversal.
- Forgetting nil markers in serialization -> ambiguous structure.

## Interview Q&A

- **Which traversal for a BST to get sorted output?** Inorder.
- **BFS vs DFS on trees?** BFS for level/shortest-depth; DFS for path/aggregate/structure problems.
- **How to avoid stack overflow on huge trees?** Iterative traversal with an explicit stack, or Morris traversal (O(1) space).
- **Why do bounds validate a BST but local checks don't?** A node deep in the left subtree must still be less than a far ancestor, which local checks ignore.
