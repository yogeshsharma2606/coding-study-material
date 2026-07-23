// Package twopointers contains two-pointer interview problems.
//
// The two-pointer pattern replaces a nested loop (O(n^2)) with two indices that
// move in a single pass (O(n)). Two flavors: (1) CONVERGING pointers from both
// ends of a SORTED array, using the sortedness to decide which side to move;
// (2) SLOW/FAST pointers where slow marks a write position and fast scans.
package twopointers

import "sort"

// TwoSumSorted returns 1-based indices of two numbers summing to target,
// on a sorted array. Move left up if the sum is too small, right down if too
// big; sortedness guarantees we never miss the answer.
func TwoSumSorted(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		switch {
		case sum == target:
			return []int{l + 1, r + 1}
		case sum < target:
			l++
		default:
			r--
		}
	}
	return nil
}

// ValidPalindromeII returns true if s can be a palindrome after deleting at
// most one character. On the first mismatch, try skipping either side.
func ValidPalindromeII(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isPal(s, l+1, r) || isPal(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isPal(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// MaxArea is Container With Most Water. Area = width * min(height). Start wide;
// always move the SHORTER wall inward, because moving the taller one can only
// shrink the area (width drops, height still bounded by the shorter wall).
func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	best := 0
	for l < r {
		h := min(height[l], height[r])
		if area := h * (r - l); area > best {
			best = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return best
}

// ThreeSum returns all unique triplets summing to zero.
// Sort, then fix i and run a two-pointer scan on the remainder. Skip duplicates
// at every level to keep triplets unique.
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate anchor
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum == 0:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			case sum < 0:
				l++
			default:
				r--
			}
		}
	}
	return res
}

// ThreeSumClosest returns the sum of the triplet closest to target.
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	best := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(sum-target) < abs(best-target) {
				best = sum
			}
			if sum == target {
				return sum
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return best
}

// RemoveDuplicates removes duplicates from a sorted array in place and returns
// the new length. slow points to the last unique value written.
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// SortedSquares returns the sorted squares of a sorted (possibly negative)
// array. The largest square is at one of the two ends, so fill the output from
// the back using converging pointers - O(n) instead of O(n log n) re-sort.
func SortedSquares(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	l, r := 0, n-1
	for i := n - 1; i >= 0; i-- {
		ls, rs := nums[l]*nums[l], nums[r]*nums[r]
		if ls > rs {
			out[i] = ls
			l++
		} else {
			out[i] = rs
			r--
		}
	}
	return out
}

// TrapRainWater computes trapped water using two pointers and running maxima.
// Water above index i = min(maxLeft, maxRight) - height[i]. Move the side with
// the smaller wall, since that side's bound is the binding constraint.
func TrapRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l, r := 0, len(height)-1
	leftMax, rightMax, water := height[l], height[r], 0
	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, height[l])
			water += leftMax - height[l]
		} else {
			r--
			rightMax = max(rightMax, height[r])
			water += rightMax - height[r]
		}
	}
	return water
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
