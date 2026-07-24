package main

import "fmt"

// For height = [0,1,0,2,1,0,1,3,2,1,2,1]:

// At index 2: leftMax=1, rightMax=3 → water = min(1,3)-0 = 1

// At index 5: leftMax=2, rightMax=3 → water = min(2,3)-0 = 2

// At index 6: leftMax=2, rightMax=3 → water = min(2,3)-1 = 1

// At index 9: leftMax=3, rightMax=2 → water = min(3,2)-1 = 1

// At index 10: leftMax=3, rightMax=2 → water = min(3,2)-2 = 0

// Total = 6

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5})) // Output: 9
}

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	return water
}
