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
func NextGreaterElementI(nums1, nums2 []int) []int {
	nextGreater := make(map[int]int, len(nums2))
	var st []int // decreasing stack of values
	for _, x := range nums2 {
		for len(st) > 0 && st[len(st)-1] < x {
			nextGreater[st[len(st)-1]] = x
			st = st[:len(st)-1]
		}
		st = append(st, x)
	}
	out := make([]int, len(nums1))
	for i, x := range nums1 {
		if g, ok := nextGreater[x]; ok {
			out[i] = g
		} else {
			out[i] = -1
		}
	}
	return out
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
func DailyTemperatures(temps []int) []int {
	res := make([]int, len(temps))
	var st []int // indices, temps decreasing
	for i, t := range temps {
		for len(st) > 0 && temps[st[len(st)-1]] < t {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			res[j] = i - j
		}
		st = append(st, i)
	}
	return res
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
func TrapStack(height []int) int {
	var st []int
	water := 0
	for i, h := range height {
		for len(st) > 0 && height[st[len(st)-1]] < h {
			bottom := st[len(st)-1]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			left := st[len(st)-1]
			width := i - left - 1
			boundedHeight := min(height[left], h) - height[bottom]
			water += width * boundedHeight
		}
		st = append(st, i)
	}
	return water
}

// SumSubarrayMins returns the sum over all subarrays of their minimum (mod 1e9+7).
// Contribution technique: each element is the min of (left)*(right) subarrays,
// where left/right are distances to the previous/next smaller elements (found
// with monotonic stacks). Strict on one side avoids double counting duplicates.
func SumSubarrayMins(arr []int) int {
	const mod = 1_000_000_007
	n := len(arr)
	left := make([]int, n)  // distance to previous element < arr[i] (strict)
	right := make([]int, n) // distance to next element <= arr[i]
	var st []int
	for i := 0; i < n; i++ {
		for len(st) > 0 && arr[st[len(st)-1]] >= arr[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			left[i] = i + 1
		} else {
			left[i] = i - st[len(st)-1]
		}
		st = append(st, i)
	}
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && arr[st[len(st)-1]] > arr[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			right[i] = n - i
		} else {
			right[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum = (sum + arr[i]*left[i]%mod*right[i]) % mod
	}
	return sum
}

// RemoveKdigits removes k digits from num to make the smallest possible number.
// Greedily pop larger preceding digits (increasing stack), then trim, then strip
// leading zeros.
func RemoveKdigits(num string, k int) string {
	var st []byte // digits, non-decreasing
	for i := 0; i < len(num); i++ {
		c := num[i]
		for k > 0 && len(st) > 0 && st[len(st)-1] > c {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, c)
	}
	st = st[:len(st)-k] // remove remaining from the end if still positive
	// strip leading zeros
	i := 0
	for i < len(st) && st[i] == '0' {
		i++
	}
	if i == len(st) {
		return "0"
	}
	return string(st[i:])
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
