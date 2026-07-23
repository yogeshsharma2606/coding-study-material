package hashing

import (
	"sort"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	if !ContainsDuplicate([]int{1, 2, 3, 1}) {
		t.Error("expected true")
	}
	if ContainsDuplicate([]int{1, 2, 3, 4}) {
		t.Error("expected false")
	}
}

func TestFirstUniqChar(t *testing.T) {
	if got := FirstUniqChar("leetcode"); got != 0 {
		t.Errorf("got %d", got)
	}
	if got := FirstUniqChar("loveleetcode"); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := FirstUniqChar("aabb"); got != -1 {
		t.Errorf("got %d", got)
	}
}

func TestIntersection(t *testing.T) {
	got := Intersection([]int{1, 2, 2, 1}, []int{2, 2})
	if len(got) != 1 || got[0] != 2 {
		t.Errorf("got %v", got)
	}
}

func TestTopKFrequent(t *testing.T) {
	got := TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	sort.Ints(got)
	if len(got) != 2 || got[0] != 1 || got[1] != 2 {
		t.Errorf("got %v", got)
	}
}

func TestLongestConsecutive(t *testing.T) {
	if got := LongestConsecutive([]int{100, 4, 200, 1, 3, 2}); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := LongestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}); got != 9 {
		t.Errorf("got %d", got)
	}
}

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	if !IsValidSudoku(board) {
		t.Error("expected valid")
	}
	board[0][0] = '8' // duplicate 8 in first column / box
	if IsValidSudoku(board) {
		t.Error("expected invalid")
	}
}

func TestFourSumCount(t *testing.T) {
	got := FourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2})
	if got != 2 {
		t.Errorf("got %d", got)
	}
}

func TestMajorityElement(t *testing.T) {
	if got := MajorityElement([]int{2, 2, 1, 1, 1, 2, 2}); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := MajorityElement([]int{3, 2, 3}); got != 3 {
		t.Errorf("got %d", got)
	}
}
