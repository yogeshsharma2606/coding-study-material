// Package strs contains classic string interview problems.
//
// A string in Go is an immutable read-only slice of bytes (UTF-8). For ASCII
// interview problems, index into bytes; for Unicode, range over runes. Most
// string problems reduce to: frequency counting (fixed-size arrays beat maps
// for a-z), two pointers (palindromes), or a hash key that groups equal things.
package strs

import (
	"sort"
	"strings"
	"unicode"
)

// ReverseString reverses a byte slice in place with two converging pointers.
func ReverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

// IsAnagram reports whether t is an anagram of s.
// Count each letter (+1 for s, -1 for t); all counts must end at zero.
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

// IsPalindrome checks alphanumeric-only, case-insensitive palindrome.
// Two pointers converge, skipping non-alphanumeric characters.
func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlnum(rune(s[i])) {
			i++
		}
		for i < j && !isAlnum(rune(s[j])) {
			j--
		}
		if lower(s[i]) != lower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlnum(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) }

func lower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

// GroupAnagrams groups words that are anagrams of each other.
// Key = the sorted letters of the word; anagrams share the same key.
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, w := range strs {
		b := []byte(w)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		groups[key] = append(groups[key], w)
	}
	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		out = append(out, g)
	}
	return out
}

// LongestCommonPrefix returns the longest common prefix of all strings.
// Vertical scan: compare column by column across all words.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for _, w := range strs[1:] {
			if i >= len(w) || w[i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// IsIsomorphic reports whether characters in s can be one-to-one mapped to t.
// Maintain two maps (s->t and t->s) so the mapping is a bijection.
func IsIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	st := make(map[byte]byte)
	ts := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if v, ok := st[a]; ok && v != b {
			return false
		}
		if v, ok := ts[b]; ok && v != a {
			return false
		}
		st[a], ts[b] = b, a
	}
	return true
}

// MyAtoi converts a string to a 32-bit signed integer (LeetCode rules):
// skip leading spaces, optional sign, read digits, clamp to int32 range.
func MyAtoi(s string) int {
	const intMax, intMin = 1<<31 - 1, -(1 << 31)
	i, n := 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	num := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		num = num*10 + int(s[i]-'0')
		if sign == 1 && num > intMax {
			return intMax
		}
		if sign == -1 && -num < intMin {
			return intMin
		}
		i++
	}
	return sign * num
}

// LongestPalindrome returns the longest palindromic substring by expanding
// around each possible center (2n-1 centers: single chars and gaps).
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, maxLen := 0, 1
	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxLen {
				start, maxLen = l, r-l+1
			}
			l--
			r++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)   // odd length center
		expand(i, i+1) // even length center
	}
	return s[start : start+maxLen]
}

// Encode/Decode serialize a list of strings to one string and back using a
// length prefix ("<len>#<payload>"), which is safe for any characters.
func Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(itoa(len(s)))
		sb.WriteByte('#')
		sb.WriteString(s)
	}
	return sb.String()
}

func Decode(s string) []string {
	var out []string
	i := 0
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		length := atoiSimple(s[i:j])
		start := j + 1
		out = append(out, s[start:start+length])
		i = start + length
	}
	return out
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}

func atoiSimple(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
