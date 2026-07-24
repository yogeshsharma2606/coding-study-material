package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(maxRequests int, refillInterval time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, maxRequests),
	}

	// Fill the bucket initially
	for i := 0; i < maxRequests; i++ {
		rl.tokens <- struct{}{}
	}

	// Refill tokens periodically
	go func() {
		ticker := time.NewTicker(refillInterval)
		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
			default: // Bucket full, drop token
			}
		}
	}()

	return rl
}

func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {
	// Allow 3 requests per second
	limiter := NewRateLimiter(3, time.Millisecond*300)

	for i := 1; i <= 20; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d: Action Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Rate Limited (429)\n", i)
		}
		time.Sleep(time.Millisecond * 200)
	}

}
