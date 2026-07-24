package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// Calculate area
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		width := right - left
		area := h * width

		if area > maxArea {
			maxArea = area
		}

		// Move the pointer at the shorter line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Maximum water container area:", maxArea(heights)) // Output: 49
}
