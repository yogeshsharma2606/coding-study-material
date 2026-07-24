# Anagrams

Two strings are anagrams if one is a rearrangement of the other — i.e. they have the **identical multiset of characters**. Every anagram problem reduces to picking a **canonical signature** that's equal for anagrams and different otherwise. Two signatures dominate:

- **Frequency count** (26-int array or map): O(n) to build.
- **Sorted characters**: O(n log n), but a convenient map key.

## Programs

### [Valid anagram](valid-anagram/valid_anagram.go)

Are `s` and `t` anagrams?

**How to think about it:** first the free rejection — **different lengths can't be anagrams**, return early. Then count: increment a frequency map for every char in `s`, decrement for every char in `t`. If any count goes **negative**, `t` has a char `s` didn't (or too many of it) → not an anagram. Equal lengths + never-negative ⇒ all counts net to zero ⇒ anagram.

**Dry run** — `"anagram"` vs `"nagaram"`: counts from `s` all return to 0 as `t` decrements them, never negative → **true**.

- **O(n) time, O(1) space** for a fixed alphabet (26 letters), O(k) for the map.
- **Alternative:** sort both and compare — O(n log n); the counting approach is strictly better.

### [Group anagrams](group-anagrams/group_anagrams.go)

Group words that are anagrams of each other.

**The key idea:** all anagrams collapse to the **same canonical key**. Here the key is the word's characters **sorted** — `"eat"`, `"tea"`, `"ate"` all become `"aet"`. Bucket words into a `map[key][]string`, then return the buckets.

**Dry run** — `["eat","tea","tan","ate","nat","bat"]`:

```
eat -> key "aet" : [eat]
tea -> key "aet" : [eat, tea]
tan -> key "ant" : [tan]
ate -> key "aet" : [eat, tea, ate]
nat -> key "ant" : [tan, nat]
bat -> key "abt" : [bat]
```

Groups: `[eat tea ate] [tan nat] [bat]`.

- **O(n · k log k) time** (n words, each length k, sorted for the key), **O(n·k) space.**
- **Faster key:** instead of sorting, use a 26-length count turned into a string — makes the key O(k) so total is **O(n·k)**. Good "can you do better?" answer.
- **Gotcha:** map iteration order in Go is random, so the group order isn't deterministic (only the grouping matters).

## Review checklist

- What is a "canonical signature" and why does every anagram problem need one?
- Count vs sort key: state the complexity of each and when you'd pick which.
- In valid-anagram, why does a count going negative immediately disprove it?
- Why is Go's group output order nondeterministic, and does it matter here?
