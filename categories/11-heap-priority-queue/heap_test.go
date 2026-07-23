package heapx

import (
	"reflect"
	"sort"
	"testing"
)

func fromSlice(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *ListNode) []int {
	var out []int
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func TestFindKthLargest(t *testing.T) {
	if got := FindKthLargest([]int{3, 2, 1, 5, 6, 4}, 2); got != 5 {
		t.Errorf("got %d", got)
	}
	if got := FindKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4); got != 4 {
		t.Errorf("got %d", got)
	}
}

func TestKthLargest(t *testing.T) {
	kl := NewKthLargest(3, []int{4, 5, 8, 2})
	want := []int{4, 5, 5, 8, 8}
	adds := []int{3, 5, 10, 9, 4}
	for i, a := range adds {
		if got := kl.Add(a); got != want[i] {
			t.Errorf("Add(%d)=%d want %d", a, got, want[i])
		}
	}
}

func TestLastStoneWeight(t *testing.T) {
	if got := LastStoneWeight([]int{2, 7, 4, 1, 8, 1}); got != 1 {
		t.Errorf("got %d", got)
	}
	if got := LastStoneWeight([]int{1, 1}); got != 0 {
		t.Errorf("got %d", got)
	}
}

func TestKClosest(t *testing.T) {
	got := KClosest([][]int{{1, 3}, {-2, 2}, {5, 8}, {0, 1}}, 2)
	sort.Slice(got, func(i, j int) bool {
		return got[i][0]*got[i][0]+got[i][1]*got[i][1] < got[j][0]*got[j][0]+got[j][1]*got[j][1]
	})
	want := [][]int{{0, 1}, {-2, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{
		fromSlice([]int{1, 4, 5}),
		fromSlice([]int{1, 3, 4}),
		fromSlice([]int{2, 6}),
	}
	got := toSlice(MergeKLists(lists))
	want := []int{1, 1, 2, 3, 4, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMedianFinder(t *testing.T) {
	m := NewMedianFinder()
	m.AddNum(1)
	m.AddNum(2)
	if got := m.FindMedian(); got != 1.5 {
		t.Errorf("got %v", got)
	}
	m.AddNum(3)
	if got := m.FindMedian(); got != 2.0 {
		t.Errorf("got %v", got)
	}
}

func TestTopKFrequent(t *testing.T) {
	got := TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	if !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("got %v", got)
	}
}

func TestReorganizeString(t *testing.T) {
	got := ReorganizeString("aab")
	if got == "" || got[0] == got[1] {
		t.Errorf("invalid result %q", got)
	}
	if ReorganizeString("aaab") != "" {
		t.Error("aaab should be impossible")
	}
}
