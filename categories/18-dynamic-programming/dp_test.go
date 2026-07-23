package dp

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct{ n, want int }{{2, 2}, {3, 3}, {5, 8}}
	for _, tt := range tests {
		if got := ClimbStairs(tt.n); got != tt.want {
			t.Errorf("ClimbStairs(%d)=%d want %d", tt.n, got, tt.want)
		}
	}
}

func TestRob(t *testing.T) {
	if got := Rob([]int{1, 2, 3, 1}); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := Rob([]int{2, 7, 9, 3, 1}); got != 12 {
		t.Errorf("got %d", got)
	}
}

func TestCoinChange(t *testing.T) {
	if got := CoinChange([]int{1, 2, 5}, 11); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := CoinChange([]int{2}, 3); got != -1 {
		t.Errorf("got %d", got)
	}
	if got := CoinChange([]int{1}, 0); got != 0 {
		t.Errorf("got %d", got)
	}
}

func TestLengthOfLIS(t *testing.T) {
	if got := LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := LengthOfLIS([]int{0, 1, 0, 3, 2, 3}); got != 4 {
		t.Errorf("got %d", got)
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	if got := LongestCommonSubsequence("abcde", "ace"); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := LongestCommonSubsequence("abc", "def"); got != 0 {
		t.Errorf("got %d", got)
	}
}

func TestKnapsack01(t *testing.T) {
	got := Knapsack01([]int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7)
	if got != 9 { // items with weight 3+4 -> value 4+5
		t.Errorf("got %d", got)
	}
}

func TestUniquePaths(t *testing.T) {
	if got := UniquePaths(3, 7); got != 28 {
		t.Errorf("got %d", got)
	}
	if got := UniquePaths(3, 2); got != 3 {
		t.Errorf("got %d", got)
	}
}

func TestWordBreak(t *testing.T) {
	if !WordBreak("leetcode", []string{"leet", "code"}) {
		t.Error("leetcode should break")
	}
	if !WordBreak("applepenapple", []string{"apple", "pen"}) {
		t.Error("applepenapple should break")
	}
	if WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) {
		t.Error("catsandog should not break")
	}
}

func TestMinDistance(t *testing.T) {
	if got := MinDistance("horse", "ros"); got != 3 {
		t.Errorf("got %d", got)
	}
	if got := MinDistance("intention", "execution"); got != 5 {
		t.Errorf("got %d", got)
	}
}

func TestCanPartition(t *testing.T) {
	if !CanPartition([]int{1, 5, 11, 5}) {
		t.Error("should partition")
	}
	if CanPartition([]int{1, 2, 3, 5}) {
		t.Error("should not partition")
	}
}
