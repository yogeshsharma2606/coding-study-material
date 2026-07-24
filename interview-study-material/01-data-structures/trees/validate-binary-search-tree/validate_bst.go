package main

import "fmt"

type NodeT struct {
	Val   int
	Left  *NodeT
	Right *NodeT
}

func isValidBST(root *NodeT) bool {
	return validate(root, nil, nil)
}

func validate(node *NodeT, min, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}
	return validate(node.Left, min, &node.Val) &&
		validate(node.Right, &node.Val, max)
}

func main() {
	root := &NodeT{Val: 10}
	root.Left = &NodeT{Val: 51}
	root.Right = &NodeT{Val: 15}
	fmt.Println(isValidBST(root))
}
