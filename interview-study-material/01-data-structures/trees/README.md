# Binary Trees

A binary tree is a recursive structure: each node has a value plus a `Left` and `Right` child (either may be `nil`). Because the definition is recursive, **most tree solutions are recursive too** — solve the subtrees, then combine.

**The universal tree template.** Almost every problem here fits this shape:

```
func solve(node) result {
    if node == nil { return base_case }      // e.g. 0, true, nil
    left  := solve(node.Left)
    right := solve(node.Right)
    return combine(node.Val, left, right)     // the only part that changes
}
```

Learn that skeleton and you "never forget" trees — you just change the base case and the `combine` step.

**Traversal cheat code (where does the root print?):**
- **Pre-order** = root **before** children (Root, L, R) → good for *copying/serializing* a tree.
- **In-order** = root **between** children (L, Root, R) → on a **BST this yields sorted order**.
- **Post-order** = root **after** children (L, R, Root) → good for *deleting/freeing* or when the answer depends on children first.
- **Level-order** = breadth-first, uses a **queue**, visits level by level.

## Programs

### [Tree traversals](tree-traversals/tree_traversals.go)

Inorder, preorder, postorder (recursive DFS) and level-order (BFS with a queue).

**How to think about it:** the three DFS orders are *the same three lines* — you only move the `Print(root.Val)` line to change the order. BFS is different: it's iterative with a **queue**; dequeue a node, print it, then enqueue its children left-then-right.

**Dry run** for the tree `A(B(D,E), C)`:

```
In-order   (L,Root,R): D B E A C
Pre-order  (Root,L,R): A B D E C
Post-order (L,R,Root): D E B C A
Level-order (queue)  : A B C D E
```

- All four are **O(n) time**. Recursive DFS uses **O(h) stack** (h = height); BFS uses **O(w) queue** (w = max width).

### [Identical trees](identical-trees/identical_trees.go)

Are two trees structurally identical **and** value-equal?

**How to think about it:** two trees are the same iff their roots match *and* their left subtrees match *and* their right subtrees match — pure recursion. Base cases first: both `nil` → equal; exactly one `nil` (or values differ) → not equal.

- **O(n) time, O(h) space.** This is the template with `combine = (p.Val==q.Val) && sameLeft && sameRight`.

### [Symmetric tree](symmetric-tree/symmetric_tree.go)

Is the tree a mirror image of itself around its center?

**Direction of thinking:** symmetry is *not* "left subtree equals right subtree" — it's "left is the **mirror** of right". Mirror means: outer child of one matches outer child of the other, inner matches inner. So compare `t1.Left` with `t2.Right` and `t1.Right` with `t2.Left` (note the **cross**). This is the identical-trees idea with the recursion crossed over.

- **O(n) time, O(h) space.**
- **Edge cases:** empty tree and single node are symmetric.

### [Count tree nodes](count-tree-nodes/count_tree_nodes.go)

Total number of nodes.

**How to think about it:** the count of a tree = 1 (this node) + count(left) + count(right). Base case: `nil` contributes 0. This file ships with a full **call-stack dry run** in its comments — read it to internalize how recursion unwinds left-first, then right, then combines on the way back up.

- **O(n) time, O(h) space.**

### [Validate a BST](validate-binary-search-tree/validate_bst.go)

Is this a valid Binary Search Tree (every left descendant `<` node `<` every right descendant)?

**The trap most people fall into:** checking only `left.Val < node.Val < right.Val` locally. That's wrong — a node deep in the left subtree could still be larger than an ancestor. **A node must fall within a valid (min, max) range inherited from all its ancestors**, not just its parent.

**Approach:** recurse carrying `min`/`max` bounds (pointers so `nil` means "unbounded"). Going **left**, the current node becomes the new **upper** bound; going **right**, it becomes the new **lower** bound. Any violation → false.

**Dry run** on root=10, left=51, right=15: at 10 bounds are (−∞,+∞) ok; go left with bounds (−∞,10) → 51 ≥ 10 → **false**. (The local check `51 vs 10` alone might look tempting to pass; the range check correctly rejects it.)

- **O(n) time, O(h) space.**
- **Alternative:** an in-order traversal of a BST must be strictly increasing — you can validate by checking that instead.

## Review checklist

- Can you write the recursive tree template and adapt it to a new problem in under a minute?
- Which traversal gives sorted order on a BST, and which do you use to serialize vs free a tree?
- Why does BST validation need inherited (min, max) bounds instead of a local comparison?
- Why does symmetric-tree cross the recursion (`Left` vs `Right`) while identical-tree does not?
