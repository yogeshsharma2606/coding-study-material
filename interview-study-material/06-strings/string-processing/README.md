# String Processing

Strings are just arrays of bytes/runes, so array patterns (two pointers, sliding window, frequency maps) all apply. The extra things interviewers watch for here:

- **Bytes vs runes.** `s[i]` indexes **bytes**; `for _, r := range s` iterates **runes** (Unicode code points). ASCII-only? bytes are fine. Unicode? use runes. Say which you assume.
- **Strings are immutable in Go.** Building a result char-by-char with `+` is O(n²) (each `+` copies). Use `strings.Builder` for O(n).

## Programs

### [Longest common prefix](longest-common-prefix/longest_common_prefix.go)

Find the longest prefix shared by all strings in a list.

**How to think about it:** start by *assuming the whole first string is the prefix*, then for each other string **shrink the candidate** until that string starts with it (`HasPrefix`). If it ever empties, there's no common prefix. This "guess big, shrink to fit" idea is simpler than comparing column-by-column and is easy to reason about.

**Dry run** — `["flower","flight","flow"]`: prefix `"flower"` → vs `"flight"` shrink to `"fl"` → vs `"flow"` already a prefix → `"fl"`.

- **O(S) time** where S = total characters, **O(1) extra space.**
- **Edge cases:** empty list → `""`; a string that's shorter than the prefix; no overlap → `""`.

### [String compression](string-compression/string_compression.go)

Run-length encode: `"aabcccccaaa"` → `"a2b1c5a3"`.

**How to think about it:** a single left-to-right scan comparing each char to the previous. While they match, bump a `count`; when they differ, **flush** the previous char and its count, then reset. Don't forget the **final flush** after the loop for the last run. Use `strings.Builder` so appends stay O(1) amortized.

**Dry run** — `"aabcccccaaa"`: `aa`→`a2`, `b`→`b1`, `ccccc`→`c5`, `aaa`→`a3` ⇒ `"a2b1c5a3"`.

- **O(n) time, O(n) space** (output buffer).
- **Real-world nuance:** the LeetCode variant only compresses if it actually shrinks the string (`"abc"` stays `"abc"`, not `"a1b1c1"`) — mention you'd add that guard.

### [Title case](title-case/title_case.go)

Capitalize the first letter of each word.

**How to think about it:** split on whitespace (`strings.Fields`), capitalize each word, join back with spaces. This one is about knowing the standard library.

- **Note:** it uses `strings.Title`, which is **deprecated** (mishandles some Unicode). The modern replacement is `golang.org/x/text/cases` (`cases.Title(language.Und)`). Flag this in an interview — awareness of deprecations reads as senior.
- **O(n) time, O(n) space.**

### [Pig Latin](pig-latin/pig_latin.go)

Transform each word: starts with a vowel → append `"yay"`; starts with consonant(s) → move the leading consonant cluster to the end and append `"ay"`, **preserving original casing** (ALL CAPS, Titlecase, lowercase).

**How to think about it:** decompose the problem — (1) a `pigLatinWord` transform on a lowercased word, (2) a **casing reapply** step so the output matches the input's style, (3) a sentence wrapper that splits, maps each word, and rejoins. The casing-preservation is the part most people forget; isolating it into helpers (`isAllUpper`, `isTitleCase`) keeps the logic clean.

**Approach for one word:** vowel-start → `word+"yay"`; else scan to the first vowel, then reassemble as `word[i:] + word[:i] + "ay"` (rotate the consonant cluster to the back).

- **O(n) per word, O(n) total.**
- **Edge cases:** words with no vowel; single-letter words; mixed casing.

## Review checklist

- When do you iterate by byte vs by rune, and why does it matter for Unicode?
- Why is `strings.Builder` O(n) where repeated `+` concatenation is O(n²)?
- In compression, why is the post-loop "final flush" necessary?
- What's deprecated about `strings.Title`, and what replaces it?
