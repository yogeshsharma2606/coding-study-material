// Package binarysearch contains binary-search interview problems.
//
// Binary search halves the search space each step: O(log n). It applies whenever
// the space is MONOTONIC - the predicate "is x a valid/answer position?" goes
// false...false...true...true (or the array is sorted). The senior-level move is
// "binary search on the ANSWER": when you can cheaply test feasibility for a
// candidate value, binary-search the value range instead of the array.
package binarysearch

// Search returns the index of target in a sorted array, or -1.
// Use the half-open invariant [lo, hi) and lo+(hi-lo)/2 to avoid overflow.
func Search(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return -1
}

// SearchInsert returns the index where target is, or where it would be inserted
// to keep the array sorted (the first index with nums[i] >= target).
func SearchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// SearchRange returns the first and last index of target, or [-1,-1].
// Two boundary searches: leftmost >= target and leftmost > target.
func SearchRange(nums []int, target int) []int {
	left := lowerBound(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := lowerBound(nums, target+1) - 1
	return []int{left, right}
}

// lowerBound returns the first index with nums[i] >= target.
func lowerBound(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// SearchRotated searches a rotated sorted array (distinct values) in O(log n).
// One half of [lo, hi] is always sorted; decide which, then whether target lies
// in that sorted half.
func SearchRotated(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] { // left half sorted
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // right half sorted
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

// FindMin returns the minimum of a rotated sorted array. Compare mid to the
// right end: if nums[mid] > nums[hi], the min is to the right; else it's mid or
// left.
func FindMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}

// FindPeakElement returns the index of any peak (nums[i] > neighbors), assuming
// nums[-1]=nums[n]=-inf. Move toward the higher neighbor - a peak must exist there.
func FindPeakElement(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < nums[mid+1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// MinEatingSpeed (Koko): min bananas/hour to finish all piles within h hours.
// Binary search on the ANSWER (speed). Feasibility is monotonic: faster speed is
// always still feasible if a slower one was.
func MinEatingSpeed(piles []int, h int) int {
	hoursNeeded := func(speed int) int {
		total := 0
		for _, p := range piles {
			total += (p + speed - 1) / speed // ceil division
		}
		return total
	}
	lo, hi := 1, 0
	for _, p := range piles {
		if p > hi {
			hi = p
		}
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		if hoursNeeded(mid) <= h {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// ShipWithinDays: least ship capacity to ship all weights (in order) within
// 'days'. Binary search on capacity; feasibility = number of days needed <= days.
func ShipWithinDays(weights []int, days int) int {
	daysNeeded := func(cap int) int {
		d, cur := 1, 0
		for _, w := range weights {
			if cur+w > cap {
				d++
				cur = 0
			}
			cur += w
		}
		return d
	}
	lo, hi := 0, 0
	for _, w := range weights {
		if w > lo {
			lo = w // capacity must fit the heaviest item
		}
		hi += w // one day for everything
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		if daysNeeded(mid) <= days {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// MySqrt returns floor(sqrt(x)) via binary search on the answer.
func MySqrt(x int) int {
	if x < 2 {
		return x
	}
	lo, hi := 1, x/2
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi // floor
}
