package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compress(s string) string {
	if len(s) == 0 {
		return s
	}

	var b strings.Builder
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			b.WriteByte(s[i-1])
			b.WriteString(strconv.Itoa(count))
			count = 1
		}
	}
	b.WriteByte(s[len(s)-1])
	b.WriteString(strconv.Itoa(count))

	return b.String()
}

func main() {
	fmt.Println(compress("aabcccccaaa"))  // "a2b1c5a3"
	fmt.Println(compress("aaabbbcccaaa")) // "a3b3c3a3"
	fmt.Println(compress("a"))            // "a1"
	fmt.Println(compress(""))             // ""
}
