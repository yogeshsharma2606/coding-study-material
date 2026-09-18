package main

import "fmt"

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

func main() {
	s := "AABABBA"
	k := 1

	fmt.Println(characterReplacement(s, k))
}
