package main

import "fmt"

type NodeD struct {
	data int
	prev *NodeD
	next *NodeD
}

type DoublyLinkedList struct {
	head *NodeD
}

func (dll *DoublyLinkedList) Insert(val int) {
	newNode := &NodeD{data: val}
	if dll.head == nil {
		dll.head = newNode
		return
	}
	curr := dll.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	newNode.prev = curr
}

func (dll *DoublyLinkedList) DisplayForward() {
	curr := dll.head
	for curr != nil {
		fmt.Print(curr.data, " <-> ")
		curr = curr.next
	}
	fmt.Println("nil")
}

func (dll *DoublyLinkedList) DisplayBackward() {
	curr := dll.head
	if curr == nil {
		return
	}
	for curr.next != nil {
		curr = curr.next
	}
	for curr != nil {
		fmt.Print(curr.data, " <-> ")
		curr = curr.prev
	}
	fmt.Println("nil")
}

func (dll *DoublyLinkedList) Search(val int) bool {
	curr := dll.head
	for curr != nil {
		if curr.data == val {
			return true
		}
		curr = curr.next
	}
	return false
}

func (dll *DoublyLinkedList) Delete(val int) {
	curr := dll.head
	for curr != nil {
		if curr.data == val {
			if curr.prev != nil {
				curr.prev.next = curr.next
			} else {
				dll.head = curr.next
			}
			if curr.next != nil {
				curr.next.prev = curr.prev
			}
			return
		}
		curr = curr.next
	}
}

func (dll *DoublyLinkedList) Reverse() {
	var temp *NodeD
	curr := dll.head
	for curr != nil {
		temp = curr.prev
		curr.prev = curr.next
		curr.next = temp
		curr = curr.prev
	}
	if temp != nil {
		dll.head = temp.prev
	}
}

func main() {
	dll := DoublyLinkedList{}
	dll.Insert(10)
	dll.Insert(20)
	dll.Insert(30)
	dll.Insert(40)

	fmt.Print("Forward: ")
	dll.DisplayForward() // 10 <-> 20 <-> 30 <-> 40 <-> nil

	fmt.Print("Backward: ")
	dll.DisplayBackward() // 40 <-> 30 <-> 20 <-> 10 <-> nil

	fmt.Println("Search 20:", dll.Search(20)) // true
	fmt.Println("Search 50:", dll.Search(50)) // false

	dll.Delete(10)
	fmt.Print("After deleting 10: ")
	dll.DisplayForward() // 10 <-> 20 <-> 40 <-> nil

	dll.Reverse()
	fmt.Print("After reverse: ")
	dll.DisplayForward() // 40 <-> 20 <-> 10 <-> nil
}
