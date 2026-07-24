package main

import "fmt"

func stockSpan(prices []int) []int {
	n := len(prices)
	result := make([]int, n)
	stack := []int{} // stack stores indices

	for i := 0; i < n; i++ {
		// Pop while current price >= stack top price
		for len(stack) > 0 && prices[stack[len(stack)-1]] <= prices[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			result[i] = i + 1
		} else {
			result[i] = i - stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return result
}

func main() {
	prices := []int{100, 80, 60, 70, 60, 75, 85} // 1, 1, 1, 2, 1, 4, 6
	fmt.Println("Prices:", prices)
	fmt.Println("Spans:", stockSpan(prices))
}
