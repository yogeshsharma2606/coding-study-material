package stack

import (
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, tt := range tests {
		if got := ValidParentheses(tt.in); got != tt.want {
			t.Errorf("ValidParentheses(%q)=%v want %v", tt.in, got, tt.want)
		}
	}
}

func TestMinStack(t *testing.T) {
	m := NewMinStack()
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	if m.Min() != -3 {
		t.Errorf("min got %d", m.Min())
	}
	m.Pop()
	if m.Top() != 0 {
		t.Errorf("top got %d", m.Top())
	}
	if m.Min() != -2 {
		t.Errorf("min got %d", m.Min())
	}
}

func TestEvalRPN(t *testing.T) {
	if got := EvalRPN([]string{"2", "1", "+", "3", "*"}); got != 9 {
		t.Errorf("got %d", got)
	}
	if got := EvalRPN([]string{"4", "13", "5", "/", "+"}); got != 6 {
		t.Errorf("got %d", got)
	}
}

func TestSimplifyPath(t *testing.T) {
	tests := []struct{ in, want string }{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
	}
	for _, tt := range tests {
		if got := SimplifyPath(tt.in); got != tt.want {
			t.Errorf("SimplifyPath(%q)=%q want %q", tt.in, got, tt.want)
		}
	}
}

func TestDecodeString(t *testing.T) {
	tests := []struct{ in, want string }{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, tt := range tests {
		if got := DecodeString(tt.in); got != tt.want {
			t.Errorf("DecodeString(%q)=%q want %q", tt.in, got, tt.want)
		}
	}
}

func TestBackspaceCompare(t *testing.T) {
	if !BackspaceCompare("ab#c", "ad#c") {
		t.Error("expected equal")
	}
	if BackspaceCompare("a#c", "b") {
		t.Error("expected not equal")
	}
}

func TestAsteroidCollision(t *testing.T) {
	if got := AsteroidCollision([]int{5, 10, -5}); !reflect.DeepEqual(got, []int{5, 10}) {
		t.Errorf("got %v", got)
	}
	if got := AsteroidCollision([]int{8, -8}); len(got) != 0 {
		t.Errorf("got %v", got)
	}
	if got := AsteroidCollision([]int{10, 2, -5}); !reflect.DeepEqual(got, []int{10}) {
		t.Errorf("got %v", got)
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
	}
	for _, tt := range tests {
		if got := Calculate(tt.in); got != tt.want {
			t.Errorf("Calculate(%q)=%d want %d", tt.in, got, tt.want)
		}
	}
}
