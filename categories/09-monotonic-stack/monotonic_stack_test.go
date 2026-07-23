package monotonicstack

import (
	"reflect"
	"testing"
)

func TestNextGreaterElementI(t *testing.T) {
	got := NextGreaterElementI([]int{4, 1, 2}, []int{1, 3, 4, 2})
	if !reflect.DeepEqual(got, []int{-1, 3, -1}) {
		t.Errorf("got %v", got)
	}
}

func TestNextGreaterElementsII(t *testing.T) {
	got := NextGreaterElementsII([]int{1, 2, 1})
	if !reflect.DeepEqual(got, []int{2, -1, 2}) {
		t.Errorf("got %v", got)
	}
}

func TestDailyTemperatures(t *testing.T) {
	got := DailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	if !reflect.DeepEqual(got, []int{1, 1, 4, 2, 1, 1, 0, 0}) {
		t.Errorf("got %v", got)
	}
}

func TestLargestRectangleArea(t *testing.T) {
	if got := LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}); got != 10 {
		t.Errorf("got %d", got)
	}
	if got := LargestRectangleArea([]int{2, 4}); got != 4 {
		t.Errorf("got %d", got)
	}
}

func TestTrapStack(t *testing.T) {
	if got := TrapStack([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}); got != 6 {
		t.Errorf("got %d", got)
	}
}

func TestSumSubarrayMins(t *testing.T) {
	if got := SumSubarrayMins([]int{3, 1, 2, 4}); got != 17 {
		t.Errorf("got %d", got)
	}
	if got := SumSubarrayMins([]int{11, 81, 94, 43, 3}); got != 444 {
		t.Errorf("got %d", got)
	}
}

func TestRemoveKdigits(t *testing.T) {
	tests := []struct {
		num  string
		k    int
		want string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
	}
	for _, tt := range tests {
		if got := RemoveKdigits(tt.num, tt.k); got != tt.want {
			t.Errorf("RemoveKdigits(%q,%d)=%q want %q", tt.num, tt.k, got, tt.want)
		}
	}
}

func TestStockSpanner(t *testing.T) {
	s := NewStockSpanner()
	prices := []int{100, 80, 60, 70, 60, 75, 85}
	want := []int{1, 1, 1, 2, 1, 4, 6}
	for i, p := range prices {
		if got := s.Next(p); got != want[i] {
			t.Errorf("Next(%d)=%d want %d", p, got, want[i])
		}
	}
}
