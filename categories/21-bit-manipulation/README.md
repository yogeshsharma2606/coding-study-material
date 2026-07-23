# 21 - Bit Manipulation

> Code: [`bitops.go`](bitops.go) - Tests: [`bitops_test.go`](bitops_test.go) - run `go test ./categories/21-bit-manipulation/`

## Overview & mental model

Bits let you manipulate sets and numbers in O(1) machine-word operations. The identities you must have memorized:

| Trick | Meaning |
|---|---|
| `x ^ x = 0`, `x ^ 0 = x` | XOR is self-canceling - pairs vanish |
| `x & (x-1)` | clears the **lowest set bit** |
| `x & (-x)` | isolates the **lowest set bit** |
| `1 << k` | mask for bit k |
| `x & (1<<k)` | test bit k |
| `x \| (1<<k)` | set bit k; `x &^ (1<<k)` clears it (Go's AND-NOT) |
| `x >> 1` | divide by 2; `x << 1` multiply by 2 |

XOR's self-canceling property is the workhorse for "find the element that doesn't pair up".

## How to recognize it

- "Appears once/twice/three times", "single number", "missing number".
- "Count set bits", "power of two", "reverse bits".
- "Without using +/-", "using O(1) space", "subsets via bitmask".
- Tight constraints hinting at O(1) space or O(bits) time.

## How to think / attack plan

1. Is there a **pairing/cancellation** structure? XOR everything.
2. Need to process each set bit? Use `x &= x-1` to iterate only over set bits.
3. Building a bitmask over a small set (n <= ~20)? Enumerate `0..(1<<n)-1` as subset masks.
4. Watch signedness and width - in Go, use `uint`/`uint32` for pure bit problems to avoid sign-extension surprises.

## How to enumerate subsets with a bitmask

```go
for mask := 0; mask < (1 << n); mask++ {
    for i := 0; i < n; i++ {
        if mask&(1<<i) != 0 { /* element i is in this subset */ }
    }
}
```

---

## Problems (easy -> hard)

### 1. Single Number (easy)
**Direction of thinking.** All elements pair up except one. XOR the whole array; pairs cancel to 0, leaving the unique value - O(1) space, no hash set.
**Complexity.** Time O(n), space O(1).

### 2. Number of 1 Bits / Hamming Weight (easy)
**Direction of thinking.** `x & (x-1)` removes the lowest set bit, so the loop runs exactly (number of set bits) times - faster than checking all 32 bits.
**Complexity.** Time O(set bits), space O(1).

### 3. Counting Bits 0..n (medium)
**Direction of thinking.** `bits[i] = bits[i>>1] + (i&1)`: i without its last bit is `i/2`, whose count we already have; add the last bit. DP over bits.
**Complexity.** Time O(n), space O(n).

### 4. Reverse Bits (easy)
**Thinking.** Shift the result left and OR in the current lowest bit of x, 32 times.
**Complexity.** Time O(32), space O(1).

### 5. Missing Number (easy)
**Direction of thinking.** XOR all indices `0..n` with all values; each present number cancels its index, leaving the missing number. (Sum formula `n(n+1)/2 - sum` also works but can overflow.)
**Complexity.** Time O(n), space O(1).

### 6. Single Number II - appears 3x (medium)
**Direction of thinking.** Simulate a base-3 counter per bit using two accumulators (`ones`, `twos`): a bit seen a third time is cleared from both. Go's `&^` (AND-NOT) implements "add unless already counted".
**Complexity.** Time O(n), space O(1).

### 7. Sum of Two Integers - no +/- (medium)
**Direction of thinking.** `a ^ b` is addition **without carry**; `(a & b) << 1` is the **carry**. Loop, folding the carry in, until there's no carry left. This is how an adder works in hardware.
**Dry run.** 1+2: xor=3, carry=0 -> 3.
**Complexity.** Time O(bits), space O(1).

### 8. Power of Two (easy)
**Thinking.** A power of two has exactly one set bit, so `n > 0 && n&(n-1) == 0`.
**Complexity.** Time O(1).

### 9. Bitwise AND of Numbers Range (medium)
**Direction of thinking.** ANDing a whole range zeroes any bit that flips within it; only the **common high prefix** of `left` and `right` survives. Right-shift both until equal (finding the common prefix), then shift back.
**Dry run.** [5,7] = 101,110,111 -> common prefix 100 = 4.
**Complexity.** Time O(bits), space O(1).

---

## Common pitfalls & edge cases

- Sign extension: right-shifting a negative signed int fills with 1s. Use unsigned types for bit puzzles.
- `n & (n-1)` on `n=0` - guard the zero case for power-of-two.
- Overflow with the sum-formula approach to Missing Number (XOR avoids it).
- Off-by-one in bit width (32 vs 64).

## Interview Q&A

- **Why does XOR find the single number?** XOR is commutative/associative and self-canceling, so duplicates annihilate.
- **What does `x & (x-1)` do and why is it useful?** Clears the lowest set bit; lets you iterate once per set bit (Brian Kernighan's algorithm).
- **How to add without arithmetic operators?** XOR for the sum bits, AND-shift for the carry, iterate.
- **When is a bitmask the right representation?** Small universes (<= ~20-30 elements) where you need fast set membership/union/intersection or subset enumeration.
