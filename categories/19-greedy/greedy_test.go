package greedy

import (
	"reflect"
	"testing"
)

func TestCanJump(t *testing.T) {
	if !CanJump([]int{2, 3, 1, 1, 4}) {
		t.Error("should reach")
	}
	if CanJump([]int{3, 2, 1, 0, 4}) {
		t.Error("should not reach")
	}
}

func TestJump(t *testing.T) {
	if got := Jump([]int{2, 3, 1, 1, 4}); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := Jump([]int{2, 3, 0, 1, 4}); got != 2 {
		t.Errorf("got %d", got)
	}
}

func TestCanCompleteCircuit(t *testing.T) {
	if got := CanCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := CanCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}); got != -1 {
		t.Errorf("got %d", got)
	}
}

func TestFindContentChildren(t *testing.T) {
	if got := FindContentChildren([]int{1, 2, 3}, []int{1, 1}); got != 1 {
		t.Errorf("got %d", got)
	}
	if got := FindContentChildren([]int{1, 2}, []int{1, 2, 3}); got != 2 {
		t.Errorf("got %d", got)
	}
}

func TestMaxProfitII(t *testing.T) {
	if got := MaxProfitII([]int{7, 1, 5, 3, 6, 4}); got != 7 {
		t.Errorf("got %d", got)
	}
	if got := MaxProfitII([]int{1, 2, 3, 4, 5}); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := MaxProfitII([]int{7, 6, 4, 3, 1}); got != 0 {
		t.Errorf("got %d", got)
	}
}

func TestPartitionLabels(t *testing.T) {
	got := PartitionLabels("ababcbacadefegdehijhklij")
	want := []int{9, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestLeastInterval(t *testing.T) {
	if got := LeastInterval([]byte("AAABBB"), 2); got != 8 {
		t.Errorf("got %d", got)
	}
	if got := LeastInterval([]byte("AAABBB"), 0); got != 6 {
		t.Errorf("got %d", got)
	}
	if got := LeastInterval([]byte("AAAA"), 2); got != 10 {
		t.Errorf("got %d", got)
	}
}

func TestCandy(t *testing.T) {
	if got := Candy([]int{1, 0, 2}); got != 5 {
		t.Errorf("got %d", got)
	}
	if got := Candy([]int{1, 2, 2}); got != 4 {
		t.Errorf("got %d", got)
	}
}
