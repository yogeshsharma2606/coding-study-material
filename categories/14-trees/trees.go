// Package trees contains binary-tree and BST interview problems.
//
// Trees are recursion made concrete: solve for the children, then combine. Ask
// "what do I need FROM my children, and what do I RETURN to my parent?" Traversal
// order matters: preorder (root first) for copying/serialization, inorder for
// BSTs (gives sorted order), postorder (children first) for bottom-up aggregates
// like height/diameter. BFS (a queue) gives level-order and shortest depth.
package trees

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// MaxDepth returns the height (max root-to-leaf node count). Postorder: depth is
// 1 + max(childDepths).
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}

// InorderTraversal returns node values in inorder (left, root, right) using an
// explicit stack (iterative to show the pattern; recursion is trivial).
func InorderTraversal(root *TreeNode) []int {
	var out []int
	var st []*TreeNode
	cur := root
	for cur != nil || len(st) > 0 {
		for cur != nil { // go as far left as possible
			st = append(st, cur)
			cur = cur.Left
		}
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		out = append(out, cur.Val)
		cur = cur.Right
	}
	return out
}

// LevelOrder returns values grouped by level using BFS with a queue.
func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// IsValidBST checks the BST property using min/max bounds passed down. Each node
// must lie strictly within (low, high) inherited from its ancestors.
func IsValidBST(root *TreeNode) bool {
	var check func(n *TreeNode, low, high *int) bool
	check = func(n *TreeNode, low, high *int) bool {
		if n == nil {
			return true
		}
		if low != nil && n.Val <= *low {
			return false
		}
		if high != nil && n.Val >= *high {
			return false
		}
		return check(n.Left, low, &n.Val) && check(n.Right, &n.Val, high)
	}
	return check(root, nil, nil)
}

// LowestCommonAncestor returns the LCA of p and q in a binary tree.
// Postorder: if a subtree contains p or q (or is one), return it up; the node
// where the two searches meet is the LCA.
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root // p and q found in different subtrees
	}
	if left != nil {
		return left
	}
	return right
}

// DiameterOfBinaryTree returns the longest path (in edges) between any two nodes.
// Postorder: at each node the through-path is leftHeight + rightHeight; track the
// global max while returning height upward.
func DiameterOfBinaryTree(root *TreeNode) int {
	best := 0
	var height func(n *TreeNode) int
	height = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := height(n.Left)
		r := height(n.Right)
		if l+r > best {
			best = l + r
		}
		return 1 + max(l, r)
	}
	height(root)
	return best
}

// InvertTree mirrors the tree (swap children recursively).
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}

// KthSmallest returns the kth smallest value in a BST. Inorder yields sorted
// order; stop after k nodes.
func KthSmallest(root *TreeNode, k int) int {
	var st []*TreeNode
	cur := root
	for cur != nil || len(st) > 0 {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
	return -1
}

// RightSideView returns the values visible from the right (last node per level).
func RightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == size-1 {
				res = append(res, node.Val) // rightmost of the level
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

// Codec serializes/deserializes a binary tree using preorder with nil markers.
type Codec struct{}

func (c *Codec) Serialize(root *TreeNode) string {
	var sb []byte
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			sb = append(sb, '#', ',')
			return
		}
		sb = append(sb, []byte(itoa(n.Val))...)
		sb = append(sb, ',')
		dfs(n.Left)
		dfs(n.Right)
	}
	dfs(root)
	return string(sb)
}

func (c *Codec) Deserialize(data string) *TreeNode {
	tokens := split(data)
	i := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		tok := tokens[i]
		i++
		if tok == "#" {
			return nil
		}
		n := &TreeNode{Val: atoi(tok)}
		n.Left = build()
		n.Right = build()
		return n
	}
	return build()
}

// --- small helpers (avoid strconv import noise) ---

func split(s string) []string {
	var out []string
	cur := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			out = append(out, cur)
			cur = ""
		} else {
			cur += string(s[i])
		}
	}
	return out
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	if neg {
		b = append([]byte{'-'}, b...)
	}
	return string(b)
}

func atoi(s string) int {
	n, sign, i := 0, 1, 0
	if len(s) > 0 && s[0] == '-' {
		sign = -1
		i = 1
	}
	for ; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return sign * n
}
