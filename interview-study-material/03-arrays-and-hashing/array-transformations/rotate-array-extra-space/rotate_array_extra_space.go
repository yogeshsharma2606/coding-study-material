package main

import "fmt"

func rotate(arr []int, k int) []int {
	n := len(arr)
	k = k % n
	fmt.Println(k)
	return append(arr[n-k:], arr[:n-k]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotated := rotate(arr, 3)
	fmt.Println("Rotated array:", rotated)
}
