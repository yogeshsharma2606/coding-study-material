package bitops

import (
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	if got := SingleNumber([]int{2, 2, 1}); got != 1 {
		t.Errorf("got %d", got)
	}
	if got := SingleNumber([]int{4, 1, 2, 1, 2}); got != 4 {
		t.Errorf("got %d", got)
	}
}

func TestHammingWeight(t *testing.T) {
	if got := HammingWeight(11); got != 3 { // 1011
		t.Errorf("got %d", got)
	}
	if got := HammingWeight(128); got != 1 {
		t.Errorf("got %d", got)
	}
}

func TestCountBits(t *testing.T) {
	got := CountBits(5)
	want := []int{0, 1, 1, 2, 1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestReverseBits(t *testing.T) {
	if got := ReverseBits(43261596); got != 964176192 {
		t.Errorf("got %d", got)
	}
}

func TestMissingNumber(t *testing.T) {
	if got := MissingNumber([]int{3, 0, 1}); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := MissingNumber([]int{0, 1}); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}); got != 8 {
		t.Errorf("got %d", got)
	}
}

func TestSingleNumberII(t *testing.T) {
	if got := SingleNumberII([]int{2, 2, 3, 2}); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := SingleNumberII([]int{0, 1, 0, 1, 0, 1, 99}); got != 99 {
		t.Errorf("got %d", got)
	}
}

func TestGetSum(t *testing.T) {
	if got := GetSum(1, 2); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := GetSum(-2, 3); got != 1 {
		t.Errorf("got %d", got)
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	if !IsPowerOfTwo(16) {
		t.Error("16 is power of two")
	}
	if IsPowerOfTwo(3) {
		t.Error("3 is not")
	}
	if IsPowerOfTwo(0) {
		t.Error("0 is not")
	}
}

func TestRangeBitwiseAnd(t *testing.T) {
	if got := RangeBitwiseAnd(5, 7); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := RangeBitwiseAnd(0, 0); got != 0 {
		t.Errorf("got %d", got)
	}
	if got := RangeBitwiseAnd(1, 2147483647); got != 0 {
		t.Errorf("got %d", got)
	}
}
