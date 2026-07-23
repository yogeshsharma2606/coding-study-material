package graphs

import (
	"sort"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}
	if got := NumIslands(grid); got != 1 {
		t.Errorf("got %d", got)
	}
	grid2 := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	if got := NumIslands(grid2); got != 3 {
		t.Errorf("got %d", got)
	}
}

func TestCloneGraph(t *testing.T) {
	a := &Node{Val: 1}
	b := &Node{Val: 2}
	c := &Node{Val: 3}
	a.Neighbors = []*Node{b, c}
	b.Neighbors = []*Node{a, c}
	c.Neighbors = []*Node{a, b}
	clone := CloneGraph(a)
	if clone == a {
		t.Error("should be a deep copy")
	}
	if clone.Val != 1 || len(clone.Neighbors) != 2 {
		t.Error("wrong structure")
	}
	if clone.Neighbors[0] == b {
		t.Error("neighbor should be cloned")
	}
}

func TestCourseSchedule(t *testing.T) {
	if !CanFinish(2, [][]int{{1, 0}}) {
		t.Error("should be possible")
	}
	if CanFinish(2, [][]int{{1, 0}, {0, 1}}) {
		t.Error("cycle -> impossible")
	}
	order := FindOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	if len(order) != 4 {
		t.Errorf("got %v", order)
	}
	// verify it's a valid topological order
	pos := map[int]int{}
	for i, c := range order {
		pos[c] = i
	}
	for _, e := range [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}} {
		if pos[e[1]] > pos[e[0]] {
			t.Errorf("prereq %d after %d", e[1], e[0])
		}
	}
}

func TestCountComponents(t *testing.T) {
	if got := CountComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := CountComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}); got != 1 {
		t.Errorf("got %d", got)
	}
}

func TestFindRedundantConnection(t *testing.T) {
	got := FindRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}})
	if got[0] != 2 || got[1] != 3 {
		t.Errorf("got %v", got)
	}
}

func TestOrangesRotting(t *testing.T) {
	if got := OrangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}); got != 4 {
		t.Errorf("got %d", got)
	}
	if got := OrangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}); got != -1 {
		t.Errorf("got %d", got)
	}
	if got := OrangesRotting([][]int{{0, 2}}); got != 0 {
		t.Errorf("got %d", got)
	}
}

func TestNetworkDelayTime(t *testing.T) {
	if got := NetworkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2); got != 2 {
		t.Errorf("got %d", got)
	}
	if got := NetworkDelayTime([][]int{{1, 2, 1}}, 2, 2); got != -1 {
		t.Errorf("got %d", got)
	}
}

func TestUnionFindBasic(t *testing.T) {
	uf := NewUnionFind(4)
	uf.Union(0, 1)
	uf.Union(2, 3)
	if uf.Find(0) != uf.Find(1) {
		t.Error("0,1 should connect")
	}
	if uf.Find(0) == uf.Find(2) {
		t.Error("0,2 should not connect")
	}
	roots := map[int]struct{}{}
	for i := 0; i < 4; i++ {
		roots[uf.Find(i)] = struct{}{}
	}
	if len(roots) != 2 {
		t.Errorf("expected 2 sets, got %d", len(roots))
	}
	sorted := []int{uf.count}
	sort.Ints(sorted)
	if uf.count != 2 {
		t.Errorf("count got %d", uf.count)
	}
}
