# 22 - Math & Number Theory

> Code: [`mathx.go`](mathx.go) - Tests: [`mathx_test.go`](mathx_test.go) - run `go test ./categories/22-math-number-theory/`

## Overview & mental model

Math problems check whether you know the standard tools and can handle **overflow** carefully:

- **Euclid's GCD**: `gcd(a,b) = gcd(b, a mod b)` in O(log). LCM from GCD.
- **Sieve of Eratosthenes**: all primes < n in O(n log log n).
- **Fast (binary) exponentiation**: `x^n` in O(log n) by squaring.
- **Overflow discipline**: check bounds *before* a multiply/add, or use a wider type. Note Go's `%` keeps the dividend's sign.

## How to recognize it

- "GCD/LCM", "primes/prime factors", "power/x^n", "reverse/palindrome number".
- "Roman numerals", "Excel columns", "happy number", "digit manipulation".
- Very large exponents/ranges (up to 1e9-1e18) hinting at O(log) methods.

## How to think / attack plan

1. Is there a closed-form or classic algorithm (GCD, sieve, fast pow)? Use it.
2. Digit problems: extract digits with `%10` and `/10`; avoid string conversion when asked.
3. **Always ask about overflow** for 32-bit constraints and guard before the operation.
4. Cycle in a numeric process (happy number)? Reuse Floyd's cycle detection.

---

## Problems (easy -> hard)

### 1. GCD & LCM (easy)
**Direction of thinking.** Euclid: repeatedly replace `(a,b)` with `(b, a mod b)` until b is 0. `lcm = a/gcd*b` (divide first to avoid overflow).
**Complexity.** Time O(log min(a,b)).

### 2. Count Primes (medium) - Sieve
**Direction of thinking.** Mark composites: for each prime p, cross out `p*p, p*p+p, ...` (smaller multiples were already marked by smaller primes). Starting at `p*p` is the key optimization.
**Dry run.** n=10: primes 2,3,5,7 -> 4.
**Complexity.** Time O(n log log n), space O(n).

### 3. Pow(x, n) (medium) - fast exponentiation
**Direction of thinking.** `x^n = (x^2)^(n/2)`; multiply the result by x whenever the current exponent bit is 1. O(log n) instead of O(n). Handle negative n by inverting x.
**Dry run.** `2^10`: bits of 10=1010 -> multiply at bits set -> 1024.
**Complexity.** Time O(log n), space O(1).

### 4. Reverse Integer (medium) - overflow
**Direction of thinking.** Peel digits with `%10`/`/10` and rebuild. The whole point is the **overflow check before** `res*10+digit`, comparing against `INT_MAX/10`.
**Complexity.** Time O(digits), space O(1).

### 5. Palindrome Number (easy)
**Direction of thinking.** Reverse only **half** the digits and compare to the other half - avoids overflow and string conversion. Negative numbers and trailing zeros are non-palindromes.
**Complexity.** Time O(digits), space O(1).

### 6. Roman to Integer (easy)
**Direction of thinking.** Scan left to right; if a symbol's value is less than the next symbol's, it's a **subtractive** pair (IV, IX), so subtract it; otherwise add.
**Complexity.** Time O(n), space O(1).

### 7. Integer to Roman (medium)
**Direction of thinking.** Greedily subtract the largest possible value, using a table that **includes the subtractive forms** (900=CM, 400=CD, ...). Append symbols until num is 0.
**Complexity.** Time O(1) (bounded to 3999), space O(1).

### 8. Happy Number (easy)
**Direction of thinking.** Repeatedly replace n with the sum of squares of its digits. Either it reaches 1, or it loops - detect the cycle with **Floyd's tortoise/hare** (O(1) space, no set).
**Complexity.** Time O(log n) per step, space O(1).

---

## Common pitfalls & edge cases

- Overflow: guard **before** the multiply/add; `math.MaxInt32` boundaries.
- Go's `%` returns a result with the dividend's sign - normalize for modular arithmetic.
- Sieve: start marking at `p*p`, and size the boolean array to n (exclusive/inclusive matters).
- Fast pow with negative exponents and `n = math.MinInt` edge (negating overflows - use a wider type or handle separately).

## Interview Q&A

- **Why does Euclid's algorithm work?** Any common divisor of a and b also divides `a mod b`, so the GCD is preserved each step; the remainder strictly shrinks.
- **Why start the sieve at p*p?** All smaller multiples of p have a smaller prime factor and were already marked.
- **How is fast exponentiation O(log n)?** Each step halves the exponent (one bit), doing O(1) work.
- **How do you avoid overflow when reversing an integer?** Compare the running result to `INT_MAX/10` before appending the next digit.
