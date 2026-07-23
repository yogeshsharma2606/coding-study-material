// Package mathx contains math & number-theory interview problems.
//
// These test whether you know the standard tools: Euclid's GCD, the Sieve of
// Eratosthenes for primes, fast (binary) exponentiation, and careful OVERFLOW
// handling. The recurring interview trap is integer overflow - check bounds
// before multiplying/adding, or use a wider type.
package mathx

// GCD via the Euclidean algorithm: gcd(a,b) = gcd(b, a mod b). O(log(min)).
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM(a,b) = a / gcd * b (divide first to reduce overflow risk).
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / GCD(a, b) * b
}

// CountPrimes counts primes < n using the Sieve of Eratosthenes: mark multiples
// of each prime starting at p*p. O(n log log n).
func CountPrimes(n int) int {
	if n < 3 {
		return 0
	}
	composite := make([]bool, n)
	count := 0
	for p := 2; p < n; p++ {
		if composite[p] {
			continue
		}
		count++
		for m := p * p; m < n; m += p {
			composite[m] = true
		}
	}
	return count
}

// MyPow computes x^n via binary (fast) exponentiation: square the base, halve
// the exponent. O(log n) multiplications. Handles negative n.
func MyPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}
	return result
}

// ReverseInteger reverses the digits of a 32-bit signed int; returns 0 on
// overflow. Check bounds BEFORE the multiply-add that would overflow.
func ReverseInteger(x int) int {
	const intMax, intMin = 1<<31 - 1, -(1 << 31)
	res := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		if res > intMax/10 || (res == intMax/10 && digit > 7) {
			return 0
		}
		if res < intMin/10 || (res == intMin/10 && digit < -8) {
			return 0
		}
		res = res*10 + digit
	}
	return res
}

// IsPalindromeNumber reports whether an integer reads the same backward, without
// converting to a string. Reverse only half the digits.
func IsPalindromeNumber(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}
	return x == rev || x == rev/10 // even or odd digit count
}

// RomanToInt converts a Roman numeral to an integer. Scan left to right; if a
// symbol is smaller than the next, subtract it (e.g. IV = 5-1).
func RomanToInt(s string) int {
	val := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	total := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && val[s[i]] < val[s[i+1]] {
			total -= val[s[i]]
		} else {
			total += val[s[i]]
		}
	}
	return total
}

// IntToRoman converts an integer (1..3999) to a Roman numeral by greedily
// subtracting the largest value symbols, including the subtractive forms.
func IntToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var sb []byte
	for i, v := range vals {
		for num >= v {
			sb = append(sb, syms[i]...)
			num -= v
		}
	}
	return string(sb)
}

// IsHappy reports whether repeatedly summing squares of digits reaches 1.
// Detect cycles with Floyd's tortoise/hare on the "next" function.
func IsHappy(n int) bool {
	next := func(x int) int {
		sum := 0
		for x > 0 {
			d := x % 10
			sum += d * d
			x /= 10
		}
		return sum
	}
	slow, fast := n, next(n)
	for fast != 1 && slow != fast {
		slow = next(slow)
		fast = next(next(fast))
	}
	return fast == 1
}
