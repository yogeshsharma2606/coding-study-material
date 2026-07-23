# 02 - Strings

> Code: [`strings.go`](strings.go) - Tests: [`strings_test.go`](strings_test.go) - run `go test ./categories/02-strings/`

## Overview & mental model

In Go a `string` is an **immutable** read-only slice of UTF-8 bytes. You cannot assign `s[i] = ...`; convert to `[]byte` (ASCII) or `[]rune` (Unicode) first. For interview problems, three tools cover ~90% of cases:

1. **Frequency counting.** For lowercase English, a fixed `[26]int` array is faster and simpler than a map.
2. **Two pointers.** Palindromes and reversals converge from both ends.
3. **A canonical key.** Group "equal-ish" strings (anagrams) by a normalized key (sorted letters, or a count signature).

## How to recognize it

- Anagram/permutation talk -> frequency counts.
- "Palindrome", "reverse" -> two pointers.
- "Group", "categorize equal strings" -> hash by canonical key.
- "Parse / to integer / valid number" -> careful character-by-character state machine.

## How to think / attack plan

1. ASCII or Unicode? Decide `[]byte` vs `[]rune` up front - it changes indexing.
2. Is the answer about *counts* (anagram), *order* (palindrome), or *grouping*? Pick the tool.
3. Parsing problems: write down the states (skip spaces -> sign -> digits) before coding; handle overflow and bad input explicitly.
4. Never build strings with `+` in a loop (O(n^2) copies) - use `strings.Builder`.

---

## Problems (easy -> hard)

### 1. Reverse String (easy)
**Problem.** Reverse a `[]byte` in place.
**Thinking.** Two pointers from both ends, swap, converge. No extra memory.
**Complexity.** Time O(n), space O(1).

### 2. Valid Anagram (easy)
**Problem.** Is `t` a rearrangement of `s`?
**Thinking.** Anagrams have identical letter multisets. Increment counts for `s`, decrement for `t`; if any count is nonzero, they differ. A `[26]int` beats a map for a-z.
**Why not sort?** Sorting is O(n log n); counting is O(n).
**Dry run.** `"anagram"`/`"nagaram"`: every letter nets to 0 -> true.
**Complexity.** Time O(n), space O(1) (fixed 26 slots).

### 3. Valid Palindrome (easy)
**Problem.** Ignoring case and non-alphanumeric, is `s` a palindrome?
**Thinking.** Two pointers converge; skip non-alphanumeric on each side; compare lowercased. Filtering into a new string works but costs O(n) space - the in-place skip is O(1).
**Dry run.** `"A man, a plan, a canal: Panama"` -> compares a=a, m=m, ... -> true.
**Complexity.** Time O(n), space O(1).

### 4. Group Anagrams (medium)
**Problem.** Group words that are anagrams.
**Direction of thinking.** Anagrams share a canonical form. Two options for the key: (a) the **sorted** letters `O(k log k)` per word, or (b) a **count signature** `O(k)`. Use a map from key to the group.
**Why sorted key.** Simple and correct; count-signature is faster if words are long.
**Dry run.** `eat,tea,ate` all sort to `aet` -> one bucket.
**Complexity.** Time O(n * k log k), space O(n*k).

### 5. Longest Common Prefix (easy)
**Problem.** Longest prefix shared by all strings.
**Thinking.** **Vertical scan**: compare character column i across all words; stop at the first mismatch or end of any word.
**Dry run.** `flower,flow,flight`: col0 f, col1 l, col2 o vs i -> stop -> `"fl"`.
**Complexity.** Time O(total characters), space O(1).

### 6. Isomorphic Strings (easy)
**Problem.** Can characters of `s` map one-to-one onto `t` (order preserved)?
**Direction of thinking.** A valid mapping is a **bijection**: `s->t` must be consistent *and* `t->s` must be consistent (otherwise two source chars map to the same target). Keep both maps.
**Why two maps.** `"badc"`/`"baba"`: s->t is consistent but t->s is not; one map alone accepts it wrongly.
**Complexity.** Time O(n), space O(1) (bounded alphabet).

### 7. String to Integer / atoi (medium)
**Problem.** Parse an int with leading spaces, optional sign, trailing junk, clamped to int32.
**Direction of thinking.** It's a small **state machine**: skip spaces -> read optional sign -> read digits until non-digit. The subtlety is **overflow**: check against `INT_MAX/INT_MIN` *during* accumulation, not after.
**Dry run.** `"   -42"` -> skip 3 spaces -> sign - -> digits 4,2 -> -42.
**Complexity.** Time O(n), space O(1).

### 8. Longest Palindromic Substring (medium)
**Problem.** Return the longest substring that is a palindrome.
**Direction of thinking.** A palindrome mirrors around a center. There are `2n-1` centers (each char, and each gap between chars). **Expand around center**: grow left/right while characters match. Beats the O(n^3) brute force of checking every substring.
**Why expand-around-center over DP.** Same O(n^2) time but O(1) space and simpler. (Manacher's is O(n) but rarely required.)
**Dry run.** `"babad"`: center at index 1 expands to `"bab"`.
**Complexity.** Time O(n^2), space O(1).

### 9. Encode and Decode Strings (medium)
**Problem.** Serialize a list of arbitrary strings into one string, then recover it.
**Direction of thinking.** A plain delimiter fails if the payload contains it. Use a **length prefix**: `"<len>#<payload>"`. On decode, read digits up to `#`, then take exactly that many bytes - payload contents can't confuse the parser.
**Why.** Length-prefixing is the standard framing technique (also how many wire protocols work).
**Complexity.** Time O(total length), space O(total length).

---

## Common pitfalls & edge cases

- Trying to mutate a `string` directly - convert to `[]byte`/`[]rune` first.
- Using `len(s)` (bytes) when the problem is about Unicode characters (runes).
- Building strings with `+=` in a loop - use `strings.Builder`.
- atoi overflow handled after the fact instead of during accumulation.
- Off-by-one on even vs odd palindrome centers.

## Interview Q&A

- **`[]byte` vs `[]rune`?** `[]byte` indexes raw UTF-8 bytes (fine for ASCII); `[]rune` decodes to Unicode code points (needed for multibyte characters).
- **Map vs `[26]int` for counts?** The fixed array avoids hashing overhead and is O(1) space; use it when the alphabet is small and known.
- **How to compare anagrams without sorting?** Compare 26-length count vectors in O(n).
- **Why length-prefix encoding over a delimiter?** Delimiters can appear in data; a length prefix is unambiguous.
