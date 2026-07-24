package main

import (
	"fmt"
)

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotateInline(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n // handle cases where k > n

	// Rotate in-place using reversals
	reverse(nums, 0, n-1) // reverse whole array
	reverse(nums, 0, k-1) // reverse first k elements
	reverse(nums, k, n-1) // reverse remaining elements
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotateInline(arr, 7)
	fmt.Println(arr) // Output: [5 6 7 1 2 3 4]
}
