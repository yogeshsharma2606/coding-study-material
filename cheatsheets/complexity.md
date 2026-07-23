# Complexity Cheatsheet

## Big-O tiers (memorize)

| Big-O | Name | Source | Feasible n |
|---|---|---|---|
| O(1) | constant | hash lookup, arithmetic, stack push | any |
| O(log n) | logarithmic | binary search, balanced BST, heap push/pop | any |
| O(n) | linear | single scan, BFS/DFS, two pointers | ~10^8 |
| O(n log n) | linearithmic | sorting, heap of n, divide & conquer | ~10^7 |
| O(n^2) | quadratic | nested loops, naive pair check | ~10^4 |
| O(n^3) | cubic | Floyd-Warshall, some DP | ~500 |
| O(2^n) | exponential | subsets, brute-force recursion | ~20-25 |
| O(n!) | factorial | permutations | ~10-12 |

"Can you do better?" almost always means **drop one tier**: n^2 -> n log n, or n log n -> n.

## Constraint -> target complexity

| n | Expected solution |
|---|---|
| <= 10-12 | O(n!) permutations |
| <= 20-25 | O(2^n) subsets / bitmask DP |
| <= 500 | O(n^3) |
| <= 5,000 | O(n^2) |
| <= 10^6 | O(n log n) or O(n) |
| <= 10^9 or 10^18 | O(log n) or O(1) (binary search / math / bits) |

## Data-structure operation costs

| Structure | Access | Search | Insert | Delete | Notes |
|---|---|---|---|---|---|
| Array / slice | O(1) | O(n) | O(n) mid, O(1)* append | O(n) | *amortized append |
| Hash map/set | - | O(1) avg | O(1) avg | O(1) avg | O(n) worst (collisions) |
| Singly linked list | O(n) | O(n) | O(1)* | O(1)* | *given the node/prev |
| Stack / queue (slice) | - | - | O(1) | O(1) | amortized |
| Binary heap | O(1) peek | O(n) | O(log n) | O(log n) | build-heap O(n) |
| Balanced BST | O(log n) | O(log n) | O(log n) | O(log n) | ordered iteration |
| Trie | - | O(L) | O(L) | O(L) | L = key length |
| Union-Find | - | O(alpha) | O(alpha) | - | alpha ~ constant |

## Algorithm complexities

| Algorithm | Time | Space |
|---|---|---|
| Binary search | O(log n) | O(1) |
| Merge sort | O(n log n) | O(n) |
| Quicksort | O(n log n) avg / O(n^2) worst | O(log n) |
| Heapsort | O(n log n) | O(1) |
| Counting/radix sort | O(n + k) | O(n + k) |
| Quickselect (kth) | O(n) avg / O(n^2) worst | O(1) |
| BFS / DFS | O(V + E) | O(V) |
| Dijkstra (heap) | O(E log V) | O(V + E) |
| Bellman-Ford | O(V * E) | O(V) |
| Topological sort | O(V + E) | O(V + E) |
| Sieve of Eratosthenes | O(n log log n) | O(n) |
| Fast exponentiation | O(log n) | O(1) |

## Space complexity reminders

- Count **extra** space, not the input.
- Recursion uses stack space = **max depth** (O(h) for trees, O(n) for a skewed recursion, O(log n) for balanced).
- A "constant number of variables" is O(1) even inside a loop.
- Output size sometimes dominates (e.g. generating all 2^n subsets is O(n * 2^n) space to store them).

## Amortized vs worst-case

- **Amortized**: average over a sequence of operations (slice append, hash map ops, two-pointer inner loops). Legit to quote, but mention the worst case.
- **Worst-case**: the single most expensive operation (a hash collision storm, a quicksort bad-pivot).
