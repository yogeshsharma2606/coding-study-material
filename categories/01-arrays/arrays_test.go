package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		if got := TwoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TwoSum(%v,%d)=%v want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{}, 0},
	}
	for _, tt := range tests {
		if got := MaxProfit(tt.in); got != tt.want {
			t.Errorf("MaxProfit(%v)=%d want %d", tt.in, got, tt.want)
		}
	}
}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{-1, -2, -3}, -1},
	}
	for _, tt := range tests {
		if got := MaxSubArray(tt.in); got != tt.want {
			t.Errorf("MaxSubArray(%v)=%d want %d", tt.in, got, tt.want)
		}
	}
}

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
	}
	for _, tt := range tests {
		if got := MaxProduct(tt.in); got != tt.want {
			t.Errorf("MaxProduct(%v)=%d want %d", tt.in, got, tt.want)
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	in := []int{0, 1, 0, 3, 12}
	MoveZeroes(in)
	want := []int{1, 3, 12, 0, 0}
	if !reflect.DeepEqual(in, want) {
		t.Errorf("MoveZeroes=%v want %v", in, want)
	}
}

func TestProductExceptSelf(t *testing.T) {
	got := ProductExceptSelf([]int{1, 2, 3, 4})
	want := []int{24, 12, 8, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ProductExceptSelf=%v want %v", got, want)
	}
}

func TestRotateRight(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7}
	RotateRight(in, 3)
	want := []int{5, 6, 7, 1, 2, 3, 4}
	if !reflect.DeepEqual(in, want) {
		t.Errorf("RotateRight=%v want %v", in, want)
	}
}

func TestSortColors(t *testing.T) {
	in := []int{2, 0, 2, 1, 1, 0}
	SortColors(in)
	want := []int{0, 0, 1, 1, 2, 2}
	if !reflect.DeepEqual(in, want) {
		t.Errorf("SortColors=%v want %v", in, want)
	}
}

func TestMergeSorted(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	MergeSorted(nums1, 3, []int{2, 5, 6}, 3)
	want := []int{1, 2, 2, 3, 5, 6}
	if !reflect.DeepEqual(nums1, want) {
		t.Errorf("MergeSorted=%v want %v", nums1, want)
	}
}
