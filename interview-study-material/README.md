# Go Interview Study Material

This curated copy groups the repository's programs by interview pattern. Each topic README is a **complete lesson**, not just a file list: it teaches how to *recognize* the pattern, the *direction of thinking*, the *approach and why it works*, a *dry run*, *time/space complexity*, and a *review checklist* — so the ideas stick after a single read. (See the parent [repo README](../README.md) for the universal interview framework, the constraints→pattern "hint decoder", and the complexity primer.)

## How to use this material

1. Read the README in a topic folder **fully** — it contains the reasoning, not just the code.
2. Explain the brute-force approach out loud before reading the optimized solution.
3. State time and space complexity aloud, and answer the review-checklist questions from memory.
4. Run a program from its own directory with `go run .` when it contains a demo entrypoint.
5. Run all compile checks and tests from this directory with `go test ./...`.

Some entries are reusable data structures or solution functions rather than executables; use `go test` to compile-check them.

## Topics (suggested order)

| # | Topic | Lessons |
|---|---|---|
| 1 | Data structures | [Linked lists](01-data-structures/linked-lists/README.md) · [Stack & queue](01-data-structures/stack-and-queue/README.md) · [Binary trees](01-data-structures/trees/README.md) |
| 2 | Searching & sorting | [Binary search](02-searching-and-sorting/binary-search/README.md) · [Sorting](02-searching-and-sorting/sorting/README.md) |
| 3 | Arrays, hashing & subarrays | [Subarrays](03-arrays-and-hashing/subarrays/README.md) · [Array transformations](03-arrays-and-hashing/array-transformations/README.md) · [Hashing](03-arrays-and-hashing/hashing/README.md) |
| 4 | Sliding window & two pointers | [Sliding window](04-sliding-window-and-two-pointers/sliding-window/README.md) · [Two pointers](04-sliding-window-and-two-pointers/two-pointers/README.md) |
| 5 | Monotonic structures | [Monotonic stack](05-monotonic-structures/monotonic-stack/README.md) |
| 6 | Strings | [String processing](06-strings/string-processing/README.md) · [Anagrams](06-strings/anagrams/README.md) |
| 7 | Graphs | [Graph traversal](07-graphs/graph-traversal/README.md) |
| 8 | Greedy | [Stock problems](08-greedy/stock-problems/README.md) |
| 9 | Go concurrency | [Concurrency patterns](09-go-concurrency/concurrency-patterns/README.md) |
| 10 | System design | [Caching](10-system-design/caching/README.md) · [Resilience & traffic control](10-system-design/resilience-and-control/README.md) |
| 11 | Applied problems | [Advanced exercises](11-applied-problems/advanced-exercises/README.md) |

## Naming conventions

- Topic and program directories use lowercase kebab-case.
- Go filenames use lowercase snake_case.
- Misspellings and unclear names from the source collection are corrected in this curated copy.
- Nonstandard demo names such as `main2` and `mainzz` are normalized to `main` in isolated copies.

Programs indexed: 57
