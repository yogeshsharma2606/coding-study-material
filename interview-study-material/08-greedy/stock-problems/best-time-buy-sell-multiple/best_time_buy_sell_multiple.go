package main

import "fmt"

func maxProfit1(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4})) // Output: 7
	fmt.Println(maxProfit1([]int{1, 2, 3, 4, 5}))    // Output: 4
	fmt.Println(maxProfit1([]int{7, 6, 4, 3, 1}))    // Output: 0
}
