// Package heapx contains heap / priority-queue interview problems.
//
// A binary heap gives O(log n) push/pop and O(1) peek at the extreme (min or
// max). Reach for it whenever you need the "top" element repeatedly, a running
// median, or a k-way merge. Key trick: to keep the k LARGEST, use a MIN-heap of
// size k (pop the smallest when it overflows) - and vice versa. In Go, implement
// heap.Interface; here we wrap it in a generic priority queue.
package heapx

import (
	"container/heap"
	"sort"
)

// PQ is a generic priority queue over any type, ordered by the provided less
// function (less(a,b)==true means a has higher priority / comes out first).
type PQ[T any] struct {
	items []T
	less  func(a, b T) bool
}

func NewPQ[T any](less func(a, b T) bool) *PQ[T] { return &PQ[T]{less: less} }

func (p PQ[T]) Len() int           { return len(p.items) }
func (p PQ[T]) Less(i, j int) bool { return p.less(p.items[i], p.items[j]) }
func (p PQ[T]) Swap(i, j int)      { p.items[i], p.items[j] = p.items[j], p.items[i] }
func (p *PQ[T]) Push(x any)        { p.items = append(p.items, x.(T)) }
func (p *PQ[T]) Pop() any {
	n := len(p.items)
	x := p.items[n-1]
	p.items = p.items[:n-1]
	return x
}
func (p *PQ[T]) push(x T) { heap.Push(p, x) }
func (p *PQ[T]) pop() T   { return heap.Pop(p).(T) }
func (p *PQ[T]) peek() T  { return p.items[0] }

// FindKthLargest returns the kth largest element using a MIN-heap of size k.
// Keep only the k largest seen so far; the heap's root is the kth largest.
func FindKthLargest(nums []int, k int) int {
	h := NewPQ(func(a, b int) bool { return a < b }) // min-heap
	for _, x := range nums {
		h.push(x)
		if h.Len() > k {
			h.pop() // drop the smallest
		}
	}
	return h.peek()
}

// KthLargest maintains the kth largest in a stream via a size-k min-heap.
type KthLargest struct {
	h *PQ[int]
	k int
}

func NewKthLargest(k int, nums []int) *KthLargest {
	kl := &KthLargest{h: NewPQ(func(a, b int) bool { return a < b }), k: k}
	for _, x := range nums {
		kl.Add(x)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	kl.h.push(val)
	if kl.h.Len() > kl.k {
		kl.h.pop()
	}
	return kl.h.peek()
}

// LastStoneWeight repeatedly smashes the two heaviest stones (max-heap).
func LastStoneWeight(stones []int) int {
	h := NewPQ(func(a, b int) bool { return a > b }) // max-heap
	for _, s := range stones {
		h.push(s)
	}
	for h.Len() > 1 {
		a, b := h.pop(), h.pop()
		if a != b {
			h.push(a - b)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h.peek()
}

// KClosest returns the k points closest to the origin using a MAX-heap of size k
// keyed by squared distance (drop the farthest when it overflows).
func KClosest(points [][]int, k int) [][]int {
	dist := func(p []int) int { return p[0]*p[0] + p[1]*p[1] }
	h := NewPQ(func(a, b []int) bool { return dist(a) > dist(b) }) // max-heap
	for _, p := range points {
		h.push(p)
		if h.Len() > k {
			h.pop()
		}
	}
	return h.items
}

// ListNode for the k-way merge problem.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeKLists merges k sorted lists using a min-heap of the current heads.
// Pop the smallest, append it, push its successor. O(N log k).
func MergeKLists(lists []*ListNode) *ListNode {
	h := NewPQ(func(a, b *ListNode) bool { return a.Val < b.Val })
	for _, l := range lists {
		if l != nil {
			h.push(l)
		}
	}
	dummy := &ListNode{}
	tail := dummy
	for h.Len() > 0 {
		node := h.pop()
		tail.Next = node
		tail = tail.Next
		if node.Next != nil {
			h.push(node.Next)
		}
	}
	return dummy.Next
}

// MedianFinder keeps a running median with two heaps: a max-heap of the lower
// half and a min-heap of the upper half, balanced so their roots straddle the
// median.
type MedianFinder struct {
	lo *PQ[int] // max-heap (lower half)
	hi *PQ[int] // min-heap (upper half)
}

func NewMedianFinder() *MedianFinder {
	return &MedianFinder{
		lo: NewPQ(func(a, b int) bool { return a > b }),
		hi: NewPQ(func(a, b int) bool { return a < b }),
	}
}

func (m *MedianFinder) AddNum(num int) {
	m.lo.push(num)
	m.hi.push(m.lo.pop()) // move lo's max to hi to keep order
	if m.hi.Len() > m.lo.Len() {
		m.lo.push(m.hi.pop()) // rebalance sizes
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.lo.Len() > m.hi.Len() {
		return float64(m.lo.peek())
	}
	return float64(m.lo.peek()+m.hi.peek()) / 2
}

// TopKFrequent returns the k most frequent elements using a size-k min-heap
// keyed by frequency.
func TopKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, x := range nums {
		freq[x]++
	}
	type pair struct{ val, cnt int }
	h := NewPQ(func(a, b pair) bool { return a.cnt < b.cnt }) // min-heap by count
	for v, c := range freq {
		h.push(pair{v, c})
		if h.Len() > k {
			h.pop()
		}
	}
	out := make([]int, 0, k)
	for h.Len() > 0 {
		out = append(out, h.pop().val)
	}
	sort.Ints(out)
	return out
}

// ReorganizeString rearranges s so no two adjacent chars are equal, or "".
// Greedily place the most frequent available char that differs from the last
// one, using a max-heap by remaining count.
func ReorganizeString(s string) string {
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	type pair struct {
		ch  byte
		cnt int
	}
	h := NewPQ(func(a, b pair) bool { return a.cnt > b.cnt }) // max-heap
	for i, c := range count {
		if c > 0 {
			h.push(pair{byte('a' + i), c})
		}
	}
	var res []byte
	var prev *pair
	for h.Len() > 0 {
		cur := h.pop()
		res = append(res, cur.ch)
		cur.cnt--
		if prev != nil && prev.cnt > 0 {
			h.push(*prev)
		}
		prev = &pair{cur.ch, cur.cnt}
	}
	if len(res) != len(s) {
		return ""
	}
	return string(res)
}
