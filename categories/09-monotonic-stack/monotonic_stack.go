// Package monotonicstack contains monotonic-stack interview problems.
//
// A monotonic stack keeps its elements sorted (increasing or decreasing). Before
// pushing x, pop everything that violates the order; those popped elements have
// just found x as their "next greater/smaller". Each element is pushed and popped
// at most once, so a whole family of "next/previous greater/smaller" problems
// that look O(n^2) become O(n). The stack usually stores INDICES, not values.
package monotonicstack

// NextGreaterElementI: for each x in nums1 (a subset of nums2), find the first
// greater element to its right in nums2. Build a value->nextGreater map with a
// decreasing monotonic stack over nums2, then look up.
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0)
	nextGreater := make(map[int]int)
	// Find next greater element for every element in nums2.
	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreater[top] = num
		}
		stack = append(stack, num)
	}
	// Elements remaining in the stack have no greater element.
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nextGreater[top] = -1
	}
	// Build result for nums1.
	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = nextGreater[num]
	}
	return result
}

// NextGreaterElementsII: circular array; find each element's next greater to the
// right, wrapping around. Iterate 2n times (mod n) over a decreasing stack.
func NextGreaterElementsII(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	var st []int // indices, values decreasing
	for i := 0; i < 2*n; i++ {
		x := nums[i%n]
		for len(st) > 0 && nums[st[len(st)-1]] < x {
			res[st[len(st)-1]] = x
			st = st[:len(st)-1]
		}
		if i < n {
			st = append(st, i)
		}
	}
	return res
}

// DailyTemperatures returns, for each day, how many days until a warmer one.
// Decreasing stack of indices; when today is warmer, pop and record the gap.
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{} // stack stores indices

	for i := 0; i < n; i++ {
		// Compare current temp with stack top
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}

	return result
}

// LargestRectangleArea finds the largest rectangle in a histogram.
// Maintain an increasing stack of indices; when a shorter bar arrives, pop and
// compute the area with the popped bar as the limiting height. A sentinel 0 at
// the end flushes the stack.
func LargestRectangleArea(heights []int) int {
	best := 0
	st := []int{} // indices, heights increasing
	hs := append([]int{}, heights...)
	hs = append(hs, 0) // sentinel to flush
	for i, h := range hs {
		for len(st) > 0 && hs[st[len(st)-1]] > h {
			top := st[len(st)-1]
			st = st[:len(st)-1]
			height := hs[top]
			var width int
			if len(st) == 0 {
				width = i
			} else {
				width = i - st[len(st)-1] - 1
			}
			if height*width > best {
				best = height * width
			}
		}
		st = append(st, i)
	}
	return best
}

// TrapStack computes trapped rain water with a decreasing stack of indices.
// When a taller bar arrives, it forms a container with the bar below the popped
// one; add the bounded water layer by layer.
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	return water
}

// SumSubarrayMins returns the sum over all subarrays of their minimum (mod 1e9+7).
// Contribution technique: each element is the min of (left)*(right) subarrays,
// where left/right are distances to the previous/next smaller elements (found
// with monotonic stacks). Strict on one side avoids double counting duplicates.
const MOD int64 = 1_000_000_007

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	var ans int64
	for i := 0; i < n; i++ {
		minValue := int64(arr[i])
		for j := i; j < n; j++ {
			if int64(arr[j]) < minValue {
				minValue = int64(arr[j])
			}

			ans = (ans + minValue)
		}
	}
	return int(ans)
}

// RemoveKdigits removes k digits from num to make the smallest possible number.
// Greedily pop larger preceding digits (increasing stack), then trim, then strip
// leading zeros.
func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}
	stack := make([]byte, 0, len(num))
	for i := 0; i < len(num); i++ {
		digit := num[i]
		// Remove larger digits from the left
		// while we still have digits to remove.
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > digit {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	// If k digits are still left to remove,
	// remove them from the end.
	if k > 0 {
		stack = stack[:len(stack)-k]
	}
	// Remove leading zeros.
	result := strings.TrimLeft(string(stack), "0")

	if result == "" {
		return "0"
	}
	return result
}

// StockSpanner returns, for each new price, the number of consecutive prior days
// (including today) with price <= today's. A decreasing stack of (price, span).
type StockSpanner struct {
	prices []int
	spans  []int
}

func NewStockSpanner() *StockSpanner { return &StockSpanner{} }

func (s *StockSpanner) Next(price int) int {
	span := 1
	for len(s.prices) > 0 && s.prices[len(s.prices)-1] <= price {
		span += s.spans[len(s.spans)-1]
		s.prices = s.prices[:len(s.prices)-1]
		s.spans = s.spans[:len(s.spans)-1]
	}
	s.prices = append(s.prices, price)
	s.spans = append(s.spans, span)
	return span
}
