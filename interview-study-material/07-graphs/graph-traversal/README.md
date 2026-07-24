# Graph Traversal

A graph is nodes + edges. Here they're stored as an **adjacency list** (`map[int][]int`: node → its neighbors), which is the standard representation for sparse graphs (space O(V+E)). Almost every graph problem is *some traversal* + *bookkeeping*, and the two traversals are **DFS** and **BFS**.

**The one thing that separates graphs from trees: cycles.** A tree can't revisit a node; a graph can. So **you must track `visited`** or you'll loop forever. Forgetting `visited` is the #1 graph bug.

**DFS vs BFS — pick by what you need:**

| | DFS | BFS |
|---|---|---|
| Structure | recursion / stack | **queue** |
| Explores | deep first (one path to the end) | wide first (level by level) |
| Best for | connectivity, cycle detection, topological sort, path existence | **shortest path in unweighted graphs**, level/distance |
| Space | O(depth) call stack | O(width) queue |

Both are **O(V + E)** time (visit each node once, each edge once).

## Programs

### [Depth-first search](depth-first-search/depth_first_search.go)

Recursively visit a node, mark it visited, then recurse into each **unvisited** neighbor.

**How to think about it:** DFS = "go as deep as possible, then back up." Mark visited **before** recursing so cycles (like `2 -> 0` in the sample graph) don't send you in circles. The `visited` map is the whole difference between tree recursion and graph recursion.

**Dry run** — graph `{0:[1,2], 1:[2], 2:[0,3], 3:[3]}` from 0: visit 0 → 1 → 2 (0 already visited, skip) → 3 → prints `0 1 2 3`.

- **O(V+E) time, O(V) space** (visited + recursion stack). For very deep graphs, an explicit stack avoids stack-overflow.

### [Breadth-first search](breadth-first-search/breadth_first_search.go)

Visit level by level using a **queue**; mark neighbors visited **when you enqueue** them (not when you dequeue).

**How to think about it:** BFS fans out in rings of increasing distance from the start. That's why it finds **shortest paths in unweighted graphs** — the first time you reach a node is via a shortest edge-count path. The critical detail: set `visited[nei] = true` at **enqueue** time, otherwise the same node gets queued multiple times before it's processed.

**Dry run** — same graph from 0: queue `[0]`→visit 0, enqueue 1,2 → visit 1 (2 already queued) → visit 2, enqueue 3 → visit 3. Prints `0 1 2 3`.

- **O(V+E) time, O(V) space.**

### [Undirected cycle detection](undirected-cycle-detection/undirected_cycle_detection.go)

Detect a cycle in an **undirected** graph via DFS with **parent tracking**.

**The subtlety that makes this problem (memorize it):** in an undirected graph every edge `u-v` looks like a "back edge" because `v`'s neighbor list contains `u`. So seeing an already-visited neighbor is **not** automatically a cycle — you must ignore the neighbor you *came from* (the `parent`). A cycle exists only when you reach an already-visited node that is **not** your parent.

- **O(V+E) time, O(V) space.**
- **Contrast (say this):** in a **directed** graph you instead track nodes *currently on the recursion stack* (a "gray" set) — parent tracking is only correct for undirected graphs.
- **Edge note:** to be fully correct on disconnected graphs, you'd loop over all nodes and start a DFS from each unvisited one.

## Review checklist

- Why must graph traversal track `visited` when tree traversal doesn't?
- BFS: why mark visited at enqueue time, and why does BFS give shortest paths in unweighted graphs?
- Undirected cycle detection: why is "already visited" not enough — what role does `parent` play?
- How does cycle detection differ between undirected (parent) and directed (recursion-stack) graphs?
