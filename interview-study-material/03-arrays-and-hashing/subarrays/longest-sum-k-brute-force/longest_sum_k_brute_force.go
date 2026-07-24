package main

import "fmt"

// Simple brute force solution: Check all possible subarrays
func largestSubarraySumKSimple(arr []int, k int) int {
	maxLength := 0

	// Try all possible starting positions
	for i := 0; i < len(arr); i++ {
		sum := 0

		// Try all possible ending positions from i
		for j := i; j < len(arr); j++ {
			sum += arr[j]

			// If sum equals k, update maxLength
			if sum == k {
				length := j - i + 1
				if length > maxLength {
					maxLength = length
				}
			}
		}
	}

	return maxLength
}

func main() {
	// Test cases
	arr1 := []int{1, 10, 5, 2, 1, 7, 1, 9}
	k1 := 15
	result1 := largestSubarraySumKSimple(arr1, k1)
	fmt.Printf("Array: %v\n", arr1)
	fmt.Printf("k = %d, Largest subarray length: %d\n", k1, result1)
	fmt.Println()

	arr2 := []int{1, 2, 3, 4, 5}
	k2 := 9
	result2 := largestSubarraySumKSimple(arr2, k2)
	fmt.Printf("Array: %v\n", arr2)
	fmt.Printf("k = %d, Largest subarray length: %d\n", k2, result2)
	fmt.Println()

	arr3 := []int{-5, 8, -14, 2, 4, 12}
	k3 := -5
	result3 := largestSubarraySumKSimple(arr3, k3)
	fmt.Printf("Array: %v\n", arr3)
	fmt.Printf("k = %d, Largest subarray length: %d\n", k3, result3)
}
