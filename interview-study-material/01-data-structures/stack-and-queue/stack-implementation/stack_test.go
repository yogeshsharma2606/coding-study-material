package main

import "testing"

// Helper function to create a stack with initial values
func setupStack(values ...int) *Stack {
	s := &Stack{}
	for _, v := range values {
		s.Push(v)
	}
	return s
}

// Helper function to assert stack length
func assertStackLength(t *testing.T, s *Stack, expected int) {
	// t.Helper()
	if got := len(s.items); got != expected {
		t.Errorf("stack length: got %d, want %d", got, expected)
	}
}

// Helper function to assert values
func assertEqual(t *testing.T, got, want int) {
	// t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestStack_Push(t *testing.T) {
	tests := []struct {
		name           string
		pushValues     []int
		expectedLength int
		expectedTop    int
	}{
		{
			name:           "push single element",
			pushValues:     []int{10},
			expectedLength: 1,
			expectedTop:    10,
		},
		{
			name:           "push multiple elements",
			pushValues:     []int{10, 20, 30},
			expectedLength: 3,
			expectedTop:    30,
		},
		{
			name:           "push negative values",
			pushValues:     []int{-10, -20, -30},
			expectedLength: 3,
			expectedTop:    -30,
		},
		{
			name:           "push zero",
			pushValues:     []int{0},
			expectedLength: 1,
			expectedTop:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setupStack(tt.pushValues...)

			assertStackLength(t, s, tt.expectedLength)
			assertEqual(t, s.Peek(), tt.expectedTop)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name         string
		initialStack []int
		popCount     int
		wantValues   []int
		finalLength  int
	}{
		{
			name:         "pop from empty stack",
			initialStack: []int{},
			popCount:     1,
			wantValues:   []int{-1},
			finalLength:  0,
		},
		{
			name:         "pop single element",
			initialStack: []int{42},
			popCount:     1,
			wantValues:   []int{42},
			finalLength:  0,
		},
		{
			name:         "pop multiple elements",
			initialStack: []int{10, 20, 30},
			popCount:     3,
			wantValues:   []int{30, 20, 10},
			finalLength:  0,
		},
		{
			name:         "pop partial elements",
			initialStack: []int{10, 20, 30, 40},
			popCount:     2,
			wantValues:   []int{40, 30},
			finalLength:  2,
		},
		{
			name:         "pop with negative values",
			initialStack: []int{-10, -20, -30},
			popCount:     2,
			wantValues:   []int{-30, -20},
			finalLength:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setupStack(tt.initialStack...)

			for i := 0; i < tt.popCount; i++ {
				got := s.Pop()
				assertEqual(t, got, tt.wantValues[i])
			}

			assertStackLength(t, s, tt.finalLength)
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name         string
		initialStack []int
		wantValue    int
		wantLength   int
	}{
		{
			name:         "peek empty stack",
			initialStack: []int{},
			wantValue:    -1,
			wantLength:   0,
		},
		{
			name:         "peek single element",
			initialStack: []int{42},
			wantValue:    42,
			wantLength:   1,
		},
		{
			name:         "peek multiple elements",
			initialStack: []int{10, 20, 30},
			wantValue:    30,
			wantLength:   3,
		},
		{
			name:         "peek negative value",
			initialStack: []int{-10, -20},
			wantValue:    -20,
			wantLength:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setupStack(tt.initialStack...)

			got := s.Peek()
			assertEqual(t, got, tt.wantValue)
			assertStackLength(t, s, tt.wantLength)
		})
	}
}

func TestStack_Operations(t *testing.T) {
	t.Run("multiple peeks should not modify stack", func(t *testing.T) {
		s := setupStack(10, 20, 30)

		for i := 0; i < 5; i++ {
			assertEqual(t, s.Peek(), 30)
		}
		assertStackLength(t, s, 3)
	})

	t.Run("push after pop maintains LIFO order", func(t *testing.T) {
		s := setupStack(10, 20)
		assertEqual(t, s.Pop(), 20)

		s.Push(30)
		assertEqual(t, s.Peek(), 30)
		assertEqual(t, s.Pop(), 30)
		assertEqual(t, s.Pop(), 10)
	})

	t.Run("empty stack after popping all elements", func(t *testing.T) {
		s := setupStack(10, 20, 30)

		s.Pop()
		s.Pop()
		s.Pop()

		assertEqual(t, s.Pop(), -1)
		assertEqual(t, s.Peek(), -1)
		assertStackLength(t, s, 0)
	})

	t.Run("mixed operations sequence", func(t *testing.T) {
		s := &Stack{}

		s.Push(5)
		assertEqual(t, s.Peek(), 5)

		s.Push(10)
		s.Push(15)
		assertEqual(t, s.Pop(), 15)
		assertEqual(t, s.Peek(), 10)

		s.Push(20)
		assertEqual(t, s.Pop(), 20)
		assertEqual(t, s.Pop(), 10)
		assertEqual(t, s.Pop(), 5)
		assertEqual(t, s.Pop(), -1)
	})
}

func TestStack_EdgeCases(t *testing.T) {
	t.Run("multiple operations on empty stack", func(t *testing.T) {
		s := &Stack{}

		for i := 0; i < 5; i++ {
			assertEqual(t, s.Pop(), -1)
			assertEqual(t, s.Peek(), -1)
		}
	})

	t.Run("large number of elements", func(t *testing.T) {
		s := &Stack{}
		size := 1000

		// Push elements
		for i := 0; i < size; i++ {
			s.Push(i)
		}

		assertStackLength(t, s, size)
		assertEqual(t, s.Peek(), size-1)

		// Pop all elements
		for i := size - 1; i >= 0; i-- {
			assertEqual(t, s.Pop(), i)
		}

		assertStackLength(t, s, 0)
	})

	t.Run("alternating push and pop", func(t *testing.T) {
		s := &Stack{}

		for i := 0; i < 100; i++ {
			s.Push(i)
			if i%2 == 1 {
				s.Pop()
			}
		}

		// Should have 50 elements remaining
		assertStackLength(t, s, 50)
	})

	t.Run("single element lifecycle", func(t *testing.T) {
		s := &Stack{}

		s.Push(42)
		assertEqual(t, s.Peek(), 42)
		assertEqual(t, s.Pop(), 42)
		assertEqual(t, s.Peek(), -1)
		assertStackLength(t, s, 0)
	})
}
