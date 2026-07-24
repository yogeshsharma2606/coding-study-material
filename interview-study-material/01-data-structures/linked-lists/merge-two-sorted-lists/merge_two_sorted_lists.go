package main

import "fmt"

type SNode struct {
	data int
	next *SNode
}

func merge(l1, l2 *SNode) *SNode {
	dummy := &SNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.data < l2.data {
			curr.next = l1
			l1 = l1.next
		} else {
			curr.next = l2
			l2 = l2.next
		}
		curr = curr.next
	}

	if l1 != nil {
		curr.next = l1
	} else {
		curr.next = l2
	}
	return dummy.next
}

func display(head *SNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.data, " -> ")
		curr = curr.next
	}
	fmt.Println("nil")
}

func main() {
	l1 := &SNode{data: 1}
	l1.next = &SNode{data: 3}
	l1.next.next = &SNode{data: 5}
	l2 := &SNode{data: 2}
	l2.next = &SNode{data: 4}
	l2.next.next = &SNode{data: 6}
	display(merge(l1, l2))
}
