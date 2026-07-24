package main

import (
	"fmt"
	"runtime"
)

func parallelSum(nums []int) int {
	n := runtime.NumCPU()
	size := len(nums) / n
	results := make(chan int, n)

	for i := 0; i < n; i++ {
		start := i * size
		end := start + size
		if i == n-1 {
			end = len(nums)
		}

		go func(part []int) {
			sum := 0
			for _, v := range part {
				sum += v
			}
			results <- sum
		}(nums[start:end])
	}

	total := 0
	for i := 0; i < n; i++ {
		total += <-results
	}
	return total
}

func main() {
	numbers := make([]int, 0, 10000)

	for i := 1; i <= 10000; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(parallelSum(numbers))
}
