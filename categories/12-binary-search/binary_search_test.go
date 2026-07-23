package binarysearch

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	if got := Search([]int{-1, 0, 3, 5, 9, 12}, 9); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := Search([]int{-1, 0, 3, 5, 9, 12}, 2); got != -1 {
		t.Errorf("got %d", got)
	}
}

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		target, want int
	}{{5, 2}, {2, 1}, {7, 4}, {0, 0}}
	nums := []int{1, 3, 5, 6}
	for _, tt := range tests {
		if got := SearchInsert(nums, tt.target); got != tt.want {
			t.Errorf("SearchInsert(%d)=%d want %d", tt.target, got, tt.want)
		}
	}
}

func TestSearchRange(t *testing.T) {
	if got := SearchRange([]int{5, 7, 7, 8, 8, 10}, 8); !reflect.DeepEqual(got, []int{3, 4}) {
		t.Errorf("got %v", got)
	}
	if got := SearchRange([]int{5, 7, 7, 8, 8, 10}, 6); !reflect.DeepEqual(got, []int{-1, -1}) {
		t.Errorf("got %v", got)
	}
}

func TestSearchRotated(t *testing.T) {
	if got := SearchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 0); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := SearchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 3); got != -1 {
		t.Errorf("got %d", got)
	}
}

func TestFindMin(t *testing.T) {
	if got := FindMin([]int{3, 4, 5, 1, 2}); got != 1 {
		t.Errorf("got %d", got)
	}
	if got := FindMin([]int{4, 5, 6, 7, 0, 1, 2}); got != 0 {
		t.Errorf("got %d", got)
	}
	if got := FindMin([]int{11, 13, 15, 17}); got != 11 {
		t.Errorf("got %d", got)
	}
}

func TestFindPeakElement(t *testing.T) {
	got := FindPeakElement([]int{1, 2, 3, 1})
	if got != 2 {
		t.Errorf("got %d", got)
	}
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	got = FindPeakElement(nums)
	if !(got == 1 || got == 5) {
		t.Errorf("got %d (expected a peak index)", got)
	}
}

func TestMinEatingSpeed(t *testing.T) {
	if got := MinEatingSpeed([]int{3, 6, 7, 11}, 8); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := MinEatingSpeed([]int{30, 11, 23, 4, 20}, 5); got != 30 {
		t.Errorf("got %d", got)
	}
}

func TestShipWithinDays(t *testing.T) {
	if got := ShipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5); got != 15 {
		t.Errorf("got %d", got)
	}
	if got := ShipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3); got != 6 {
		t.Errorf("got %d", got)
	}
}

func TestMySqrt(t *testing.T) {
	if got := MySqrt(8); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := MySqrt(16); got != 4 {
		t.Errorf("got %d", got)
	}
}
