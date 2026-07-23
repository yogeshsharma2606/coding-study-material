package slidingwindow

import (
	"reflect"
	"testing"
)

func TestMaxSumSubarrayK(t *testing.T) {
	if got := MaxSumSubarrayK([]int{2, 1, 5, 1, 3, 2}, 3); got != 9 {
		t.Errorf("got %d", got)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}
	for _, tt := range tests {
		if got := LengthOfLongestSubstring(tt.in); got != tt.want {
			t.Errorf("LengthOfLongestSubstring(%q)=%d want %d", tt.in, got, tt.want)
		}
	}
}

func TestMinSubArrayLen(t *testing.T) {
	if got := MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := MinSubArrayLen(11, []int{1, 1, 1, 1}); got != 0 {
		t.Errorf("got %d", got)
	}
}

func TestLongestKDistinct(t *testing.T) {
	if got := LongestKDistinct("eceba", 2); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := LongestKDistinct("aa", 1); got != 2 {
		t.Errorf("got %d", got)
	}
}

func TestCharacterReplacement(t *testing.T) {
	if got := CharacterReplacement("AABABBA", 1); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := CharacterReplacement("ABAB", 2); got != 4 {
		t.Errorf("got %d", got)
	}
}

func TestFindAnagrams(t *testing.T) {
	if got := FindAnagrams("cbaebabacd", "abc"); !reflect.DeepEqual(got, []int{0, 6}) {
		t.Errorf("got %v", got)
	}
}

func TestLongestOnes(t *testing.T) {
	if got := LongestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2); got != 6 {
		t.Errorf("got %d", got)
	}
}

func TestMinWindow(t *testing.T) {
	if got := MinWindow("ADOBECODEBANC", "ABC"); got != "BANC" {
		t.Errorf("got %q", got)
	}
	if got := MinWindow("a", "aa"); got != "" {
		t.Errorf("got %q", got)
	}
}
