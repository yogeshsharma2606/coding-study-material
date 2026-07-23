// Package arrays contains classic array-manipulation interview problems.
//
// Theme: arrays are contiguous memory, so index math is O(1). Most "optimal"
// array solutions come from (a) a single scan that maintains an invariant,
// (b) two indices moving toward/with each other, or (c) trading space for a
// hash map. Watch for the words "in-place" (O(1) extra space) and "sorted".
package arrays

// TwoSum returns indices of the two numbers that add up to target.
// Approach: one pass, hash map of value -> index. For each x we check whether
// target-x was seen already. Trades O(n) space for O(n) time (vs O(n^2) brute).
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, x := range nums {
		if j, ok := seen[target-x]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return nil
}

// MaxProfit is Best Time to Buy and Sell Stock (one transaction).
// Track the minimum price so far; the best profit ending today is
// price[today] - minSoFar. Single scan, O(1) space.
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	best := 0
	for _, p := range prices[1:] {
		if p-minPrice > best {
			best = p - minPrice
		}
		if p < minPrice {
			minPrice = p
		}
	}
	return best
}

// MaxSubArray returns the largest sum of any contiguous subarray (Kadane).
// Invariant: cur = best subarray sum ending at i. Either extend the previous
// subarray or start fresh at nums[i], whichever is larger.
func MaxSubArray(nums []int) int {
	cur, best := nums[0], nums[0]
	for _, x := range nums[1:] {
		if cur+x > x {
			cur = cur + x
		} else {
			cur = x
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

// MaxProduct returns the largest product of any contiguous subarray.
// A negative can flip the biggest product to the smallest and vice versa,
// so we track BOTH the running max and min ending here.
func MaxProduct(nums []int) int {
	curMax, curMin, best := nums[0], nums[0], nums[0]
	for _, x := range nums[1:] {
		if x < 0 {
			curMax, curMin = curMin, curMax
		}
		curMax = max(x, curMax*x)
		curMin = min(x, curMin*x)
		if curMax > best {
			best = curMax
		}
	}
	return best
}

// MoveZeroes moves all zeros to the end in-place, preserving order.
// A slow pointer marks the next slot for a non-zero; the fast pointer scans.
func MoveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

// ProductExceptSelf returns out[i] = product of all elements except nums[i],
// without division and in O(n). Two sweeps: prefix products, then suffix.
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	out[0] = 1
	for i := 1; i < n; i++ {
		out[i] = out[i-1] * nums[i-1] // product of everything to the left
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		out[i] *= right // multiply by product of everything to the right
		right *= nums[i]
	}
	return out
}

// RotateRight rotates nums to the right by k using the reverse trick, in-place.
// reverse(all); reverse(first k); reverse(rest).
func RotateRight(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(a []int, i, j int) {
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

// SortColors sorts an array of 0s, 1s, 2s in one pass (Dutch National Flag).
// low/high bound the unknown region; mid scans it.
func SortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high-- // don't advance mid: the swapped-in value is unexamined
		}
	}
}

// MergeSorted merges nums2 into nums1 in-place. nums1 has length m+n with the
// last n slots empty (zeros). Fill from the back to avoid overwriting.
func MergeSorted(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}
