// Package greedy contains greedy interview problems.
//
// A greedy algorithm builds a solution by always taking the locally optimal
// choice, never reconsidering. It's fast (usually O(n) or O(n log n)) but only
// CORRECT when the problem has the "greedy-choice property" - a local optimum
// leads to a global optimum. The senior skill is proving (or arguing via an
// EXCHANGE ARGUMENT) that greed is safe here, and recognizing when it is NOT
// (then you need DP).
package greedy

import "sort"

// CanJump: can you reach the last index? Track the farthest reachable index;
// if you ever stand beyond it, you're stuck.
func CanJump(nums []int) bool {
	farthest := 0
	for i, n := range nums {
		if i > farthest {
			return false
		}
		if i+n > farthest {
			farthest = i + n
		}
	}
	return true
}

// Jump (Jump Game II): min jumps to reach the last index (guaranteed reachable).
// BFS-like level expansion: within the current jump's range, track the farthest
// next reach; when you hit the range end, you must jump.
func Jump(nums []int) int {
	jumps, curEnd, farthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == curEnd {
			jumps++
			curEnd = farthest
		}
	}
	return jumps
}

// CanCompleteCircuit (Gas Station): starting index to complete the loop, or -1.
// If total gas >= total cost, a solution exists; the start is just after the
// point where the running tank dips lowest (goes negative).
func CanCompleteCircuit(gas, cost []int) int {
	total, tank, start := 0, 0, 0
	for i := range gas {
		diff := gas[i] - cost[i]
		total += diff
		tank += diff
		if tank < 0 {
			start = i + 1 // can't start anywhere in [start..i]
			tank = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

// FindContentChildren (Assign Cookies): max children satisfied. Sort both greed
// factors and cookie sizes; give the smallest adequate cookie to each child.
func FindContentChildren(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	child, cookie := 0, 0
	for child < len(g) && cookie < len(s) {
		if s[cookie] >= g[child] {
			child++
		}
		cookie++
	}
	return child
}

// MaxProfitII: unlimited transactions; capture every upward move. Sum all
// positive day-to-day differences.
func MaxProfitII(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// PartitionLabels: split s into as many parts as possible so each letter appears
// in one part. A part must extend to the last occurrence of every letter in it.
func PartitionLabels(s string) []int {
	last := [26]int{}
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}
	var res []int
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}
		if i == end { // every letter so far ends within [start, end]
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

// LeastInterval (Task Scheduler): min time units to run tasks with cooldown n
// between identical tasks. Governed by the most frequent task: arrange it in
// (maxCount-1) gap frames plus the final row.
func LeastInterval(tasks []byte, n int) int {
	var count [26]int
	maxCount := 0
	for _, t := range tasks {
		count[t-'A']++
		if count[t-'A'] > maxCount {
			maxCount = count[t-'A']
		}
	}
	maxCountTasks := 0 // how many tasks share the max frequency
	for _, c := range count {
		if c == maxCount {
			maxCountTasks++
		}
	}
	// (maxCount-1) full frames of size (n+1), plus the last group
	idle := (maxCount-1)*(n+1) + maxCountTasks
	if idle < len(tasks) {
		return len(tasks) // cooldown fully absorbed by other tasks
	}
	return idle
}

// Candy: min candies so each child gets >= 1 and higher-rated children get more
// than a lower-rated neighbor. Two sweeps: left-to-right for the left neighbor,
// right-to-left for the right neighbor; take the max at each position.
func Candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	total := 0
	for _, c := range candies {
		total += c
	}
	return total
}
