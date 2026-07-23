package sorting

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	got := MergeSort([]int{5, 2, 9, 1, 5, 6})
	if !reflect.DeepEqual(got, []int{1, 2, 5, 5, 6, 9}) {
		t.Errorf("got %v", got)
	}
}

func TestQuickSort(t *testing.T) {
	a := []int{5, 2, 9, 1, 5, 6, 3}
	QuickSort(a)
	if !sort.IntsAreSorted(a) {
		t.Errorf("not sorted: %v", a)
	}
}

func TestQuickSelect(t *testing.T) {
	if got := QuickSelect([]int{3, 2, 1, 5, 6, 4}, 2); got != 2 {
		t.Errorf("got %d", got) // 2nd smallest
	}
	if got := QuickSelect([]int{7, 10, 4, 3, 20, 15}, 3); got != 7 {
		t.Errorf("got %d", got)
	}
}

func TestCountingSort(t *testing.T) {
	got := CountingSort([]int{4, 2, 2, 8, 3, 3, 1})
	if !reflect.DeepEqual(got, []int{1, 2, 2, 3, 3, 4, 8}) {
		t.Errorf("got %v", got)
	}
}

func TestLargestNumber(t *testing.T) {
	if got := LargestNumber([]int{10, 2}); got != "210" {
		t.Errorf("got %q", got)
	}
	if got := LargestNumber([]int{3, 30, 34, 5, 9}); got != "9534330" {
		t.Errorf("got %q", got)
	}
	if got := LargestNumber([]int{0, 0}); got != "0" {
		t.Errorf("got %q", got)
	}
}

func TestHIndex(t *testing.T) {
	if got := HIndex([]int{3, 0, 6, 1, 5}); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := HIndex([]int{1, 3, 1}); got != 1 {
		t.Errorf("got %d", got)
	}
}

func TestFrequencySort(t *testing.T) {
	got := FrequencySort("tree")
	// validate: frequencies are non-increasing and it's a permutation
	freq := map[rune]int{}
	for _, r := range got {
		freq[r]++
	}
	if len(got) != 4 || freq['e'] != 2 || freq['t'] != 1 || freq['r'] != 1 {
		t.Errorf("bad multiset: %q", got)
	}
	if got[0] != 'e' || got[1] != 'e' {
		t.Errorf("most frequent should lead: %q", got)
	}
}

func TestWiggleSort(t *testing.T) {
	nums := []int{3, 5, 2, 1, 6, 4}
	WiggleSort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 && nums[i] > nums[i+1] {
			t.Errorf("wiggle broken at %d: %v", i, nums)
		}
		if i%2 == 1 && nums[i] < nums[i+1] {
			t.Errorf("wiggle broken at %d: %v", i, nums)
		}
	}
}
