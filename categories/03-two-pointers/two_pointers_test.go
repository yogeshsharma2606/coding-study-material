package twopointers

import (
	"reflect"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {
	if got := TwoSumSorted([]int{2, 7, 11, 15}, 9); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("got %v", got)
	}
	if got := TwoSumSorted([]int{2, 3, 4}, 6); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("got %v", got)
	}
}

func TestValidPalindromeII(t *testing.T) {
	if !ValidPalindromeII("aba") {
		t.Error("aba")
	}
	if !ValidPalindromeII("abca") {
		t.Error("abca should be true (delete c)")
	}
	if ValidPalindromeII("abc") {
		t.Error("abc should be false")
	}
}

func TestMaxArea(t *testing.T) {
	if got := MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}); got != 49 {
		t.Errorf("got %d", got)
	}
}

func TestThreeSum(t *testing.T) {
	got := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	want := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestThreeSumClosest(t *testing.T) {
	if got := ThreeSumClosest([]int{-1, 2, 1, -4}, 1); got != 2 {
		t.Errorf("got %d", got)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := RemoveDuplicates(nums)
	if !reflect.DeepEqual(nums[:n], []int{0, 1, 2, 3, 4}) {
		t.Errorf("got %v", nums[:n])
	}
}

func TestSortedSquares(t *testing.T) {
	got := SortedSquares([]int{-4, -1, 0, 3, 10})
	want := []int{0, 1, 9, 16, 100}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTrapRainWater(t *testing.T) {
	if got := TrapRainWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}); got != 6 {
		t.Errorf("got %d", got)
	}
}
