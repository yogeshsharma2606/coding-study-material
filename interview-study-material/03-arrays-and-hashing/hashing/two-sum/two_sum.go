package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value → index

	for i, num := range nums {
		complement := target - num
		if j, found := seen[complement]; found {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil // no solution found
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 11
	fmt.Println(twoSum(nums, target)) // Output: [0, 1]
}
