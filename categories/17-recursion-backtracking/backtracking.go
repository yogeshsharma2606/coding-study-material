// Package backtracking contains recursion & backtracking interview problems.
//
// Backtracking is systematic brute force over a decision tree: at each step try a
// choice, recurse, then UNDO the choice (backtrack) and try the next. The
// skeleton is choose -> explore -> unchoose. Pruning invalid branches early is
// what keeps it tractable. Use it for "generate all / count all / find any"
// combinations, permutations, partitions, and constraint-satisfaction problems.
package backtracking

import "sort"

// Subsets returns all subsets (the power set). At each index, choose to include
// the element or not; every recursion path is a subset.
func Subsets(nums []int) [][]int {
	var res [][]int
	var cur []int
	var backtrack func(start int)
	backtrack = func(start int) {
		// record a copy of the current subset at every node
		res = append(res, append([]int(nil), cur...))
		for i := start; i < len(nums); i++ {
			cur = append(cur, nums[i]) // choose
			backtrack(i + 1)           // explore
			cur = cur[:len(cur)-1]     // unchoose
		}
	}
	backtrack(0)
	return res
}

// Permutations returns all orderings. Track used elements; each position picks an
// unused element.
func Permutations(nums []int) [][]int {
	var res [][]int
	var cur []int
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		if len(cur) == len(nums) {
			res = append(res, append([]int(nil), cur...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			cur = append(cur, nums[i])
			backtrack()
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	backtrack()
	return res
}

// Combine returns all k-combinations of 1..n.
func Combine(n, k int) [][]int {
	var res [][]int
	var cur []int
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(cur) == k {
			res = append(res, append([]int(nil), cur...))
			return
		}
		// prune: not enough numbers left to reach size k
		for i := start; i <= n-(k-len(cur))+1; i++ {
			cur = append(cur, i)
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(1)
	return res
}

// CombinationSum returns unique combinations summing to target; each candidate
// may be reused. Sort to enable pruning when the remainder goes negative.
func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var cur []int
	var backtrack func(start, remain int)
	backtrack = func(start, remain int) {
		if remain == 0 {
			res = append(res, append([]int(nil), cur...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > remain {
				break // sorted: no larger candidate fits
			}
			cur = append(cur, candidates[i])
			backtrack(i, remain-candidates[i]) // i (not i+1): reuse allowed
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0, target)
	return res
}

// GenerateParenthesis returns all valid combinations of n pairs of parentheses.
// Constraint pruning: can add '(' while open < n, ')' while close < open.
func GenerateParenthesis(n int) []string {
	var res []string
	var backtrack func(cur []byte, open, close int)
	backtrack = func(cur []byte, open, close int) {
		if len(cur) == 2*n {
			res = append(res, string(cur))
			return
		}
		if open < n {
			backtrack(append(cur, '('), open+1, close)
		}
		if close < open {
			backtrack(append(cur, ')'), open, close+1)
		}
	}
	backtrack([]byte{}, 0, 0)
	return res
}

// LetterCombinations returns all letter strings a phone number could spell.
func LetterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	mapping := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	var res []string
	var backtrack func(idx int, cur []byte)
	backtrack = func(idx int, cur []byte) {
		if idx == len(digits) {
			res = append(res, string(cur))
			return
		}
		for _, ch := range mapping[digits[idx]] {
			backtrack(idx+1, append(cur, byte(ch)))
		}
	}
	backtrack(0, []byte{})
	return res
}

// PalindromePartition returns all ways to split s into palindromic substrings.
func PalindromePartition(s string) [][]string {
	var res [][]string
	var cur []string
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			res = append(res, append([]string(nil), cur...))
			return
		}
		for end := start + 1; end <= len(s); end++ {
			if isPalindrome(s, start, end-1) {
				cur = append(cur, s[start:end])
				backtrack(end)
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack(0)
	return res
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// Exist (Word Search): does word exist in the grid via 4-directional adjacency
// without reusing a cell? DFS with in-place marking, restored on backtrack.
func Exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	var dfs func(r, c, idx int) bool
	dfs = func(r, c, idx int) bool {
		if idx == len(word) {
			return true
		}
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[idx] {
			return false
		}
		tmp := board[r][c]
		board[r][c] = '#' // mark visited
		found := dfs(r+1, c, idx+1) || dfs(r-1, c, idx+1) ||
			dfs(r, c+1, idx+1) || dfs(r, c-1, idx+1)
		board[r][c] = tmp // restore
		return found
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

// SolveNQueens returns all distinct board configurations placing n queens so
// none attack each other. Place one queen per row; track threatened columns and
// diagonals for O(1) validity checks.
func SolveNQueens(n int) [][]string {
	var res [][]string
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n) // r+c
	diag2 := make([]bool, 2*n) // r-c+n
	queens := make([]int, n)   // queens[r] = column
	var backtrack func(r int)
	backtrack = func(r int) {
		if r == n {
			res = append(res, buildBoard(queens, n))
			return
		}
		for c := 0; c < n; c++ {
			if cols[c] || diag1[r+c] || diag2[r-c+n] {
				continue
			}
			cols[c], diag1[r+c], diag2[r-c+n] = true, true, true
			queens[r] = c
			backtrack(r + 1)
			cols[c], diag1[r+c], diag2[r-c+n] = false, false, false
		}
	}
	backtrack(0)
	return res
}

func buildBoard(queens []int, n int) []string {
	board := make([]string, n)
	for r := 0; r < n; r++ {
		row := make([]byte, n)
		for c := 0; c < n; c++ {
			if queens[r] == c {
				row[c] = 'Q'
			} else {
				row[c] = '.'
			}
		}
		board[r] = string(row)
	}
	return board
}
