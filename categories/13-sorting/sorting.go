// Package sorting contains sorting algorithms and sort-based interview problems.
//
// Comparison sorts are bounded by O(n log n). Merge sort is stable and
// predictable (O(n) extra space); quicksort is in-place and fast in practice but
// O(n^2) worst case. When keys are small integers, COUNTING/RADIX sort beats the
// bound at O(n). Many "hard" problems are easy once you sort and/or supply the
// right comparator - recognizing that is the skill.
package sorting

import (
	"sort"
	"strconv"
)

// MergeSort returns a sorted copy using divide-and-conquer. Stable, O(n log n),
// O(n) extra space. Split in half, sort each, merge.
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return append([]int(nil), nums...)
	}
	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)
}

func merge(a, b []int) []int {
	out := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	out = append(out, a[i:]...)
	out = append(out, b[j:]...)
	return out
}

// QuickSort sorts in place using Lomuto partition. Average O(n log n), worst
// O(n^2) (mitigated by a mid-element pivot here).
func QuickSort(nums []int) {
	quicksort(nums, 0, len(nums)-1)
}

func quicksort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quicksort(a, lo, p-1)
	quicksort(a, p+1, hi)
}

func partition(a []int, lo, hi int) int {
	mid := lo + (hi-lo)/2
	a[mid], a[hi] = a[hi], a[mid] // use middle element as pivot, park at end
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

// QuickSelect returns the kth smallest element (1-indexed) in O(n) average by
// partitioning only the side containing the answer.
func QuickSelect(nums []int, k int) int {
	a := append([]int(nil), nums...)
	lo, hi, target := 0, len(a)-1, k-1
	for lo <= hi {
		p := partition(a, lo, hi)
		if p == target {
			return a[p]
		} else if p < target {
			lo = p + 1
		} else {
			hi = p - 1
		}
	}
	return -1
}

// CountingSort sorts non-negative ints in O(n + maxVal) using a frequency table
// - beats the comparison bound when the value range is small.
func CountingSort(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	maxV := nums[0]
	for _, x := range nums {
		if x > maxV {
			maxV = x
		}
	}
	count := make([]int, maxV+1)
	for _, x := range nums {
		count[x]++
	}
	out := make([]int, 0, len(nums))
	for v, c := range count {
		for ; c > 0; c-- {
			out = append(out, v)
		}
	}
	return out
}

// LargestNumber arranges numbers to form the largest concatenation.
// Custom comparator: a before b iff a+b > b+a as strings.
func largestNumber(nums []int) string {
	// Convert numbers to strings
	strs := make([]string, len(nums))

	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	// Custom comparator
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	// Edge case:
	// [0, 0, 0] -> "0", not "000"
	if strs[0] == "0" {
		return "0"
	}

	return strings.Join(strs, "")
}

// HIndex returns the researcher's h-index. Sort descending; the h-index is the
// largest i where citations[i-1] >= i.
func HIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	h := 0
	for i, c := range citations {
		if c >= i+1 {
			h = i + 1
		} else {
			break
		}
	}
	return h
}

// FrequencySort sorts characters by descending frequency (ties arbitrary).
func FrequencySort(s string) string {
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	chars := make([]rune, 0, len(freq))
	for r := range freq {
		chars = append(chars, r)
	}
	sort.Slice(chars, func(i, j int) bool { return freq[chars[i]] > freq[chars[j]] })
	out := make([]rune, 0, len(s))
	for _, r := range chars {
		for c := 0; c < freq[r]; c++ {
			out = append(out, r)
		}
	}
	return string(out)
}

// WiggleSort rearranges so nums[0] <= nums[1] >= nums[2] <= nums[3]...
// One pass: whenever the local order is wrong, swap - each fix can't break the
// previous relation.
func WiggleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if (i%2 == 0) == (nums[i] > nums[i+1]) {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
}
