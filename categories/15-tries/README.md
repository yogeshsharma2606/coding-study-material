# 15 - Tries (Prefix Trees)

> Code: [`tries.go`](tries.go) - Tests: [`tries_test.go`](tries_test.go) - run `go test ./categories/15-tries/`

## Overview & mental model

A trie stores a set of strings by their **shared prefixes**. Each node represents one character; the path from the root to a node spells a prefix, and a flag marks where complete words end. The payoff: insert and lookup are **O(L)** in the word length, **independent of the number of stored words**, and prefix queries are natural. That's why tries power autocomplete, spell-check, IP routing, and dictionary problems.

A **bit trie** (branching on 0/1) over the binary digits of integers turns XOR/bitmask problems into greedy root-to-leaf walks.

## How to recognize it

- "Prefix", "startsWith", "autocomplete", "dictionary", "word with wildcards".
- Many queries against a fixed word set (amortize the build).
- "Maximum XOR pair", "shortest root/prefix replacement".

## How to think / attack plan

1. Am I doing repeated **prefix** work over a set of strings? Build a trie once, query many times.
2. Node layout: a fixed `[26]*Node` array (lowercase) is simple and fast; a `map[byte]*Node` handles arbitrary alphabets.
3. Wildcards (`.`) -> DFS that branches over all children at that position.
4. For "every prefix must also be a word" or "shortest root", check the `isEnd` flag as you descend.
5. Integer XOR problems -> bit trie, greedily choosing the opposite bit.

## Core structure

```go
type Trie struct {
    children [26]*Trie
    isEnd    bool
}
```

---

## Problems (easy -> hard)

### 1. Implement Trie (medium)
**Thinking.** `Insert` walks/creates nodes and marks the end; `Search` requires `isEnd`; `StartsWith` just needs the path to exist. Each is O(L).
**Complexity.** O(L) per op, space O(total chars).

### 2. Add and Search Word - wildcards (medium)
**Direction of thinking.** `.` matches any character, so on a `.` the search **branches** into every existing child via DFS. Concrete characters follow the single matching edge.
**Dry run.** dict {bad,dad,mad}; `".ad"` branches at the root into b/d/m -> matches.
**Complexity.** O(L) normal; wildcards can cost O(26^k) in the worst case.

### 3. Replace Words (medium)
**Direction of thinking.** Build a trie of dictionary **roots**; for each word, walk until you hit a node marked `isEnd` - that's the shortest root prefix. Replace with it.
**Complexity.** Build O(sum of roots), replace O(sum of words).

### 4. Longest Word in Dictionary (medium)
**Direction of thinking.** A word qualifies only if **every prefix is also a word**. DFS the trie moving only into children that are themselves word-ends; track the longest (lexicographically smallest on ties).
**Complexity.** O(total chars).

### 5. Search Suggestions System (medium)
**Direction of thinking.** After inserting all products, for each growing prefix descend to the prefix node and collect up to 3 words in lexicographic order (children are visited a..z, so a DFS yields sorted results).
**Complexity.** Build O(sum of products); each query O(prefix + output).

### 6. Map Sum Pairs (medium)
**Direction of thinking.** To answer prefix-sum queries in O(prefix), store a running `sum` at every node and add the value delta along the insert path. Support overwrites by tracking the previous value and applying the difference.
**Complexity.** O(L) insert and query.

### 7. Maximum XOR of Two Numbers (medium-hard) - bit trie
**Direction of thinking.** Insert each number's 32 bits (MSB first) into a binary trie. For each number, walk the trie preferring the **opposite** bit at each level (that sets the XOR bit to 1). The greedy path yields the max XOR partner. Turns an O(n^2) pairwise check into O(32n).
**Dry run.** `[3,10,5,25,2,8]` -> 5 XOR 25 = 28.
**Complexity.** Time O(32n), space O(32n).

---

## Common pitfalls & edge cases

- Forgetting the `isEnd` flag - then `Search` can't distinguish a word from a mere prefix.
- Wildcard search without pruning can explode; note the worst case.
- Using a `[26]` array for non-lowercase input - switch to a map.
- Bit trie: iterate bits from most significant to least, and pick a consistent bit width.

## Interview Q&A

- **Trie vs hash set for words?** A hash set answers exact membership in O(L) too, but a trie also answers **prefix** queries and shares memory across common prefixes.
- **Space cost of a trie?** O(total characters), with per-node overhead; a `[26]` array wastes space on sparse nodes (use a map to save memory).
- **Why does the bit-trie greedy work for max XOR?** Higher bits dominate the value, so choosing an opposing high bit whenever possible is optimal.
- **How does autocomplete use tries?** Descend to the prefix node, then enumerate/rank the subtree's words.
