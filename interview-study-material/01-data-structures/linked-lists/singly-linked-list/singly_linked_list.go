package main

import "fmt"

// Node represents a single element in the list
type SNode struct {
	data int
	next *SNode
}

// LinkedList represents the list itself
type LinkedList struct {
	head *SNode
}

// Insert adds a new node at the end of the list
func (l *LinkedList) Insert(val int) {
	newNode := &SNode{data: val}
	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

// Display prints all elements of the list
func (l *LinkedList) Display() {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.data, " -> ")
		curr = curr.next
	}
	fmt.Println("nil")
}

// Search checks if a value exists in the list
func (l *LinkedList) Search(val int) bool {
	curr := l.head
	for curr != nil {
		if curr.data == val {
			return true
		}
		curr = curr.next
	}
	return false
}

// Delete removes the first occurrence of a value from the list
func (l *LinkedList) Delete(val int) {
	if l.head == nil {
		return
	}
	if l.head.data == val {
		l.head = l.head.next
		return
	}
	curr := l.head
	for curr.next != nil {
		if curr.next.data == val {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

// Reverse reverses the linked list
func (l *LinkedList) Reverse() {
	var prev *SNode
	curr := l.head
	for curr != nil {
		next := curr.next // save next
		curr.next = prev  // reverse link
		prev = curr       // move prev forward
		curr = next       // move curr forward
	}
	l.head = prev
}

func main() {
	list := LinkedList{}

	// Insert elements
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)

	fmt.Print("Linked List: ")
	list.Display() // 10 -> 20 -> 30 -> 40 -> nil

	// Search elements
	fmt.Println("Search 20:", list.Search(20)) // true
	fmt.Println("Search 50:", list.Search(50)) // false

	// Delete an element
	list.Delete(30)
	fmt.Print("After deleting 30: ")
	list.Display() // 10 -> 20 -> 40 -> nil

	// Reverse the list
	list.Reverse()
	fmt.Print("After reversing: ")
	list.Display() // 40 -> 20 -> 10 -> nil

}
