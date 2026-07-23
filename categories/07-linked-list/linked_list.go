// Package linkedlist contains singly-linked-list interview problems.
//
// Linked lists trade O(1) insert/delete (given a node) for O(n) access - there's
// no indexing. The recurring techniques are: a DUMMY head node to simplify edge
// cases at the front, POINTER REVERSAL (rewire next pointers), and the FAST/SLOW
// (tortoise-hare) pattern for middle/cycle/nth-from-end in one pass.
package linkedlist

// ListNode is a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// FromSlice builds a list from a slice (test helper).
func FromSlice(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// ToSlice flattens a list to a slice (test helper).
func ToSlice(head *ListNode) []int {
	var out []int
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

// ReverseList reverses the list iteratively by flipping each next pointer.
// prev trails, cur leads; save next before rewiring.
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// HasCycle detects a cycle with Floyd's tortoise and hare. If a fast pointer
// (2x speed) ever meets the slow pointer, there is a loop.
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// DetectCycleStart returns the node where the cycle begins, or nil.
// After the meeting point, moving one pointer to head and advancing both at
// 1x meets at the cycle entrance (Floyd's second phase).
func DetectCycleStart(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// MergeTwoLists merges two sorted lists into one sorted list.
// A dummy head avoids special-casing the first node.
func MergeTwoLists(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}

// RemoveNthFromEnd removes the nth node from the end in one pass using a gap of
// n between two pointers. The dummy handles removing the head.
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// MiddleNode returns the middle node (second middle if even length).
func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// ReorderList reorders L0->L1->...->Ln to L0->Ln->L1->Ln-1->...
// Steps: find middle, reverse second half, merge the two halves alternately.
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// 1. find middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 2. reverse second half
	second := ReverseList(slow.Next)
	slow.Next = nil
	// 3. merge alternately
	first := head
	for second != nil {
		n1, n2 := first.Next, second.Next
		first.Next = second
		second.Next = n1
		first = n1
		second = n2
	}
}

// AddTwoNumbers adds two numbers stored as reversed digit lists.
func AddTwoNumbers(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0
	for a != nil || b != nil || carry != 0 {
		sum := carry
		if a != nil {
			sum += a.Val
			a = a.Next
		}
		if b != nil {
			sum += b.Val
			b = b.Next
		}
		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
	}
	return dummy.Next
}

// IsPalindromeList reports whether the list reads the same forwards/backwards.
// Find middle, reverse the second half, compare. O(1) extra space.
func IsPalindromeList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := ReverseList(slow)
	p1, p2 := head, second
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}
