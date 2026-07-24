package main

import (
	"fmt"
)

const MOD = 1000000007
const MAXP = 1000000

// Sieve of Eratosthenes
func buildPrimeSieve() []bool {
	isPrime := make([]bool, MAXP+1)
	for i := 2; i <= MAXP; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= MAXP; i++ {
		if isPrime[i] {
			for j := i * i; j <= MAXP; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}

func countPrimeSplits(s string) int {

	n := len(s)
	isPrime := buildPrimeSieve()

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {

		val := 0
		mult := 1

		for j := i - 1; j >= 0 && i-j <= 7; j-- {

			digit := int(s[j] - '0')
			val = digit*mult + val
			mult *= 10

			if s[j] == '0' {
				continue // leading zero
			}

			if val >= 2 && val <= MAXP && isPrime[val] {
				dp[i] = (dp[i] + dp[j]) % MOD
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(countPrimeSplits("11375")) // 3
	fmt.Println(countPrimeSplits("3175"))  // Example 2
}
