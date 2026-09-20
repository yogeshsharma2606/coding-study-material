// Package stack contains stack (LIFO) interview problems.
//
// A stack answers "what was the most recent unmatched thing?" Push to defer work,
// pop to resolve it against the current element. It shines for matching
// (brackets), evaluation (postfix/infix), and undo-like collisions. In Go a slice
// is a perfect stack: append to push, reslice to pop.
package stack

import (
	"strconv"
	"strings"
)

// ValidParentheses reports whether brackets are balanced and correctly nested.
// Push openers; on a closer, the top must be the matching opener.
func ValidParentheses(s string) bool {
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var st []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if open, isClose := pairs[c]; isClose {
			if len(st) == 0 || st[len(st)-1] != open {
				return false
			}
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return len(st) == 0
}

// MinStack supports push/pop/top and retrieving the min, all in O(1).
// A parallel stack stores the min so far at each level.
type MinStack struct {
	data []int
	mins []int
}

func NewMinStack() *MinStack { return &MinStack{} }

func (m *MinStack) Push(x int) {
	m.data = append(m.data, x)
	if len(m.mins) == 0 || x < m.mins[len(m.mins)-1] {
		m.mins = append(m.mins, x)
	} else {
		m.mins = append(m.mins, m.mins[len(m.mins)-1])
	}
}

func (m *MinStack) Pop() {
	m.data = m.data[:len(m.data)-1]
	m.mins = m.mins[:len(m.mins)-1]
}

func (m *MinStack) Top() int { return m.data[len(m.data)-1] }
func (m *MinStack) Min() int { return m.mins[len(m.mins)-1] }

// EvalRPN evaluates Reverse Polish Notation. Operands push; an operator pops two
// and pushes the result. Postfix needs no precedence handling - that's its point.
func EvalRPN(tokens []string) int {
	var st []int
	pop := func() int {
		x := st[len(st)-1]
		st = st[:len(st)-1]
		return x
	}
	for _, tok := range tokens {
		switch tok {
		case "+", "-", "*", "/":
			b, a := pop(), pop()
			switch tok {
			case "+":
				st = append(st, a+b)
			case "-":
				st = append(st, a-b)
			case "*":
				st = append(st, a*b)
			case "/":
				st = append(st, a/b)
			}
		default:
			n, _ := strconv.Atoi(tok)
			st = append(st, n)
		}
	}
	return st[0]
}

// SimplifyPath canonicalizes a Unix path. Split on '/'; ".." pops, "." and ""
// are skipped, names push. Rebuild from the stack.
func SimplifyPath(path string) string {
	var st []string
	for _, part := range strings.Split(path, "/") {
		switch part {
		case "", ".":
			// skip
		case "..":
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		default:
			st = append(st, part)
		}
	}
	return "/" + strings.Join(st, "/")
}

// DecodeString expands strings like "3[a2[c]]" -> "accaccacc".
// Two stacks: counts and the partial string built before each '['. On ']',
// pop and repeat.
func DecodeString(s string) string {
	var counts []int
	var strs []string
	cur := ""
	num := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			num = num*10 + int(c-'0')
		case c == '[':
			counts = append(counts, num)
			strs = append(strs, cur)
			num, cur = 0, ""
		case c == ']':
			k := counts[len(counts)-1]
			counts = counts[:len(counts)-1]
			prev := strs[len(strs)-1]
			strs = strs[:len(strs)-1]
			cur = prev + strings.Repeat(cur, k)
		default:
			cur += string(c)
		}
	}
	return cur
}

// BackspaceCompare compares two strings where '#' is a backspace.
// Build each result on a stack; '#' pops.
func buildString(s string) []byte {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return stack
}

func backspaceCompare(s string, t string) bool {
	sStack := buildString(s)
	tStack := buildString(t)
	if len(sStack) != len(tStack) {
		return false
	}
	for i := range sStack {
		if sStack[i] != tStack[i] {
			return false
		}
	}
	return true
}

// AsteroidCollision resolves collisions. Positive move right, negative left.
// A right-mover on the stack may collide with an incoming left-mover; smaller
// explodes. Only right-then-left pairs collide.
func AsteroidCollision(asteroids []int) []int {
	var st []int
	for _, a := range asteroids {
		alive := true
		for alive && a < 0 && len(st) > 0 && st[len(st)-1] > 0 {
			top := st[len(st)-1]
			if top < -a {
				st = st[:len(st)-1] // top explodes, keep checking
			} else if top == -a {
				st = st[:len(st)-1] // both explode
				alive = false
			} else {
				alive = false // incoming explodes
			}
		}
		if alive {
			st = append(st, a)
		}
	}
	return st
}

// Calculate evaluates a basic expression with + - ( ) and non-negative ints.
// Keep a running result, a sign, and use a stack to save state at '('.
func Calculate(s string) int {
	var st []int // saved (result, sign) pairs, flattened
	result, sign, num := 0, 1, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			num = num*10 + int(c-'0')
		case c == '+' || c == '-':
			result += sign * num
			num = 0
			if c == '+' {
				sign = 1
			} else {
				sign = -1
			}
		case c == '(':
			st = append(st, result, sign)
			result, sign = 0, 1
		case c == ')':
			result += sign * num
			num = 0
			prevSign := st[len(st)-1]
			prevResult := st[len(st)-2]
			st = st[:len(st)-2]
			result = prevResult + prevSign*result
		}
	}
	return result + sign*num
}
