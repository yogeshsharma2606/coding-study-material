package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	got := ToSlice(ReverseList(FromSlice([]int{1, 2, 3, 4, 5})))
	if !reflect.DeepEqual(got, []int{5, 4, 3, 2, 1}) {
		t.Errorf("got %v", got)
	}
}

func TestHasCycleAndDetect(t *testing.T) {
	head := FromSlice([]int{3, 2, 0, -4})
	// create a cycle: tail -> node index 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = head.Next
	if !HasCycle(head) {
		t.Error("expected cycle")
	}
	if DetectCycleStart(head) != head.Next {
		t.Error("wrong cycle start")
	}

	noCycle := FromSlice([]int{1, 2})
	if HasCycle(noCycle) {
		t.Error("expected no cycle")
	}
}

func TestMergeTwoLists(t *testing.T) {
	got := ToSlice(MergeTwoLists(FromSlice([]int{1, 2, 4}), FromSlice([]int{1, 3, 4})))
	if !reflect.DeepEqual(got, []int{1, 1, 2, 3, 4, 4}) {
		t.Errorf("got %v", got)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	got := ToSlice(RemoveNthFromEnd(FromSlice([]int{1, 2, 3, 4, 5}), 2))
	if !reflect.DeepEqual(got, []int{1, 2, 3, 5}) {
		t.Errorf("got %v", got)
	}
	got = ToSlice(RemoveNthFromEnd(FromSlice([]int{1}), 1))
	if len(got) != 0 {
		t.Errorf("got %v", got)
	}
}

func TestMiddleNode(t *testing.T) {
	if MiddleNode(FromSlice([]int{1, 2, 3, 4, 5})).Val != 3 {
		t.Error("odd middle wrong")
	}
	if MiddleNode(FromSlice([]int{1, 2, 3, 4, 5, 6})).Val != 4 {
		t.Error("even middle wrong")
	}
}

func TestReorderList(t *testing.T) {
	head := FromSlice([]int{1, 2, 3, 4, 5})
	ReorderList(head)
	if got := ToSlice(head); !reflect.DeepEqual(got, []int{1, 5, 2, 4, 3}) {
		t.Errorf("got %v", got)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	got := ToSlice(AddTwoNumbers(FromSlice([]int{2, 4, 3}), FromSlice([]int{5, 6, 4})))
	if !reflect.DeepEqual(got, []int{7, 0, 8}) {
		t.Errorf("got %v", got)
	}
}

func TestIsPalindromeList(t *testing.T) {
	if !IsPalindromeList(FromSlice([]int{1, 2, 2, 1})) {
		t.Error("expected palindrome")
	}
	if IsPalindromeList(FromSlice([]int{1, 2})) {
		t.Error("expected not palindrome")
	}
}
