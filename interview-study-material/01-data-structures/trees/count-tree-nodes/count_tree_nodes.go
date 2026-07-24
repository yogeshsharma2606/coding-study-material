package main

import "fmt"

type NodeT struct {
	Val   int
	Left  *NodeT
	Right *NodeT
}

func countNodes(root *NodeT) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

/*
DRY RUN — Stack execution (tree: 1→2,3; 2→4,5)

Go evaluates left-to-right: first countNodes(Left), then countNodes(Right), then 1 + L + R.

  STEP  │ CALL STACK (top → bottom)           │ ACTION / RETURN
  ──────┼────────────────────────────────────┼────────────────────────────────────────
    1   │ countNodes(1)                      │ root≠nil → need 1 + countNodes(2) + countNodes(3)
    2   │ countNodes(2)  ← countNodes(1)     │ call Left first → countNodes(2)
    3   │ countNodes(4)  ← countNodes(2) ← (1)│ call Left first → countNodes(4)
    4   │ countNodes(nil)← countNodes(4) ←..  │ root==nil → return 0
    5   │ countNodes(4)  ← countNodes(2) ← (1)│ Left=0, now Right → countNodes(nil)→0
    6   │ countNodes(4)  ← countNodes(2) ← (1)│ 1+0+0 = 1 → return 1
    7   │ countNodes(5)  ← countNodes(2) ← (1)│ countNodes(2): Left=1, now Right → countNodes(5)
    8   │ countNodes(nil)← countNodes(5) ←..  │ return 0 (Left)
    9   │ countNodes(nil)← countNodes(5) ←..  │ return 0 (Right)
   10   │ countNodes(5)  ← countNodes(2) ← (1)│ 1+0+0 = 1 → return 1
   11   │ countNodes(2)  ← countNodes(1)      │ 1+1+1 = 3 → return 3
   12   │ countNodes(3)  ← countNodes(1)      │ countNodes(1): Left=3, now Right → countNodes(3)
   13   │ countNodes(nil)← countNodes(3) ← (1)│ return 0 (Left)
   14   │ countNodes(nil)← countNodes(3) ← (1)│ return 0 (Right)
   15   │ countNodes(3)  ← countNodes(1)      │ 1+0+0 = 1 → return 1
   16   │ countNodes(1)                      │ 1+3+1 = 5 → return 5

Result: 5 nodes.
*/

func main() {
	// Example tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	root := &NodeT{Val: 1}
	root.Left = &NodeT{Val: 2}
	root.Right = &NodeT{Val: 3}
	root.Left.Left = &NodeT{Val: 4}
	root.Left.Right = &NodeT{Val: 5}

	fmt.Println("Total nodes:", countNodes(root)) // Output: 5
}
