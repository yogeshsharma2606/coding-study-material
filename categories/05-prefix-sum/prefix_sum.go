// Package prefixsum contains prefix-sum interview problems.
//
// A prefix sum P[i] = a[0]+...+a[i-1] lets you answer "sum of range [l, r]" in
// O(1) as P[r+1]-P[l], after O(n) preprocessing. The deeper trick: the number of
// subarrays with a given property (sum == k, sum divisible by k, equal 0s/1s)
// equals the number of matching prefix pairs, which a HASH MAP of prefix values
// counts in one pass. Unlike sliding window, prefix sums handle NEGATIVE numbers.
package prefixsum

// NumArray answers immutable range-sum queries in O(1) after O(n) setup.
type NumArray struct {
	prefix []int // prefix[i] = sum of nums[0:i]
}

func NewNumArray(nums []int) *NumArray {
	p := make([]int, len(nums)+1)
	for i, x := range nums {
		p[i+1] = p[i] + x
	}
	return &NumArray{prefix: p}
}

// SumRange returns nums[l] + ... + nums[r].
func (na *NumArray) SumRange(l, r int) int {
	return na.prefix[r+1] - na.prefix[l]
}

// SubarraySumEqualsK counts contiguous subarrays summing to k.
// A subarray (l, r] has sum k iff P[r] - P[l] = k, i.e. P[l] = P[r] - k.
// Walk prefixes, counting how many earlier prefixes equal (running - k).
func SubarraySumEqualsK(nums []int, k int) int {
	count := map[int]int{0: 1} // empty prefix seen once
	running, res := 0, 0
	for _, x := range nums {
		running += x
		res += count[running-k]
		count[running]++
	}
	return res
}

// FindMaxLength returns the longest contiguous subarray with equal 0s and 1s.
// Map 0 -> -1 so "equal count" becomes "prefix sum returns to a prior value".
// Store the first index of each prefix value; length = i - firstIndex[prefix].
func FindMaxLength(nums []int) int {
	first := map[int]int{
		0: -1,
	}
	sum := 0
	maxLen := 0
	for i, num := range nums {
		if num == 0 {
			sum--
		} else {
			sum++
		}
		if start, ok := first[sum]; ok {
			length := i - start
			if length > maxLen {
				maxLen = length
			}
		} else {
			// Store only the first occurrence.
			first[sum] = i
		}
	}
	return maxLen
}

// PivotIndex returns the leftmost index where left sum == right sum, else -1.
// left + nums[i] + right = total; pivot when left == total - left - nums[i].
func PivotIndex(nums []int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	left := 0
	for i, x := range nums {
		if left == total-left-x {
			return i
		}
		left += x
	}
	return -1
}

// SubarraysDivByK counts subarrays whose sum is divisible by k.
// Two prefixes with the same remainder mod k bound a divisible subarray.
// Normalize the remainder to [0, k) to handle negatives in Go.
func SubarraysDivByK(nums []int, k int) int {
	count := map[int]int{0: 1}
	running, res := 0, 0
	for _, x := range nums {
		running += x
		r := ((running % k) + k) % k
		res += count[r]
		count[r]++
	}
	return res
}

// CheckSubarraySum reports whether there is a subarray of length >= 2 whose sum
// is a multiple of k. Same remainder at two indices >= 2 apart => divisible.
// Store the earliest index for each remainder.
func CheckSubarraySum(nums []int, k int) bool {
	first := map[int]int{0: -1}
	running := 0
	for i, x := range nums {
		running += x
		r := running % k
		if r < 0 {
			r += k
		}
		if j, ok := first[r]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			first[r] = i
		}
	}
	return false
}

// MaxSubArrayLen returns the longest subarray summing to exactly k (handles
// negatives). Store the earliest index of each prefix; if running-k was seen,
// candidate length = i - firstIndex[running-k].
func MaxSubArrayLen(nums []int, k int) int {
	first := map[int]int{0: -1}
	running, best := 0, 0
	for i, x := range nums {
		running += x
		if j, ok := first[running-k]; ok {
			if i-j > best {
				best = i - j
			}
		}
		if _, ok := first[running]; !ok {
			first[running] = i
		}
	}
	return best
}

// NumMatrix answers immutable 2D region-sum queries in O(1) using a 2D prefix
// sum where P[i][j] = sum of the rectangle from (0,0) to (i-1,j-1).
func (nm *NumMatrix) SumRegion(
	row1, col1, row2, col2 int,
) int {

	sum := 0

	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += nm.matrix[i][j]
		}
	}

	return sum
}

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		matrix: matrix,
	}
}
