// Package graphs contains graph interview problems.
//
// A graph is nodes + edges. The two traversals cover most problems: BFS (queue)
// finds shortest paths in UNWEIGHTED graphs and explores level by level; DFS
// (stack/recursion) explores fully and detects cycles/components. Beyond those:
// TOPOLOGICAL SORT orders a DAG by dependencies, UNION-FIND tracks connectivity
// incrementally, and DIJKSTRA (a heap) finds shortest paths with non-negative
// weights. Grids are just implicit graphs (4-directional neighbors).
package graphs

import "container/heap"

// --- 1. Number of Islands: grid DFS flood fill ---

func NumIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	var sink func(r, c int)
	sink = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0' // mark visited by sinking
		sink(r+1, c)
		sink(r-1, c)
		sink(r, c+1)
		sink(r, c-1)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				sink(r, c)
			}
		}
	}
	return count
}

// --- 2. Clone Graph: DFS with a visited map ---

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	clones := map[*Node]*Node{}
	var dfs func(n *Node) *Node
	dfs = func(n *Node) *Node {
		if c, ok := clones[n]; ok {
			return c
		}
		copy := &Node{Val: n.Val}
		clones[n] = copy // register before recursing to handle cycles
		for _, nb := range n.Neighbors {
			copy.Neighbors = append(copy.Neighbors, dfs(nb))
		}
		return copy
	}
	return dfs(node)
}

// --- 3 & 4. Course Schedule (cycle detection) and ordering (Kahn's topo sort) ---

// CanFinish reports whether all courses can be finished (no cycle in prereqs).
func CanFinish(numCourses int, prerequisites [][]int) bool {
	return len(topoOrder(numCourses, prerequisites)) == numCourses
}

// FindOrder returns a valid course order, or empty if impossible (cycle).
func FindOrder(numCourses int, prerequisites [][]int) []int {
	order := topoOrder(numCourses, prerequisites)
	if len(order) != numCourses {
		return []int{}
	}
	return order
}

// topoOrder runs Kahn's algorithm: repeatedly remove nodes with in-degree 0.
func topoOrder(n int, edges [][]int) []int {
	adj := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		course, pre := e[0], e[1]
		adj[pre] = append(adj[pre], course) // pre -> course
		indeg[course]++
	}
	var queue []int
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}
	var order []int
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		for _, v := range adj[u] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return order
}

// --- 5 & 8. Union-Find (Disjoint Set Union) ---

type UnionFind struct {
	parent []int
	rank   []int
	count  int // number of disjoint sets
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{parent: make([]int, n), rank: make([]int, n), count: n}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

// Find returns the set representative with path compression.
func (uf *UnionFind) Find(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]] // path halving
		x = uf.parent[x]
	}
	return x
}

// Union merges two sets by rank; returns false if already connected.
func (uf *UnionFind) Union(a, b int) bool {
	ra, rb := uf.Find(a), uf.Find(b)
	if ra == rb {
		return false
	}
	if uf.rank[ra] < uf.rank[rb] {
		ra, rb = rb, ra
	}
	uf.parent[rb] = ra
	if uf.rank[ra] == uf.rank[rb] {
		uf.rank[ra]++
	}
	uf.count--
	return true
}

// CountComponents counts connected components in an undirected graph.
func CountComponents(n int, edges [][]int) int {
	uf := NewUnionFind(n)
	for _, e := range edges {
		uf.Union(e[0], e[1])
	}
	return uf.count
}

// FindRedundantConnection returns the edge that creates a cycle (the last one
// joining two already-connected nodes).
func FindRedundantConnection(edges [][]int) []int {
	uf := NewUnionFind(len(edges) + 1)
	for _, e := range edges {
		if !uf.Union(e[0], e[1]) {
			return e
		}
	}
	return nil
}

// --- 6. Rotting Oranges: multi-source BFS ---

// OrangesRotting returns minutes until no fresh orange remains, or -1.
// 0 empty, 1 fresh, 2 rotten. Start BFS from ALL rotten cells simultaneously.
func OrangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var queue [][2]int
	fresh := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0
	for len(queue) > 0 && fresh > 0 {
		minutes++
		for sz := len(queue); sz > 0; sz-- {
			cell := queue[0]
			queue = queue[1:]
			for _, d := range dirs {
				nr, nc := cell[0]+d[0], cell[1]+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}

// --- 7. Network Delay Time: Dijkstra with a min-heap ---

type edge struct{ node, dist int }
type edgeHeap []edge

func (h edgeHeap) Len() int           { return len(h) }
func (h edgeHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h edgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *edgeHeap) Push(x any)        { *h = append(*h, x.(edge)) }
func (h *edgeHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// NetworkDelayTime returns the time for a signal from k to reach all n nodes,
// or -1 if unreachable. times[i] = [u, v, w].
func NetworkDelayTime(times [][]int, n, k int) int {
	adj := make(map[int][][2]int) // u -> list of (v, w)
	for _, t := range times {
		adj[t[0]] = append(adj[t[0]], [2]int{t[1], t[2]})
	}
	dist := make(map[int]int)
	h := &edgeHeap{{k, 0}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(edge)
		if _, seen := dist[cur.node]; seen {
			continue // already finalized (lazy deletion)
		}
		dist[cur.node] = cur.dist
		for _, nb := range adj[cur.node] {
			if _, seen := dist[nb[0]]; !seen {
				heap.Push(h, edge{nb[0], cur.dist + nb[1]})
			}
		}
	}
	if len(dist) != n {
		return -1
	}
	maxDist := 0
	for _, d := range dist {
		if d > maxDist {
			maxDist = d
		}
	}
	return maxDist
}
