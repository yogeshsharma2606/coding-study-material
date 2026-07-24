package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{} // stack stores indices

	for i := 0; i < n; i++ {
		// Compare current temp with stack top
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}

	return result
}

func main() {
	temps := []int{73, 74, 75, 71, 69, 72, 76, 73} // 1, 1, 4, 2, 1, 1, 0, 0
	fmt.Println("Input:", temps)
	fmt.Println("Output:", dailyTemperatures(temps))
}
