package main

import "fmt"

// Merge two sorted slices
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// MergeSort function
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original:", arr)
	sorted := mergeSort(arr)
	fmt.Println("Sorted:", sorted)

	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{2, 4, 6, 8, 10}
	merged := merge(arr1, arr2)
	fmt.Println("Merged:", merged)
}
