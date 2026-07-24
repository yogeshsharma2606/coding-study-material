package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	indexes := []int{} // stores indices

	for i := 0; i < len(nums); i++ {
		// remove out-of-window elements
		if len(indexes) > 0 && indexes[0] == i-k {
			indexes = indexes[1:]
		}

		// remove smaller elements
		for len(indexes) > 0 && nums[indexes[len(indexes)-1]] < nums[i] {
			indexes = indexes[:len(indexes)-1]
		}

		indexes = append(indexes, i)

		if i >= k-1 {
			result = append(result, nums[indexes[0]])
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k)) // Output: [3, 3, 5, 5, 6, 7]
}
