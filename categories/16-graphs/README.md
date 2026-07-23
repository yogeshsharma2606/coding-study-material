# 16 - Graphs

> Code: [`graphs.go`](graphs.go) - Tests: [`graphs_test.go`](graphs_test.go) - run `go test ./categories/16-graphs/`

## Overview & mental model

A graph is nodes connected by edges (directed or undirected, weighted or not). Most interview graph problems are one of a handful of algorithms:

- **BFS** (queue) - shortest path in an **unweighted** graph, level-by-level spread, multi-source flood.
- **DFS** (recursion/stack) - full exploration, connectivity, cycle detection, flood fill.
- **Topological sort** - order a **DAG** by dependencies (Kahn's in-degree BFS, or DFS post-order).
- **Union-Find (DSU)** - incremental connectivity ("are these connected?", counting components, cycle detection) in near-O(1) amortized.
- **Dijkstra** (min-heap) - shortest path with **non-negative weights**.

A grid is an implicit graph: each cell connects to its 4 (or 8) neighbors.

## How to recognize it

- "Islands / regions / connected", "reachable", "shortest steps", "spread/infection".
- "Order with prerequisites / dependencies" -> topological sort.
- "Are X and Y connected", "number of groups", "does adding this edge form a cycle" -> Union-Find.
- "Shortest/cheapest path with weights" -> Dijkstra (non-negative) or Bellman-Ford (negative).

## How to think / attack plan

1. Model it: what are nodes, what are edges? For grids, neighbors are the graph.
2. Weighted? No -> BFS for shortest, DFS for reach. Yes, non-negative -> Dijkstra.
3. Dependencies/ordering -> topological sort; if no valid order exists, there's a cycle.
4. Dynamic connectivity / grouping -> Union-Find.
5. Track **visited** to avoid infinite loops (cycles) and redundant work.

## How to represent a graph in Go

- **Adjacency list**: `map[int][]int` or `[][]int` - best for sparse graphs.
- **Edge list**: `[][]int` of `[u,v,w]` - convenient for Union-Find / Kruskal.
- **Grid**: the 2D slice itself; compute neighbors on the fly.

---

## Problems (easy -> hard)

### 1. Number of Islands (medium) - grid DFS flood fill
**Direction of thinking.** Each unvisited `'1'` starts a new island; DFS/BFS "sinks" its whole component so it's counted once. Marking visited in place (turn `'1'`->`'0'`) saves a visited set.
**Complexity.** Time O(rows*cols), space O(rows*cols) recursion worst case.

### 2. Clone Graph (medium)
**Direction of thinking.** DFS while mapping original->clone. **Register the clone before recursing** into neighbors, so cycles resolve to the already-created copy instead of looping forever.
**Complexity.** Time O(V+E), space O(V).

### 3. Course Schedule (medium) - cycle detection
**Direction of thinking.** Courses with prerequisites form a directed graph; you can finish all iff it's a **DAG** (no cycle). Kahn's algorithm removes in-degree-0 nodes repeatedly; if you can't remove all, a cycle remains.
**Complexity.** Time O(V+E), space O(V+E).

### 4. Course Schedule II (medium) - topological order
**Direction of thinking.** Same Kahn's process, but record the removal order - that's a valid topological sort. Empty result signals a cycle.
**Dry run.** prereqs 1,2 need 0; 3 needs 1,2 -> order like `[0,1,2,3]`.
**Complexity.** Time O(V+E).

### 5. Number of Connected Components (medium) - Union-Find
**Direction of thinking.** Start with n singleton sets; each edge unions two sets. The remaining set count is the component count. With path compression + union by rank, operations are near O(1) amortized (inverse-Ackermann).
**Complexity.** Time O(E * alpha(n)), space O(n).

### 6. Rotting Oranges (medium) - multi-source BFS
**Direction of thinking.** All initially-rotten cells rot their neighbors simultaneously, so seed the BFS queue with **every** rotten cell and process level by level; each level is one minute. Track fresh count to detect unreachable oranges (-1).
**Complexity.** Time O(rows*cols), space O(rows*cols).

### 7. Network Delay Time (medium) - Dijkstra
**Direction of thinking.** Shortest time to reach all nodes from a source with non-negative edge weights -> Dijkstra with a min-heap. Pop the closest unfinalized node, relax its edges. The answer is the max finalized distance (or -1 if some node is unreachable).
**Complexity.** Time O(E log V), space O(V+E).

### 8. Redundant Connection (medium) - Union-Find cycle
**Direction of thinking.** Process edges; the first edge that connects two **already-connected** nodes closes a cycle - return it. Union-Find detects this in one pass.
**Complexity.** Time O(E * alpha(n)), space O(n).

---

## Common pitfalls & edge cases

- Forgetting `visited` -> infinite loops on cyclic graphs.
- BFS gives shortest paths only for **unweighted** (or uniform-weight) graphs; use Dijkstra otherwise.
- Dijkstra breaks with **negative** weights - use Bellman-Ford / SPFA.
- Topological sort requires a DAG; detect cycles (leftover nodes) explicitly.
- Union-Find without path compression/union-by-rank degrades toward O(n) per op.
- Deep grid DFS can overflow the stack; BFS or an explicit stack is safer for huge grids.

## Interview Q&A

- **BFS vs DFS for shortest path?** BFS (unweighted) guarantees shortest by levels; DFS does not.
- **How does topological sort detect a cycle?** If fewer than V nodes get ordered, a cycle blocked the rest.
- **Why is Union-Find near O(1)?** Path compression flattens trees and union-by-rank keeps them shallow; amortized inverse-Ackermann.
- **When Dijkstra vs Bellman-Ford?** Dijkstra for non-negative weights (faster, O(E log V)); Bellman-Ford handles negatives and detects negative cycles (O(VE)).
- **Adjacency list vs matrix?** List for sparse graphs (O(V+E) space); matrix for dense graphs or O(1) edge lookups (O(V^2) space).
