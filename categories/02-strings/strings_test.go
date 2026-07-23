package strs

import (
	"reflect"
	"sort"
	"testing"
)

func TestReverseString(t *testing.T) {
	b := []byte("hello")
	ReverseString(b)
	if string(b) != "olleh" {
		t.Errorf("got %s", b)
	}
}

func TestIsAnagram(t *testing.T) {
	if !IsAnagram("anagram", "nagaram") {
		t.Error("expected anagram")
	}
	if IsAnagram("rat", "car") {
		t.Error("expected not anagram")
	}
}

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("A man, a plan, a canal: Panama") {
		t.Error("expected palindrome")
	}
	if IsPalindrome("race a car") {
		t.Error("expected not palindrome")
	}
}

func TestGroupAnagrams(t *testing.T) {
	got := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	// normalize for comparison
	for _, g := range got {
		sort.Strings(g)
	}
	sort.Slice(got, func(i, j int) bool { return got[i][0] < got[j][0] })
	want := [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	if got := LongestCommonPrefix([]string{"flower", "flow", "flight"}); got != "fl" {
		t.Errorf("got %q", got)
	}
	if got := LongestCommonPrefix([]string{"dog", "racecar", "car"}); got != "" {
		t.Errorf("got %q", got)
	}
}

func TestIsIsomorphic(t *testing.T) {
	if !IsIsomorphic("egg", "add") {
		t.Error("egg/add should be isomorphic")
	}
	if IsIsomorphic("foo", "bar") {
		t.Error("foo/bar should not be isomorphic")
	}
	if IsIsomorphic("badc", "baba") {
		t.Error("badc/baba should not be isomorphic")
	}
}

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
	}
	for _, tt := range tests {
		if got := MyAtoi(tt.in); got != tt.want {
			t.Errorf("MyAtoi(%q)=%d want %d", tt.in, got, tt.want)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	got := LongestPalindrome("babad")
	if got != "bab" && got != "aba" {
		t.Errorf("got %q", got)
	}
	if got := LongestPalindrome("cbbd"); got != "bb" {
		t.Errorf("got %q", got)
	}
}

func TestEncodeDecode(t *testing.T) {
	in := []string{"hello", "wor#ld", "", "42#"}
	if got := Decode(Encode(in)); !reflect.DeepEqual(got, in) {
		t.Errorf("roundtrip got %v want %v", got, in)
	}
}
