// Package hashing contains hash-map / hash-set interview problems.
//
// A hash map gives amortized O(1) insert/lookup/delete, letting you trade O(n)
// space for O(1) "have I seen this?" checks - collapsing an O(n^2) search into
// O(n). Use a set (map[T]struct{}) for membership, a map[T]int for counts, and a
// map keyed on a derived signature to group/relate items.
package hashing

// ContainsDuplicate reports whether any value appears at least twice.
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, x := range nums {
		if _, ok := seen[x]; ok {
			return true
		}
		seen[x] = struct{}{}
	}
	return false
}

// FirstUniqChar returns the index of the first non-repeating character, or -1.
// Count all chars in one pass, then scan for the first with count 1.
func FirstUniqChar(s string) int {
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if count[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

// Intersection returns the unique values present in both arrays.
func Intersection(a, b []int) []int {
	set := make(map[int]struct{}, len(a))
	for _, x := range a {
		set[x] = struct{}{}
	}
	var out []int
	for _, x := range b {
		if _, ok := set[x]; ok {
			out = append(out, x)
			delete(set, x) // avoid duplicates in output
		}
	}
	return out
}

// TopKFrequent returns the k most frequent elements using bucket sort by
// frequency: index = frequency, so we avoid a full O(n log n) sort.
func TopKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, x := range nums {
		freq[x]++
	}
	// buckets[f] holds all values with frequency f; f in [1, len(nums)]
	buckets := make([][]int, len(nums)+1)
	for val, f := range freq {
		buckets[f] = append(buckets[f], val)
	}
	var out []int
	for f := len(buckets) - 1; f >= 1 && len(out) < k; f-- {
		for _, v := range buckets[f] {
			out = append(out, v)
			if len(out) == k {
				break
			}
		}
	}
	return out
}

// LongestConsecutive returns the length of the longest run of consecutive
// integers (in any order) in O(n). Put all in a set; only start counting from a
// value with no predecessor (x-1 absent), so each run is walked once.
func LongestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, x := range nums {
		set[x] = struct{}{}
	}
	best := 0
	for x := range set {
		if _, ok := set[x-1]; ok {
			continue // not the start of a run
		}
		length := 1
		for {
			if _, ok := set[x+length]; !ok {
				break
			}
			length++
		}
		if length > best {
			best = length
		}
	}
	return best
}

// IsValidSudoku checks the filled cells of a 9x9 board for row/col/box conflicts
// using nine hash sets each. '.' means empty.
func IsValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9]map[byte]struct{}
	for i := 0; i < 9; i++ {
		rows[i] = map[byte]struct{}{}
		cols[i] = map[byte]struct{}{}
		boxes[i] = map[byte]struct{}{}
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			v := board[r][c]
			if v == '.' {
				continue
			}
			b := (r/3)*3 + c/3
			if _, ok := rows[r][v]; ok {
				return false
			}
			if _, ok := cols[c][v]; ok {
				return false
			}
			if _, ok := boxes[b][v]; ok {
				return false
			}
			rows[r][v] = struct{}{}
			cols[c][v] = struct{}{}
			boxes[b][v] = struct{}{}
		}
	}
	return true
}

// FourSumCount counts tuples (i,j,k,l) with a[i]+b[j]+c[k]+d[l] == 0.
// Split into halves: count sums of a+b in a map, then for each c+d look up the
// negation. O(n^2) instead of O(n^4).
func FourSumCount(a, b, c, d []int) int {
	sumAB := make(map[int]int)
	for _, x := range a {
		for _, y := range b {
			sumAB[x+y]++
		}
	}
	res := 0
	for _, x := range c {
		for _, y := range d {
			res += sumAB[-(x + y)]
		}
	}
	return res
}

// MajorityElement returns the element appearing more than n/2 times using
// Boyer-Moore voting: O(1) space. A majority element survives cancellation.
func MajorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, x := range nums {
		if count == 0 {
			candidate = x
		}
		if x == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
