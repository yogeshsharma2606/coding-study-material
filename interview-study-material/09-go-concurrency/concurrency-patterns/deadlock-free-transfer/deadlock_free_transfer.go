package main

import (
	"fmt"
	"sync"
)

type Account struct {
	ID      int
	balance int
	mu      sync.Mutex
}

// Transfer safely transfers money without deadlock
func Transfer(a, b *Account, amount int) {
	// Enforce lock ordering by account ID
	first, second := a, b
	if a.ID > b.ID {
		first, second = b, a
	}

	first.mu.Lock()
	defer first.mu.Unlock()

	second.mu.Lock()
	defer second.mu.Unlock()

	// Perform transfer
	a.balance += amount
	b.balance -= amount
}

func main() {
	acc1 := &Account{ID: 1, balance: 1000}
	acc2 := &Account{ID: 2, balance: 1000}

	var wg sync.WaitGroup

	// Concurrent transfers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Transfer(acc1, acc2, 10)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			Transfer(acc2, acc1, 10)
		}()
	}

	wg.Wait()

	fmt.Println("Account 1 balance:", acc1.balance)
	fmt.Println("Account 2 balance:", acc2.balance)
}
