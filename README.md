# Coding Interview Study Material (Go)

A **senior / principal-engineer**-grade study repository for the coding rounds of technical interviews. It is designed to be *enough on its own*: every category teaches **how to recognize the pattern**, **how to think your way to the solution**, and then walks real problems from **easy to hard** with a **direction of thinking**, an **approach + why**, a **dry run**, a **runnable Go solution**, and **time/space complexity**.

Unlike a raw list of solutions, this repo focuses on the thing interviews actually test at the senior level: **can you map an unseen problem onto a known pattern and justify your trade-offs out loud?**

---

## How to use this repo

1. Read this README fully once - especially the **"How to think in an interview"** framework and the **complexity primer** below.
2. Study one category at a time from [`categories/`](categories/). Each folder has a `README.md` (the lesson) plus one `.go` file per problem and a matching `_test.go`.
3. Run the code. Everything is standard-library only, so:
   ```bash
   go test ./...          # run every solution's tests
   go test ./categories/01-arrays/...   # one category
   ```
4. Before looking at a solution, try it yourself, then compare your reasoning to the **Direction of thinking** section.
5. Use the [cheatsheets](cheatsheets/) for last-minute revision.

---

## How to think in an interview (the universal framework)

Most candidates fail not because they don't know algorithms, but because they jump to code. Use this sequence out loud - it is what senior interviewers grade:

1. **Restate & clarify.** Repeat the problem in your own words. Ask about input size, value ranges, duplicates, sortedness, empty/negative inputs, and whether the input can be mutated. Constraints *are* hints (see the table below).
2. **Work a small example by hand.** Concrete examples expose the pattern and become your test cases.
3. **State the brute force.** Always have a correct O(n^2)/O(2^n) baseline. Say its complexity. This shows you can solve it at all, and gives you something to optimize.
4. **Find the bottleneck & the redundancy.** Ask: *what work am I repeating?* Removing repeated work is the core of almost every optimization (hashing to avoid re-search, DP to avoid re-computation, two pointers to avoid re-scan).
5. **Map to a pattern.** Match the signals to one of the categories here (see [pattern-decision cheatsheet](cheatsheets/pattern-decision.md)).
6. **Design, then dry run, then code.** Verbally trace your approach on the small example *before* writing code. Then write clean code.
7. **Test & analyze.** Walk your own code on the example, hit edge cases, and state final time/space complexity.

### Constraints -> pattern (the "hint decoder")

| You see... | Likely direction |
|---|---|
| Sorted array / "find pair/target" | Two pointers or binary search |
| "Contiguous subarray/substring" | Sliding window or prefix sum |
| "Subarray sum equals k", range sums | Prefix sum + hash map |
| n up to ~10^18, or "count steps" | O(log n): binary search / bit / math |
| n up to ~10^5-10^6 | Need O(n) or O(n log n) |
| n up to ~5000 | O(n^2) is probably fine |
| n up to ~20-25 | Exponential / bitmask / backtracking |
| "Next greater / previous smaller" | Monotonic stack |
| "Top / k largest / k closest / stream" | Heap (priority queue) |
| "All combinations / permutations / partitions" | Backtracking |
| "Number of ways / min cost / can you reach" | Dynamic programming |
| "Shortest path / levels / spread" | BFS |
| "Connected components / cycle / islands" | DFS / Union-Find |
| "Order with prerequisites" | Topological sort |
| "Prefix / autocomplete / dictionary" | Trie |
| "Design a cache / O(1) get&put" | Hash map + doubly linked list |
| "Concurrent / parallel / rate limit" | Goroutines + channels / sync |

---

## Complexity primer (the tiers you must know cold)

From best to worst; assume n is the input size:

| Big-O | Name | Typical source | Feasible n (approx.) |
|---|---|---|---|
| O(1) | Constant | Hash lookup, arithmetic | any |
| O(log n) | Logarithmic | Binary search, balanced tree, heap push/pop | any |
| O(n) | Linear | Single scan, two pointers, BFS/DFS | ~10^8 |
| O(n log n) | Linearithmic | Sorting, heap of n items | ~10^7 |
| O(n^2) | Quadratic | Nested loops, naive pair check | ~10^4 |
| O(n^2 ... n^3) | Poly | Many DP tables | ~10^3 |
| O(2^n) | Exponential | Subsets, brute-force recursion | ~20-25 |
| O(n!) | Factorial | Permutations | ~10-12 |

Rules of thumb: an interviewer's "can you do better?" almost always means *drop one tier* (n^2 -> n log n, or n log n -> n). **Space** counts your extra structures and, for recursion, the **call-stack depth** (often O(n) or O(log n)).

Common data-structure operation costs are in [cheatsheets/complexity.md](cheatsheets/complexity.md).

---

## Curriculum (25 categories)

### Fundamentals & linear scans
| # | Category | Core idea |
|---|---|---|
| 01 | [Arrays](categories/01-arrays/) | In-place tricks, scanning invariants, Kadane, Dutch flag |
| 02 | [Strings](categories/02-strings/) | Frequency maps, parsing, palindromes, encoding |
| 03 | [Two Pointers](categories/03-two-pointers/) | Converging/ fast-slow pointers on sorted/linked data |
| 04 | [Sliding Window](categories/04-sliding-window/) | Variable/fixed windows for contiguous subarrays |
| 05 | [Prefix Sum](categories/05-prefix-sum/) | Range queries and "subarray-sum = k" in O(n) |
| 06 | [Hashing](categories/06-hashing/) | Trade space for O(1) lookup; grouping & dedup |

### Linked structures, LIFO/FIFO, and ordering by priority
| # | Category | Core idea |
|---|---|---|
| 07 | [Linked List](categories/07-linked-list/) | Pointer surgery, dummy nodes, cycle detection |
| 08 | [Stack](categories/08-stack/) | LIFO, matching, expression evaluation |
| 09 | [Monotonic Stack](categories/09-monotonic-stack/) | Next/previous greater/smaller in O(n) |
| 10 | [Queue & Deque](categories/10-queue-deque/) | FIFO, monotonic deque for window max |
| 11 | [Heap / Priority Queue](categories/11-heap-priority-queue/) | Top-k, streams, k-way merge |

### Searching & sorting
| # | Category | Core idea |
|---|---|---|
| 12 | [Binary Search](categories/12-binary-search/) | Search space reduction, "binary search on the answer" |
| 13 | [Sorting](categories/13-sorting/) | Comparators, counting sort, quickselect, merge |

### Trees & graphs
| # | Category | Core idea |
|---|---|---|
| 14 | [Trees](categories/14-trees/) | Traversals, BST invariants, recursion on trees |
| 15 | [Tries](categories/15-tries/) | Prefix trees for dictionaries/autocomplete |
| 16 | [Graphs](categories/16-graphs/) | BFS/DFS, topological sort, Union-Find, Dijkstra |

### Algorithmic paradigms
| # | Category | Core idea |
|---|---|---|
| 17 | [Recursion & Backtracking](categories/17-recursion-backtracking/) | Systematic search of the solution space |
| 18 | [Dynamic Programming](categories/18-dynamic-programming/) | Overlapping subproblems + optimal substructure |
| 19 | [Greedy](categories/19-greedy/) | Locally optimal choices with an exchange argument |
| 20 | [Intervals](categories/20-intervals/) | Sort-by-endpoint sweeping and merging |

### Low-level & math
| # | Category | Core idea |
|---|---|---|
| 21 | [Bit Manipulation](categories/21-bit-manipulation/) | XOR tricks, masks, power-of-two |
| 22 | [Math & Number Theory](categories/22-math-number-theory/) | GCD, primes, modular arithmetic, overflow |
| 23 | [Matrix / Grid](categories/23-matrix-grid/) | 2D traversal, rotation, grid BFS/DFS |

### Systems-flavored (senior focus)
| # | Category | Core idea |
|---|---|---|
| 24 | [Design & Cache](categories/24-design-cache/) | LRU/LFU, rate limiter, O(1) data-structure design |
| 25 | [Concurrency](categories/25-concurrency/) | Goroutines, channels, worker pools, sync primitives |

---

## Cheatsheets

- [Complexity cheatsheet](cheatsheets/complexity.md) - Big-O tiers and per-structure operation costs
- [Pattern-decision cheatsheet](cheatsheets/pattern-decision.md) - signal -> pattern recognition
- [Go idioms cheatsheet](cheatsheets/go-idioms.md) - slices, maps, `container/heap`, `sort`, generics gotchas

---

## Repo conventions

- **Standard library only** - clones and runs anywhere with a Go toolchain (Go 1.22+).
- Each problem is an **exported function** with a table-driven test (`*_test.go`).
- Markdown lessons reference the real `.go` files, so lessons and code never drift.
- Category folders are independent Go packages under [`categories/`](categories/).

Run the full suite:

```bash
go test ./...
```
