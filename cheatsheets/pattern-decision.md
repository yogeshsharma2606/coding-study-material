# Pattern-Decision Cheatsheet

How to map an unseen problem onto a known pattern. Read the **signals**, pick the **pattern**, confirm with the **complexity target**.

## Signal -> pattern table

| Signal in the problem | Pattern | Category |
|---|---|---|
| Sorted array, find a pair/triplet with target | Two pointers | [03](../categories/03-two-pointers/) |
| Sorted array, find/insert/first-last/rotated | Binary search | [12](../categories/12-binary-search/) |
| "Minimize the max / smallest X that's feasible" | Binary search on the answer | [12](../categories/12-binary-search/) |
| "Contiguous subarray/substring" + optimize length | Sliding window | [04](../categories/04-sliding-window/) |
| "Subarray sum == k", range sums, +negatives | Prefix sum (+ hash map) | [05](../categories/05-prefix-sum/) |
| "Have I seen it", counts, grouping, dedup | Hashing | [06](../categories/06-hashing/) |
| Reverse/reorder/cycle in a linked list | Linked list (dummy, fast/slow) | [07](../categories/07-linked-list/) |
| Matching brackets, evaluate expression, nesting | Stack | [08](../categories/08-stack/) |
| "Next/previous greater/smaller", histogram area | Monotonic stack | [09](../categories/09-monotonic-stack/) |
| Sliding-window max/min, FIFO processing | Queue / deque | [10](../categories/10-queue-deque/) |
| "Top/kth largest/smallest/closest", stream median | Heap | [11](../categories/11-heap-priority-queue/) |
| Custom order, kth element, small integer keys | Sorting / quickselect / counting | [13](../categories/13-sorting/) |
| Tree traversal, depth, ancestor, validate BST | Trees (DFS/BFS) | [14](../categories/14-trees/) |
| Prefix / autocomplete / dictionary / max XOR | Trie | [15](../categories/15-tries/) |
| Islands, connected, shortest steps, prerequisites | Graphs (BFS/DFS/topo/UF/Dijkstra) | [16](../categories/16-graphs/) |
| "Generate all / list every / any valid config" | Backtracking | [17](../categories/17-recursion-backtracking/) |
| "Number of ways / min cost / can reach", overlapping subproblems | Dynamic programming | [18](../categories/18-dynamic-programming/) |
| Locally-optimal choice works (prove it) | Greedy | [19](../categories/19-greedy/) |
| List of [start, end], merge/overlap/rooms | Intervals (sort + sweep) | [20](../categories/20-intervals/) |
| Appears once/twice/3x, count bits, no +/- | Bit manipulation | [21](../categories/21-bit-manipulation/) |
| GCD/primes/power/overflow/digits | Math & number theory | [22](../categories/22-math-number-theory/) |
| Rotate/spiral/2D search/regions | Matrix / grid | [23](../categories/23-matrix-grid/) |
| "Design cache/LRU/LFU, O(1) get&put" | Design (map + aux structure) | [24](../categories/24-design-cache/) |
| "Parallel/concurrent/rate-limit/cancel" | Concurrency | [25](../categories/25-concurrency/) |

## Decision flow

```mermaid
flowchart TD
    start["Read problem, restate, note constraints"] --> arr{"Data type?"}
    arr -->|"Array/String"| sorted{"Sorted?"}
    arr -->|"Linked list"| ll["Two pointers / dummy / fast-slow (07)"]
    arr -->|"Tree"| tree["DFS/BFS; BST -> inorder (14)"]
    arr -->|"Graph/Grid"| graph["BFS/DFS/topo/UF/Dijkstra (16, 23)"]

    sorted -->|"Yes"| bs["Two pointers or binary search (03, 12)"]
    sorted -->|"No"| goal{"What is asked?"}

    goal -->|"Contiguous subarray optimize"| win["Sliding window (04)"]
    goal -->|"Subarray sum/count with =k"| pre["Prefix sum + hashmap (05)"]
    goal -->|"Next greater/smaller"| mono["Monotonic stack (09)"]
    goal -->|"Top-k / median / stream"| heap["Heap (11)"]
    goal -->|"Seen-before / group / count"| hash["Hashing (06)"]

    goal -->|"Generate all / combinations"| bt["Backtracking (17)"]
    goal -->|"Count ways / min cost / reach"| dp["Dynamic programming (18)"]
    goal -->|"Local best is optimal"| greedy["Greedy - prove it! (19)"]
```

## When two patterns compete

- **Sliding window vs prefix sum**: window needs non-negative values (monotonic shrink); prefix sum handles negatives and exact-sum counting.
- **Two pointers vs hashing** (pair sum): sorted or O(1) space -> two pointers; need original indices or can't sort -> hashing.
- **Heap vs quickselect vs sort** (kth): heap for streaming/top-k O(n log k); quickselect O(n) avg one-shot; full sort if you also need order.
- **Greedy vs DP**: try greedy, then break it with a counterexample. If a locally worse choice can win globally, use DP.
- **BFS vs DFS**: shortest path in unweighted -> BFS; connectivity/paths/aggregates -> DFS.

## The 7-step interview process (say it out loud)

1. Restate + clarify (size, ranges, duplicates, sortedness, mutability).
2. Small example by hand.
3. Brute force + its complexity.
4. Find the repeated work / bottleneck.
5. Map to a pattern; state target complexity.
6. Dry run the approach, then code cleanly.
7. Test edge cases; state final time/space.
