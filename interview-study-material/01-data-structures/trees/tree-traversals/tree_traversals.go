package main

import "fmt"

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

//DFS In-order
func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}

//DFS Pre-order
func preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preorder(root.Left)
	preorder(root.Right)
}

//DFS Post-order
func postorder(root *Node) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Print(root.Val, " ")
}

//BFS Level-order
func levelOrder(root *Node) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.Val, " ")
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	root := &Node{Val: "A"}
	root.Left = &Node{Val: "B"}
	root.Right = &Node{Val: "C"}
	root.Left.Left = &Node{Val: "D"}
	root.Left.Right = &Node{Val: "E"}

	fmt.Print("In-order: ")
	inorder(root)
	fmt.Println()
	fmt.Print("Pre-order: ")
	preorder(root)
	fmt.Println()
	fmt.Print("Post-order: ")
	postorder(root)
	fmt.Println()
	fmt.Print("Level-order: ")
	levelOrder(root)
	fmt.Println()
}

/*
TRAVERSAL EXAMPLES (tree from mainaa)

Tree structure:
            A
           / \
          B   C
         / \
        D   E

DFS In-order   (Left → Root → Right):  D B E A C
DFS Pre-order  (Root → Left → Right):  A B D E C
DFS Post-order (Left → Right → Root):  D E B C A
BFS Level-order (level by level):      A B C D E

Another example (larger tree):

              F
             / \
            G   H
           / \   \
          I   J   K
         /
        L

In-order:   L I G J F H K
Pre-order:  F G I L J H K
Post-order: L I J G K H F
Level-order: F G H I J K L
*/
