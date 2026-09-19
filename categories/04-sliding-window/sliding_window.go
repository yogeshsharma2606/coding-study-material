// Package slidingwindow contains sliding-window interview problems.
//
// A sliding window maintains a contiguous range [left, right] and slides it over
// the data, adding the element entering on the right and removing the one
// leaving on the left. It turns "check every subarray" (O(n^2)) into O(n) by
// reusing work between overlapping windows. Two shapes: FIXED size, and VARIABLE
// size that grows on the right and shrinks from the left until a constraint holds.
package slidingwindow

// MaxSumSubarrayK returns the maximum sum of any subarray of length k (fixed).
// Slide the window: add nums[right], subtract nums[right-k].
func MaxSumSubarrayK(nums []int, k int) int {
	if len(nums) < k || k == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	best := sum
	for r := k; r < len(nums); r++ {
		sum += nums[r] - nums[r-k]
		if sum > best {
			best = sum
		}
	}
	return best
}

// LengthOfLongestSubstring returns the length of the longest substring without
// repeating characters. Grow right; when a duplicate appears, jump left past
// its previous occurrence. Variable window.
func LengthOfLongestSubstring(s string) int {
	last := make(map[byte]int) // char -> last index seen
	best, left := 0, 0
	for right := 0; right < len(s); right++ {
		if idx, ok := last[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		last[s[right]] = right
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

// MinSubArrayLen returns the minimal length of a contiguous subarray with sum
// >= target (0 if none). Grow right to reach target, then shrink left as long
// as the constraint still holds, recording the min length.
func MinSubArrayLen(target int, nums []int) int {
	left, sum, best := 0, 0, len(nums)+1
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < best {
				best = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if best == len(nums)+1 {
		return 0
	}
	return best
}

// LongestKDistinct returns the length of the longest substring with at most k
// distinct characters. Track a char->count map; shrink when distinct exceeds k.
func LongestKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}
	count := make(map[byte]int)
	left, best := 0, 0
	for right := 0; right < len(s); right++ {
		count[s[right]]++
		for len(count) > k {
			count[s[left]]--
			if count[s[left]] == 0 {
				delete(count, s[left])
			}
			left++
		}
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

// CharacterReplacement returns the length of the longest substring that can be
// made of a single repeated character by replacing at most k others.
// Window valid while (windowLen - maxFreq) <= k. Track running maxFreq.
func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	left := 0
	maxFreq := 0
	result := 0
	for right := 0; right < len(s); right++ {
		count[s[right]]++

		// Maximum frequency character in current window
		if count[s[right]] > maxFreq {
			maxFreq = count[s[right]]
		}

		// Characters that need to be replaced
		windowSize := right - left + 1
		replacements := windowSize - maxFreq

		// If replacements > k, shrink window
		if replacements > k {
			count[s[left]]--
			left++
		}
		windowSize = right - left + 1

		if windowSize > result {
			result = windowSize
		}
	}
	return result
}

// FindAnagrams returns start indices of p's anagrams within s (fixed window).
// Compare frequency arrays as the window of len(p) slides over s.
func FindAnagrams(s, p string) []int {
	var res []int
	if len(s) < len(p) {
		return res
	}
	var need, win [26]int
	for i := 0; i < len(p); i++ {
		need[p[i]-'a']++
		win[s[i]-'a']++
	}
	if win == need {
		res = append(res, 0)
	}
	for r := len(p); r < len(s); r++ {
		win[s[r]-'a']++
		win[s[r-len(p)]-'a']--
		if win == need {
			res = append(res, r-len(p)+1)
		}
	}
	return res
}

// LongestOnes returns the max number of consecutive 1s if you may flip at most
// k zeros. Window valid while zeros in it <= k.
func LongestOnes(nums []int, k int) int {
	left, zeros, best := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

// MinWindow returns the smallest substring of s containing all chars of t
// (with multiplicity), or "" if none. Expand right to satisfy, then contract
// left to minimize while still satisfied.
func MinWindow(s, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	required := len(need) // distinct chars still to satisfy
	formed := 0
	window := make(map[byte]int)
	left, bestLen, bestL := 0, len(s)+1, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++
		if need[c] > 0 && window[c] == need[c] {
			formed++
		}
		for formed == required {
			if right-left+1 < bestLen {
				bestLen = right - left + 1
				bestL = left
			}
			lc := s[left]
			window[lc]--
			if need[lc] > 0 && window[lc] < need[lc] {
				formed--
			}
			left++
		}
	}
	if bestLen == len(s)+1 {
		return ""
	}
	return s[bestL : bestL+bestLen]
}
