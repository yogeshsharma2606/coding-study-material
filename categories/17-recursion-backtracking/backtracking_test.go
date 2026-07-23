package backtracking

import (
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	got := Subsets([]int{1, 2, 3})
	if len(got) != 8 { // 2^3
		t.Errorf("got %d subsets", len(got))
	}
}

func TestPermutations(t *testing.T) {
	got := Permutations([]int{1, 2, 3})
	if len(got) != 6 { // 3!
		t.Errorf("got %d permutations", len(got))
	}
}

func TestCombine(t *testing.T) {
	got := Combine(4, 2)
	if len(got) != 6 { // C(4,2)
		t.Errorf("got %d", len(got))
	}
}

func TestCombinationSum(t *testing.T) {
	got := CombinationSum([]int{2, 3, 6, 7}, 7)
	// expect [2,2,3] and [7]
	if len(got) != 2 {
		t.Errorf("got %v", got)
	}
	for _, c := range got {
		sum := 0
		for _, x := range c {
			sum += x
		}
		if sum != 7 {
			t.Errorf("bad combo %v", c)
		}
	}
}

func TestGenerateParenthesis(t *testing.T) {
	got := GenerateParenthesis(3)
	if len(got) != 5 { // Catalan(3)
		t.Errorf("got %d", len(got))
	}
}

func TestLetterCombinations(t *testing.T) {
	got := LetterCombinations("23")
	if len(got) != 9 { // 3*3
		t.Errorf("got %d", len(got))
	}
	if LetterCombinations("") != nil {
		t.Error("empty input")
	}
}

func TestPalindromePartition(t *testing.T) {
	got := PalindromePartition("aab")
	// [["a","a","b"],["aa","b"]]
	if len(got) != 2 {
		t.Errorf("got %v", got)
	}
}

func TestExist(t *testing.T) {
	board := [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}
	if !Exist(board, "ABCCED") {
		t.Error("ABCCED should exist")
	}
	if !Exist(board, "SEE") {
		t.Error("SEE should exist")
	}
	if Exist(board, "ABCB") {
		t.Error("ABCB should not exist")
	}
}

func TestSolveNQueens(t *testing.T) {
	got := SolveNQueens(4)
	if len(got) != 2 {
		t.Errorf("got %d solutions", len(got))
	}
	if len(SolveNQueens(1)) != 1 {
		t.Error("n=1 should have 1")
	}
	// n=8 has 92 solutions
	sols := SolveNQueens(8)
	sort.Slice(sols, func(i, j int) bool { return len(sols[i]) < len(sols[j]) })
	if len(sols) != 92 {
		t.Errorf("n=8 got %d", len(sols))
	}
}
