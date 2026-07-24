package main

import "fmt"

// Kadane's Algorithm
func maxSubArray(arr []int) ([]int, int) {
	maxSoFar := arr[0]
	currSum := arr[0]
	start, end := 0, 0 // indices of best subarray
	tempStart := 0
	for i := 1; i < len(arr); i++ {
		if currSum < 0 {
			currSum = arr[i]
			tempStart = i
		} else {
			currSum += arr[i]
		}
		if currSum > maxSoFar {
			maxSoFar = currSum
			start = tempStart
			end = i
		}
	}
	return arr[start : end+1], maxSoFar
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} // Output: 6
	subarray, maxSum := maxSubArray(arr)
	fmt.Println("Maximum subarray sum:", maxSum)
	fmt.Println("Subarray:", subarray)
}
