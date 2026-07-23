package queuedeque

import (
	"reflect"
	"testing"
)

func TestMyQueue(t *testing.T) {
	q := NewMyQueue()
	q.Push(1)
	q.Push(2)
	if q.Peek() != 1 {
		t.Errorf("peek got %d", q.Peek())
	}
	if q.Pop() != 1 {
		t.Error("pop 1")
	}
	q.Push(3)
	if q.Pop() != 2 {
		t.Error("pop 2")
	}
	if q.Pop() != 3 {
		t.Error("pop 3")
	}
	if !q.Empty() {
		t.Error("should be empty")
	}
}

func TestMovingAverage(t *testing.T) {
	m := NewMovingAverage(3)
	if got := m.Next(1); got != 1.0 {
		t.Errorf("got %v", got)
	}
	if got := m.Next(10); got != 5.5 {
		t.Errorf("got %v", got)
	}
	m.Next(3)
	if got := m.Next(5); got != 6.0 {
		t.Errorf("got %v", got)
	}
}

func TestRecentCounter(t *testing.T) {
	r := NewRecentCounter()
	if r.Ping(1) != 1 {
		t.Error("1")
	}
	if r.Ping(100) != 2 {
		t.Error("2")
	}
	if r.Ping(3001) != 3 {
		t.Error("3")
	}
	if r.Ping(3002) != 3 {
		t.Error("3 again")
	}
}

func TestMyCircularQueue(t *testing.T) {
	c := NewMyCircularQueue(3)
	if !c.EnQueue(1) || !c.EnQueue(2) || !c.EnQueue(3) {
		t.Error("enqueue")
	}
	if c.EnQueue(4) {
		t.Error("should be full")
	}
	if c.Rear() != 3 {
		t.Error("rear")
	}
	if !c.IsFull() {
		t.Error("full")
	}
	if !c.DeQueue() {
		t.Error("dequeue")
	}
	if !c.EnQueue(4) {
		t.Error("enqueue after dequeue")
	}
	if c.Rear() != 4 {
		t.Error("rear 4")
	}
	if c.Front() != 2 {
		t.Error("front 2")
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	got := MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	want := []int{3, 3, 5, 5, 6, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestShortestSubarray(t *testing.T) {
	if got := ShortestSubarray([]int{1}, 1); got != 1 {
		t.Errorf("got %d", got)
	}
	if got := ShortestSubarray([]int{1, 2}, 4); got != -1 {
		t.Errorf("got %d", got)
	}
	if got := ShortestSubarray([]int{2, -1, 2}, 3); got != 3 {
		t.Errorf("got %d", got)
	}
}

func TestMaxResult(t *testing.T) {
	if got := MaxResult([]int{1, -1, -2, 4, -7, 3}, 2); got != 7 {
		t.Errorf("got %d", got)
	}
	if got := MaxResult([]int{10, -5, -2, 4, 0, 3}, 3); got != 17 {
		t.Errorf("got %d", got)
	}
}
