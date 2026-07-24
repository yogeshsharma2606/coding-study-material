package main

import "fmt"

// Partition function
func partition(arr []int, low, high int) int {
	pivot := arr[high] // choose last element as pivot
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// QuickSort function
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Original:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted:", arr)
}

/*
DRY RUN — arr = [10, 7, 8, 9, 1, 5], indices 0..5

Step 1: partition(arr, 0, 5)
  pivot = arr[5] = 5
  i = -1, j runs 0..4
  j=0: 10 < 5? no
  j=1: 7 < 5?  no
  j=2: 8 < 5?  no
  j=3: 9 < 5?  no
  j=4: 1 < 5?  yes → i=0, swap arr[0],arr[4] → [1, 7, 8, 9, 10, 5]
  place pivot: swap arr[i+1], arr[high] → swap arr[1], arr[5] → [1, 5, 8, 9, 10, 7]
  return pi = 1

  Recurse: quickSort(0,0) skip; quickSort(2,5) on [8,9,10,7]

Step 2: partition(arr, 2, 5)
  pivot = arr[5] = 7
  j=2: 8<7? no. j=3: 9<7? no. j=4: 10<7? no. i stays 1
  place pivot: swap arr[2], arr[5] → [1, 5, 7, 9, 10, 8], return pi = 2

Step 3: partition(arr, 3, 5)  pivot=8 → no moves, swap arr[3],arr[5] → [1,5,7,8,10,9], return 3
Step 4: partition(arr, 4, 5)  pivot=9 → swap arr[4],arr[5] → [1,5,7,8,9,10], return 4

Result: [1, 5, 7, 8, 9, 10]
*/
