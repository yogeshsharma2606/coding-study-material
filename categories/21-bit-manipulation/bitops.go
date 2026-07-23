// Package bitops contains bit-manipulation interview problems.
//
// Bits let you do set-like operations in O(1) machine words. Key identities:
//
//	x ^ x = 0, x ^ 0 = x, XOR is associative/commutative (self-canceling);
//	x & (x-1) clears the lowest set bit; x & (-x) isolates it;
//	1<<k is a mask for bit k; x & (1<<k) tests it; x | (1<<k) sets it.
//
// These turn "find the unique element", "count bits", and "no-arithmetic add"
// into elegant O(1)/O(bits) solutions.
package bitops

// SingleNumber: every element appears twice except one. XOR everything - pairs
// cancel to 0, leaving the unique value. O(1) space (no hash set needed).
func SingleNumber(nums []int) int {
	x := 0
	for _, n := range nums {
		x ^= n
	}
	return x
}

// HammingWeight counts set bits (population count). x & (x-1) removes the lowest
// set bit each step, so the loop runs once per set bit.
func HammingWeight(x uint) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

// CountBits returns, for each i in [0,n], the number of set bits.
// DP: bits[i] = bits[i>>1] + (i&1) (i is i/2 shifted, plus its last bit).
func CountBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i>>1] + (i & 1)
	}
	return res
}

// ReverseBits reverses the 32 bits of x. Pull the low bit off x and push it onto
// the growing result.
func ReverseBits(x uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = (res << 1) | (x & 1)
		x >>= 1
	}
	return res
}

// MissingNumber: [0..n] with one missing. XOR all indices and values; every
// present number cancels with its index, leaving the missing one.
func MissingNumber(nums []int) int {
	x := len(nums) // start with n (the highest index)
	for i, v := range nums {
		x ^= i ^ v
	}
	return x
}

// SingleNumberII: every element appears three times except one. Track bits that
// have appeared once and twice; a bit appearing a third time is cleared.
func SingleNumberII(nums []int) int {
	ones, twos := 0, 0
	for _, n := range nums {
		ones = (ones ^ n) &^ twos // add to ones unless already in twos
		twos = (twos ^ n) &^ ones // add to twos unless now in ones
	}
	return ones
}

// GetSum adds two integers without + or -. XOR is addition without carry; AND<<1
// is the carry. Repeat until no carry.
func GetSum(a, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

// IsPowerOfTwo: a power of two has exactly one set bit, so x & (x-1) == 0.
func IsPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// RangeBitwiseAnd returns the AND of all numbers in [left, right]. The result is
// the common binary prefix of left and right (differing low bits get zeroed by
// some number in the range). Shift both right until equal, then shift back.
func RangeBitwiseAnd(left, right int) int {
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}
