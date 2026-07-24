package main

import (
	"fmt"
)

/*
ALGORITHM: Find Largest Subarray with Sum K

KEY INSIGHT: If prefix_sum[i] - prefix_sum[j] = k, then subarray from (j+1) to i has sum k

Example: arr = [10, 5, 2, 7, 1, 9], k = 15

STEP-BY-STEP WALKTHROUGH:

Index:  0   1   2   3   4   5
Array: [10,  5,  2,  7,  1,  9]
        ──────────────────────

┌─────────────────────────────────────────────────────────────────┐
│ ITERATION 0: i=0, arr[0]=10                                     │
├─────────────────────────────────────────────────────────────────┤
│ sum = 0 + 10 = 10                                               │
│                                                                  │
│ Check: sum == k? 10 == 15? NO                                  │
│ Check: prefixSum[10-15] = prefixSum[-5]? NOT FOUND             │
│ Store: prefixSum[10] = 0                                        │
│                                                                  │
│ State:                                                           │
│   sum = 10                                                       │
│   prefixSum = {10: 0}                                           │
│   maxLength = 0                                                  │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│ ITERATION 1: i=1, arr[1]=5                                      │
├─────────────────────────────────────────────────────────────────┤
│ sum = 10 + 5 = 15                                               │
│                                                                  │
│ Check: sum == k? 15 == 15? YES! ✓                              │
│   → Subarray [0..1] = [10, 5] has sum 15                       │
│   → maxLength = max(0, 2) = 2                                   │
│                                                                  │
│ Check: prefixSum[15-15] = prefixSum[0]? NOT FOUND              │
│ Store: prefixSum[15] = 1                                        │
│                                                                  │
│ State:                                                           │
│   sum = 15                                                       │
│   prefixSum = {10: 0, 15: 1}                                    │
│   maxLength = 2                                                  │
│                                                                  │
│ Visual:                                                          │
│   [10, 5, 2, 7, 1, 9]                                           │
│    ────                                                          │
│    sum=15 ✓                                                      │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│ ITERATION 2: i=2, arr[2]=2                                      │
├─────────────────────────────────────────────────────────────────┤
│ sum = 15 + 2 = 17                                               │
│                                                                  │
│ Check: sum == k? 17 == 15? NO                                   │
│ Check: prefixSum[17-15] = prefixSum[2]? NOT FOUND              │
│ Store: prefixSum[17] = 2                                        │
│                                                                  │
│ State:                                                           │
│   sum = 17                                                       │
│   prefixSum = {10: 0, 15: 1, 17: 2}                             │
│   maxLength = 2                                                  │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│ ITERATION 3: i=3, arr[3]=7                                      │
├─────────────────────────────────────────────────────────────────┤
│ sum = 17 + 7 = 24                                               │
│                                                                  │
│ Check: sum == k? 24 == 15? NO                                   │
│ Check: prefixSum[24-15] = prefixSum[9]? NOT FOUND              │
│ Store: prefixSum[24] = 3                                        │
│                                                                  │
│ State:                                                           │
│   sum = 24                                                       │
│   prefixSum = {10: 0, 15: 1, 17: 2, 24: 3}                      │
│   maxLength = 2                                                  │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│ ITERATION 4: i=4, arr[4]=1                                      │
├─────────────────────────────────────────────────────────────────┤
│ sum = 24 + 1 = 25                                               │
│                                                                  │
│ Check: sum == k? 25 == 15? NO                                   │
│ Check: prefixSum[25-15] = prefixSum[10]? YES! ✓                │
│   → prefixSum[10] = 0                                           │
│   → Subarray from (0+1) to 4 = [1..4] = [5, 2, 7, 1]           │
│   → Length = 4 - 0 = 4                                          │
│   → maxLength = max(2, 4) = 4                                   │
│                                                                  │
│ Store: prefixSum[25] = 4                                        │
│                                                                  │
│ State:                                                           │
│   sum = 25                                                       │
│   prefixSum = {10: 0, 15: 1, 17: 2, 24: 3, 25: 4}              │
│   maxLength = 4                                                  │
│                                                                  │
│ Visual:                                                          │
│   [10, 5, 2, 7, 1, 9]                                           │
│        ────────                                                  │
│        sum=15 ✓ (prefixSum[25] - prefixSum[10] = 15)            │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│ ITERATION 5: i=5, arr[5]=9                                      │
├─────────────────────────────────────────────────────────────────┤
│ sum = 25 + 9 = 34                                               │
│                                                                  │
│ Check: sum == k? 34 == 15? NO                                   │
│ Check: prefixSum[34-15] = prefixSum[19]? NOT FOUND             │
│ Store: prefixSum[34] = 5                                        │
│                                                                  │
│ State:                                                           │
│   sum = 34                                                       │
│   prefixSum = {10: 0, 15: 1, 17: 2, 24: 3, 25: 4, 34: 5}       │
│   maxLength = 4                                                  │
└─────────────────────────────────────────────────────────────────┘

RESULT: maxLength = 4
Subarray [5, 2, 7, 1] has sum 15 and is the longest.
*/

func largestSubarraySumK(arr []int, k int) int {
	// Map to store prefix sum -> earliest index where this sum occurred
	prefixSum := make(map[int]int)
	sum := 0
	maxLength := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i] // Calculate prefix sum up to index i

		// Case 1: If prefix sum equals k, subarray from 0 to i has sum k
		if sum == k {
			maxLength = i + 1
		}

		// Case 2: If (sum - k) exists in map, there's a subarray with sum k
		// Why? If prefixSum[i] - prefixSum[j] = k, then subarray (j+1..i) has sum k
		// We found prefixSum[j] = sum - k, so subarray from (j+1) to i has sum k
		if startIdx, exists := prefixSum[sum-k]; exists {
			length := i - startIdx
			if length > maxLength {
				maxLength = length
			}
		}

		// Store the earliest index for this prefix sum
		// Only store if it doesn't exist (we want earliest index for maximum length)
		if _, exists := prefixSum[sum]; !exists {
			prefixSum[sum] = i
		}
	}

	return maxLength
}

func main() {
	// Other test cases
	fmt.Println("\nOther test cases:")
	arr1 := []int{5, 5, 5, 1, 1, 3, 6, 2, 2}
	k1 := 15
	fmt.Printf("Array: %v, k: %d, Largest subarray length: %d\n", arr1, k1, largestSubarraySumK(arr1, k1))

	arr3 := []int{-5, 8, -14, 2, 4, 12}
	k3 := -5
	fmt.Printf("Array: %v, k: %d, Largest subarray length: %d\n", arr3, k3, largestSubarraySumK(arr3, k3))
}
