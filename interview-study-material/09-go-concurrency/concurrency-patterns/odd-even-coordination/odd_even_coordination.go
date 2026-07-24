package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i + 1
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Channel to coordinate turns
	oddCh := make(chan int)
	evenCh := make(chan int)

	// Odd goroutine
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			if num%2 != 0 {
				<-oddCh // wait for signal
				fmt.Println("Odd:", num)
				evenCh <- 1 // signal even
			}
		}
	}()

	// Even goroutine
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			if num%2 == 0 {
				<-evenCh // wait for signal
				fmt.Println("Even:", num)
				if num < n {
					oddCh <- 1 // signal odd
				} else {
					close(evenCh)
					close(oddCh)
				}

			}
		}
	}()

	// Kick off with odd
	oddCh <- 1

	wg.Wait()
}
