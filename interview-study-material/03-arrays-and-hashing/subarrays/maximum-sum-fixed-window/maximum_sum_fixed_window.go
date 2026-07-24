package main

import "fmt"

func maxSumSubarray(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return -1 // not enough elements
	}

	// Compute sum of first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum := windowSum

	// Slide the window
	for i := k; i < n; i++ {
		windowSum += arr[i] - arr[i-k] // add new, remove old
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func main() {
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3
	fmt.Println("Maximum Sum Subarray of Size K:", maxSumSubarray(arr, k)) // Output: 9
}
