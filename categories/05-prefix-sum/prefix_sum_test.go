package prefixsum

import "testing"

func TestNumArray(t *testing.T) {
	na := NewNumArray([]int{-2, 0, 3, -5, 2, -1})
	if got := na.SumRange(0, 2); got != 1 {
		t.Errorf("got %d", got)
	}
	if got := na.SumRange(2, 5); got != -1 {
		t.Errorf("got %d", got)
	}
}

func TestSubarraySumEqualsK(t *testing.T) {
	if got := SubarraySumEqualsK([]int{1, 1, 1}, 2); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := SubarraySumEqualsK([]int{1, 2, 3}, 3); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := SubarraySumEqualsK([]int{1, -1, 0}, 0); got != 3 {
		t.Errorf("got %d", got)
	}
}

func TestFindMaxLength(t *testing.T) {
	if got := FindMaxLength([]int{0, 1}); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := FindMaxLength([]int{0, 1, 0}); got != 2 {
		t.Errorf("got %d", got)
	}
}

func TestPivotIndex(t *testing.T) {
	if got := PivotIndex([]int{1, 7, 3, 6, 5, 6}); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := PivotIndex([]int{1, 2, 3}); got != -1 {
		t.Errorf("got %d", got)
	}
}

func TestSubarraysDivByK(t *testing.T) {
	if got := SubarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5); got != 7 {
		t.Errorf("got %d", got)
	}
}

func TestCheckSubarraySum(t *testing.T) {
	if !CheckSubarraySum([]int{23, 2, 4, 6, 7}, 6) {
		t.Error("expected true")
	}
	if CheckSubarraySum([]int{23, 2, 6, 4, 7}, 13) {
		t.Error("expected false")
	}
}

func TestMaxSubArrayLen(t *testing.T) {
	if got := MaxSubArrayLen([]int{1, -1, 5, -2, 3}, 3); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := MaxSubArrayLen([]int{-2, -1, 2, 1}, 1); got != 2 {
		t.Errorf("got %d", got)
	}
}

func TestNumMatrix(t *testing.T) {
	nm := NewNumMatrix([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	if got := nm.SumRegion(2, 1, 4, 3); got != 8 {
		t.Errorf("got %d", got)
	}
	if got := nm.SumRegion(1, 1, 2, 2); got != 11 {
		t.Errorf("got %d", got)
	}
}
