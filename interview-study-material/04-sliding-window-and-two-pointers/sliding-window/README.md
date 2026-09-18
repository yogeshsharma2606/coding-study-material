# Sliding Window

A sliding window keeps a **contiguous range `[left, right]`** and slides it across the array/string, maintaining just enough state (a sum, a count, a char-frequency map, a deque) so each element enters and leaves the window **once**. That's how "try every subarray" (O(n²)) becomes **O(n)**.

**Two flavors — know which you're in:**
- **Fixed window** (size `k` given): `right` advances every step; `left` follows exactly `k` behind. You *add the entering element and evict the leaving one*.
- **Variable window** (grow/shrink to satisfy a constraint): `right` always expands; `left` advances **only when the window becomes invalid** (e.g. a duplicate appears, sum exceeds target).

**Recognition cue:** *"contiguous subarray/substring"* + *"longest/shortest/max/min/count with property P"* → sliding window.

## Programs

### [First negative in every window](first-negative-in-window/first_negative_in_window.go)

For each fixed window of size `k`, report the first negative number (or 0 if none).

**How to think about it:** you only care about **negatives, in order of appearance**. Keep a queue of the *indices* of negatives currently inside the window. Two maintenance steps each iteration: (1) **evict** the front index if it slid out of the window (`indexes[0] <= i-k`), (2) **enqueue** the current index if `arr[i] < 0`. Once the first full window is formed (`i >= k-1`), the answer is `arr[indexes[0]]` (front) or 0 if empty.

**Dry run** — `[12,-1,-7,8,-15,30,16,28], k=3` → `[-1,-1,-7,-15,-15,0]`. (Last window `30,16,28` has no negative → 0.)

- **O(n) time, O(k) space.** Each index is pushed/popped at most once.

### [Sliding-window maximum](sliding-window-maximum/sliding_window_maximum.go)

Report the maximum of every window of size `k`. Uses a **monotonic deque**.

**The insight (the reason this is O(n), not O(nk)):** keep a deque of indices whose values are **decreasing** front→back. Before adding `i`, pop from the **back** every index whose value is `≤ nums[i]` — they can never be a max while `nums[i]` is in the window, so they're dead weight. Also pop from the **front** any index that slid out of the window. The front of the deque is *always* the current window's max.

**Dry run** — `[1,3,-1,-3,5,3,6,7], k=3` → `[3,3,5,5,6,7]`. When `5` arrives it wipes out `-1,-3` from the back and eventually the older maxima, so the front stays the true max cheaply.

- **O(n) time, O(k) space.** Each index enters and leaves the deque once (amortized O(1) per step).
- **Why not a heap?** A max-heap gives O(n log k); the monotonic deque beats it at O(n). See also category 05 (monotonic structures).

### [Longest substring without repeating characters](longest-substring-without-repeats/longest_substring_without_repeats.go)

Length of the longest substring with all-distinct characters — a **variable** window.

**How to think about it:** expand `right` one char at a time. Keep a map `char -> last index seen`. When the new char was already seen **inside the current window** (`lastIndex >= left`), you can't keep both, so **jump `left` to just past the previous occurrence**. The window `[left, right]` is always duplicate-free, so track its max length.

**Dry run** — `"pwwkew"`:

```
p: window "p"        len1
w: window "pw"       len2
w: dup! left->2      window "w"    len1
k: "wk"  len2
e: "wke" len3  <- max
w: dup (last w at 2 >= left 2) left->3  "kew" len3
```

Answer **3** (`"wke"`).

- **O(n) time, O(min(n, alphabet)) space.** `left` only moves forward, never resets — that's what keeps it linear.
- **Gotcha:** the `lastIndex >= left` check is essential; a stale occurrence *outside* the window must not drag `left` backward.

### [Longest Repeating Character Replacement](longest-repeating-character-replacement/character_replacement.go)

Length of the longest substring that can be converted into the **same character** using at most `k` replacements — a **variable sliding window**.

**How to think about it:** expand `right` one character at a time and maintain a frequency map `char -> count` for the current window.

The character with the highest frequency is the character we keep. Every other character needs to be replaced.

The key formula is:

```text
replacements = windowSize - maxFreq
```

If `replacements <= k`, the current window is valid.

If `replacements > k`, shrink the window from the left until it becomes valid again.

**Dry run** — `s = "AABABBA", k = 1`:

```text
A:       "A"       maxFreq=1  replacements=0  len=1
A:       "AA"      maxFreq=2  replacements=0  len=2
B:       "AAB"     maxFreq=2  replacements=1  len=3
A:       "AABA"    maxFreq=3  replacements=1  len=4  <- max

B:       "AABAB"   maxFreq=3  replacements=2  > k
         shrink -> "ABAB"

B:       "ABABB"   maxFreq=3  replacements=2  > k
         shrink -> "BABB"

A:       "BABBA"   maxFreq=3  replacements=2  > k
         shrink -> "ABBA"
```

Answer **4** — substring `"AABA"` can become `"AAAA"` with 1 replacement.

### Complexity

* **Time:** `O(n)`
* **Space:** `O(alphabet)` — effectively `O(1)` for a fixed alphabet.

`right` moves forward once and `left` also only moves forward, so every character is processed a constant number of times.

### Important Insight

We don't need to explicitly determine which characters to replace.

For example:

```text
Window = "AABA"

A appears 3 times
B appears 1 time

windowSize = 4
maxFreq    = 3

replacements = 4 - 3
             = 1
```

Therefore, with `k = 1`, `"AABA"` is valid because we can replace `B` with `A`.

### Gotcha

`maxFreq` does **not need to decrease** when the left side of the window moves.

It may become stale, but this still produces the correct maximum-length answer and keeps the algorithm `O(n)`.

### Pattern

**Sliding Window + Frequency Map**

This pattern is useful when:

* We need the longest/shortest substring.
* The window has a constraint.
* We need to track character frequencies.
* We can expand and shrink the window based on a calculated condition.



## Review checklist

- Fixed vs variable window: how does `left` move in each?
- Why does the monotonic deque give O(n) window-max where a heap gives O(n log k)?
- In "longest substring", why must `left` only ever move forward, and why the `>= left` guard?
- What state does each window carry (sum / frequency map / deque), and why is O(n) achievable?
