package main

import "fmt"

func subarraySum(nums []int, target int) []int {

	subarray := []int{}

	// Try all possible starting positions
	for i := 0; i < len(nums); i++ {
		sum := 0

		// Try all possible ending positions from i
		for j := i; j < len(nums); j++ {
			sum += nums[j]

			if sum == target {
				subarray = nums[i : j+1]
			}
		}
	}

	return subarray

}

func main() {
	arr := []int{2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Output: [20, 3, 10]
	target := 15
	fmt.Println("Subarray:", subarraySum(arr, target))
}
