// Package dp contains dynamic-programming interview problems.
//
// DP applies when a problem has (1) OVERLAPPING SUBPROBLEMS - the same smaller
// problem recurs - and (2) OPTIMAL SUBSTRUCTURE - the optimal answer is built
// from optimal answers to subproblems. The recipe: define the state, write the
// recurrence (transition), set base cases, and choose an evaluation order
// (top-down memoized recursion, or bottom-up table). Then optimize space.
package dp

// ClimbStairs: ways to climb n steps taking 1 or 2 at a time.
// ways(n) = ways(n-1) + ways(n-2) (Fibonacci). Only the last two states matter.
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Rob (House Robber): max sum with no two adjacent houses.
// dp[i] = max(dp[i-1], dp[i-2] + nums[i]). Rolling variables give O(1) space.
func Rob(nums []int) int {
	prev, cur := 0, 0
	for _, x := range nums {
		prev, cur = cur, max(cur, prev+x)
	}
	return cur
}

// CoinChange: fewest coins to make amount, or -1. Unbounded knapsack.
// dp[a] = 1 + min over coins c of dp[a-c].
func CoinChange(coins []int, amount int) int {
	const inf = 1 << 30
	dp := make([]int, amount+1)
	for a := 1; a <= amount; a++ {
		dp[a] = inf
		for _, c := range coins {
			if c <= a && dp[a-c]+1 < dp[a] {
				dp[a] = dp[a-c] + 1
			}
		}
	}
	if dp[amount] >= inf {
		return -1
	}
	return dp[amount]
}

// LengthOfLIS: length of the longest strictly increasing subsequence.
// O(n^2) dp[i] = longest ending at i; the O(n log n) patience-sorting version is
// in the comment below.
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// Patience sorting: tails[k] = smallest tail of an increasing subseq of len k+1.
	var tails []int
	for _, x := range nums {
		lo, hi := 0, len(tails)
		for lo < hi { // lower_bound
			mid := (lo + hi) / 2
			if tails[mid] < x {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		if lo == len(tails) {
			tails = append(tails, x)
		} else {
			tails[lo] = x
		}
	}
	return len(tails)
}

// LongestCommonSubsequence: length of the LCS of two strings.
// dp[i][j] = LCS of a[:i], b[:j]. If chars match, +1 diagonal; else max of
// dropping one char from either string.
func LongestCommonSubsequence(a, b string) int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// Knapsack01: max value with total weight <= capacity, each item used at most
// once. 1D dp over capacity, iterated DOWNWARD so each item is counted once.
func Knapsack01(weights, values []int, capacity int) int {
	dp := make([]int, capacity+1)
	for i := range weights {
		for c := capacity; c >= weights[i]; c-- {
			if dp[c-weights[i]]+values[i] > dp[c] {
				dp[c] = dp[c-weights[i]] + values[i]
			}
		}
	}
	return dp[capacity]
}

// UniquePaths: number of paths from top-left to bottom-right moving only right or
// down. dp[j] += dp[j-1] over rows; a single row suffices.
func UniquePaths(m, n int) int {
	dp := make([]int, n)
	for j := range dp {
		dp[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

// WordBreak: can s be segmented into dictionary words?
// dp[i] = true if s[:i] is segmentable; try every split point j with dp[j] and
// s[j:i] in the dictionary.
func WordBreak(s string, wordDict []string) bool {
	words := make(map[string]struct{}, len(wordDict))
	maxLen := 0
	for _, w := range wordDict {
		words[w] = struct{}{}
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0 && i-j <= maxLen; j-- {
			if dp[j] {
				if _, ok := words[s[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)]
}

// MinDistance (Edit Distance): min insert/delete/replace ops to turn a into b.
// dp[i][j] = edits between a[:i] and b[:j].
func MinDistance(a, b string) int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i // delete all of a[:i]
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // insert all of b[:j]
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
			}
		}
	}
	return dp[m][n]
}

// CanPartition: can the array be split into two equal-sum subsets?
// Subset-sum to total/2; 1D boolean knapsack.
func CanPartition(nums []int) bool {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, x := range nums {
		for c := target; c >= x; c-- {
			if dp[c-x] {
				dp[c] = true
			}
		}
	}
	return dp[target]
}
