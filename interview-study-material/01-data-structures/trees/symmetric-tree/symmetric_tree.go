package main

func isMirror(t1, t2 *Node) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val &&
		isMirror(t1.Left, t2.Right) &&
		isMirror(t1.Right, t2.Left)
}

func isSymmetric(root *Node) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

/*
SYMMETRIC TREE EXAMPLES

Symmetric (isSymmetric → true): left subtree mirrors right

        1
       / \
      2   2
     / \ / \
    3  4 4  3

        1
       / \
      2   2
       \   \
        3   3

Single node is symmetric:
        1

Empty tree (nil root) is symmetric.


Non-symmetric (isSymmetric → false):

        1
       / \
      2   2
       \   \
        3   4     ← values differ (3 vs 4)

        1
       / \
      2   3       ← structure differs (2 has children, 3 is leaf)
     /
    4
*/
