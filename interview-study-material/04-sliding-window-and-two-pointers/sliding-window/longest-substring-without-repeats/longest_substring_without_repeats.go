package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// characterMap stores the last seen index of each character
	characterMap := make(map[byte]int)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		// If the character is in the map and within our current window
		if lastIndex, found := characterMap[char]; found && lastIndex >= left {
			// Move left pointer to the right of the previous occurrence
			left = lastIndex + 1
		}

		// Update the character's last seen position
		characterMap[char] = right

		// Calculate the current window size
		currentWindowSize := right - left + 1
		if currentWindowSize > maxLen {
			maxLen = currentWindowSize
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabd")) //	4 (cabd)
	fmt.Println(lengthOfLongestSubstring("bbbbb"))  // 1 (b)
	fmt.Println(lengthOfLongestSubstring("pwwkew")) // 3 (wke)
}
