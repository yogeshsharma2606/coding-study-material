package main

import "fmt"

type Node1 struct {
	Val   int
	Left  *Node1
	Right *Node1
}

// height or max depth of the tree
func height(root *Node1) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func heightBFS(root *Node1) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*Node1{root}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

func main() {
	// Example tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	root := &Node1{Val: 1}
	root.Left = &Node1{Val: 2}
	root.Right = &Node1{Val: 3}
	root.Left.Left = &Node1{Val: 4}
	root.Left.Right = &Node1{Val: 5}

	fmt.Println("Height of tree:", height(root))         // Output: 3
	fmt.Println("Level-order of tree:", heightBFS(root)) // Output: 3

}
