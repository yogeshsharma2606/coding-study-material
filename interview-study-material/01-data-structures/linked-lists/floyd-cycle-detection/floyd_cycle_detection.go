package main

import "fmt"

func hasCycle(head *SNode) bool {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	head := &SNode{data: 1}
	head.next = &SNode{data: 2}
	head.next.next = &SNode{data: 3}
	head.next.next.next = &SNode{data: 4}
	head.next.next.next.next = head
	fmt.Println(hasCycle(head))
}
