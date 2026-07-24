package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	longestStreak := 1
	currentStreak := 1

	for i := 1; i < len(nums); i++ {
		// Skip duplicates
		if nums[i] == nums[i-1] {
			continue
		}

		// Check if consecutive
		if nums[i] == nums[i-1]+1 {
			currentStreak++
		} else {
			// Sequence broken, reset
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
			currentStreak = 1
		}
	}

	// Final check for the last sequence
	if currentStreak > longestStreak {
		return currentStreak
	}
	return longestStreak
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2})) // Output: 4 ( [1, 2, 3, 4] )
}
