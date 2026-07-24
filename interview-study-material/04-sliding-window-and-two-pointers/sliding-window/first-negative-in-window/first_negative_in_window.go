package main

import "fmt"

func firstNegativeInWindow(arr []int, k int) []int {
	var result []int
	var indexes []int // store indices of negative numbers

	for i := 0; i < len(arr); i++ {
		// Remove indices out of current window
		if len(indexes) > 0 && indexes[0] <= i-k {
			indexes = indexes[1:]
		}

		// Add current index if negative
		if arr[i] < 0 {
			indexes = append(indexes, i)
		}

		// Record answer once we have a full window
		if i >= k-1 {
			if len(indexes) > 0 {
				result = append(result, arr[indexes[0]])
			} else {
				result = append(result, 0)
			}
		}
	}
	return result
}

func main() {
	arr := []int{12, -1, -7, 8, -15, 30, 16, 28}
	k := 3
	fmt.Println(firstNegativeInWindow(arr, k)) // Output: [-1 -1 -7 -15 -15 0]
}
