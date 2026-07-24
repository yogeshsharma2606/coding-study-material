package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	mutex            sync.Mutex
	state            State
	failureThreshold int
	failures         int
	lastFailureTime  time.Time
	retryTimeout     time.Duration
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            Closed,
		failureThreshold: threshold,
		retryTimeout:     timeout,
	}
}

func (cb *CircuitBreaker) Execute(request func() (string, error)) (string, error) {
	cb.mutex.Lock()

	// Check if we should move from Open to Half-Open
	if cb.state == Open && time.Since(cb.lastFailureTime) > cb.retryTimeout {
		fmt.Println(">>> Circuit transitioning to HALF-OPEN")
		cb.state = HalfOpen
	}

	if cb.state == Open {
		cb.mutex.Unlock()
		return "", errors.New("circuit is OPEN - request rejected")
	}
	cb.mutex.Unlock()

	// Execute the actual request
	result, err := request()

	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if err != nil {
		cb.failures++
		cb.lastFailureTime = time.Now()
		if cb.failures >= cb.failureThreshold {
			fmt.Println("!!! Threshold reached: Tripping circuit to OPEN")
			cb.state = Open
		}
		return "", err
	}

	// If successful and state was HalfOpen, reset to Closed
	if cb.state == HalfOpen {
		fmt.Println("<<< Service recovered: Closing circuit")
		cb.state = Closed
		cb.failures = 0
	}

	return result, nil
}

func main() {
	// 3 failures allowed, 5 seconds recovery time
	cb := NewCircuitBreaker(3, 5*time.Second)

	// Mocking a failing service call
	failingService := func() (string, error) {
		return "", errors.New("timeout error")
	}

	for i := 1; i <= 7; i++ {
		res, err := cb.Execute(failingService)
		if err != nil {
			fmt.Printf("Attempt %d: %v\n", i, err)
		} else {
			fmt.Printf("Attempt %d: %s\n", i, res)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
