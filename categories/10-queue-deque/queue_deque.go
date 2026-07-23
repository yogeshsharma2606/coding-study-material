// Package queuedeque contains queue and deque interview problems.
//
// A queue is FIFO (first in, first out) - the backbone of BFS and stream
// processing. A deque (double-ended queue) allows push/pop at both ends and is
// the key to O(n) sliding-window max/min via a MONOTONIC DEQUE: keep only
// candidates that could still be the answer, discarding dominated ones.
package queuedeque

// MyQueue implements a FIFO queue using two LIFO stacks. Amortized O(1) per op:
// push onto 'in'; when popping, if 'out' is empty, dump 'in' into 'out'
// (reversing order), then pop from 'out'.
type MyQueue struct {
	in, out []int
}

func NewMyQueue() *MyQueue { return &MyQueue{} }

func (q *MyQueue) Push(x int) { q.in = append(q.in, x) }

func (q *MyQueue) transfer() {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			q.out = append(q.out, q.in[len(q.in)-1])
			q.in = q.in[:len(q.in)-1]
		}
	}
}

func (q *MyQueue) Pop() int {
	q.transfer()
	x := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return x
}

func (q *MyQueue) Peek() int {
	q.transfer()
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool { return len(q.in) == 0 && len(q.out) == 0 }

// MovingAverage returns the average of the last 'size' values from a stream.
type MovingAverage struct {
	window []int
	size   int
	sum    int
}

func NewMovingAverage(size int) *MovingAverage { return &MovingAverage{size: size} }

func (m *MovingAverage) Next(val int) float64 {
	m.window = append(m.window, val)
	m.sum += val
	if len(m.window) > m.size {
		m.sum -= m.window[0]
		m.window = m.window[1:]
	}
	return float64(m.sum) / float64(len(m.window))
}

// RecentCounter counts requests in the last 3000 ms. Pings arrive in increasing
// time; evict from the front anything older than t-3000.
type RecentCounter struct {
	q []int
}

func NewRecentCounter() *RecentCounter { return &RecentCounter{} }

func (r *RecentCounter) Ping(t int) int {
	r.q = append(r.q, t)
	for r.q[0] < t-3000 {
		r.q = r.q[1:]
	}
	return len(r.q)
}

// MyCircularQueue is a fixed-capacity ring buffer with O(1) operations.
type MyCircularQueue struct {
	buf        []int
	head, size int
}

func NewMyCircularQueue(k int) *MyCircularQueue {
	return &MyCircularQueue{buf: make([]int, k)}
}

func (c *MyCircularQueue) EnQueue(v int) bool {
	if c.IsFull() {
		return false
	}
	c.buf[(c.head+c.size)%len(c.buf)] = v
	c.size++
	return true
}

func (c *MyCircularQueue) DeQueue() bool {
	if c.IsEmpty() {
		return false
	}
	c.head = (c.head + 1) % len(c.buf)
	c.size--
	return true
}

func (c *MyCircularQueue) Front() int {
	if c.IsEmpty() {
		return -1
	}
	return c.buf[c.head]
}

func (c *MyCircularQueue) Rear() int {
	if c.IsEmpty() {
		return -1
	}
	return c.buf[(c.head+c.size-1)%len(c.buf)]
}

func (c *MyCircularQueue) IsEmpty() bool { return c.size == 0 }
func (c *MyCircularQueue) IsFull() bool  { return c.size == len(c.buf) }

// MaxSlidingWindow returns the maximum of each window of size k.
// A decreasing monotonic deque of INDICES: the front is always the current max;
// pop smaller values from the back (they can never be the max while this one
// lives) and evict indices that fall out of the window from the front.
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}
	var dq []int // indices, nums decreasing
	var res []int
	for i, x := range nums {
		// evict out-of-window index from front
		if len(dq) > 0 && dq[0] <= i-k {
			dq = dq[1:]
		}
		// pop smaller-or-equal from back
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= x {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		if i >= k-1 {
			res = append(res, nums[dq[0]])
		}
	}
	return res
}

// ShortestSubarray returns the length of the shortest subarray with sum >= k
// (values may be negative), or -1. Uses prefix sums + a monotonic-increasing
// deque of prefix indices.
func ShortestSubarray(nums []int, k int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + x
	}
	best := n + 1
	var dq []int // indices into prefix, prefix values increasing
	for i := 0; i <= n; i++ {
		// while current prefix satisfies from the smallest front
		for len(dq) > 0 && prefix[i]-prefix[dq[0]] >= k {
			if i-dq[0] < best {
				best = i - dq[0]
			}
			dq = dq[1:]
		}
		// maintain increasing prefixes: drop larger-or-equal tails
		for len(dq) > 0 && prefix[dq[len(dq)-1]] >= prefix[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}
	if best == n+1 {
		return -1
	}
	return best
}

// MaxResult (Jump Game VI): max score from index 0 to n-1, each jump up to k
// steps. dp[i] = nums[i] + max(dp[i-k..i-1]); a monotonic deque yields the window
// max in O(1) amortized, so overall O(n).
func MaxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	var dq []int // indices, dp decreasing
	dq = append(dq, 0)
	for i := 1; i < n; i++ {
		if dq[0] < i-k {
			dq = dq[1:]
		}
		dp[i] = nums[i] + dp[dq[0]]
		for len(dq) > 0 && dp[dq[len(dq)-1]] <= dp[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}
	return dp[n-1]
}
