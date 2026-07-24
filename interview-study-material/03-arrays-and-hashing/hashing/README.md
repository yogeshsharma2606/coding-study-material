# Array Hashing

Hashing's core deal: **trade space for time.** A hash map/set turns "does this value exist / where is it?" from an O(n) scan into an **O(1) average** lookup. The instant a problem needs *"have I seen X before?"*, *"pair that sums to target"*, *"group by key"*, or *"dedup"*, think hash map.

**The recognition reflex:** whenever your brute force is "for each element, scan the rest to find a match", a hash map usually collapses it from O(n²) to O(n) — you remember what you've seen instead of re-searching.

## Programs

### [Two Sum](two-sum/two_sum.go)

Return indices of the two numbers that add up to `target`.

**How to think about it:** brute force checks every pair — O(n²). The redundancy: for each `num` you re-scan for its partner `target − num`. Instead, **remember every number you've passed in a map (value → index)**. For the current `num`, if its complement is already in the map, you've found the pair in one pass. This is the archetype for "hashing removes the inner loop".

**Dry run** — `[2,7,11,15], target=11`:

```
num=2  complement=9  not seen -> store {2:0}
num=7  complement=4  not seen -> store {2:0, 7:1}
... wait target=11: at num=2 complement=9; at num=7 complement=4 (not 2)
```

(For `target=9`: `2`→need 7 (no), store; `7`→need 2 (yes!) → return `[0,1]`.)

- **O(n) time, O(n) space.** One pass; store *after* checking so you don't match an element with itself.

### [Longest consecutive sequence](longest-consecutive-sequence/longest_consecutive_sequence.go)

Length of the longest run of consecutive integers (e.g. `[100,4,200,1,3,2]` → `4` for `1,2,3,4`).

**How to think about it (shipped approach — sort):** once sorted, consecutive numbers sit next to each other, so you just walk once, extend the streak when `nums[i] == nums[i-1]+1`, skip duplicates, and reset on a gap.

**Dry run** — sorted `[1,2,3,4,100,200]`: streak grows 1→2→3→4 across `1..4`, breaks at 100, so longest = **4**.

- **This version: O(n log n) time** (dominated by the sort), **O(1) extra space.**
- **The O(n) upgrade to mention (classic hash-set trick):** put all numbers in a set; only start counting a streak from a number whose predecessor `x−1` is **absent** (i.e. a true sequence start), then walk `x+1, x+2, ...` in the set. Each number is visited at most twice → **O(n) time, O(n) space**. Knowing *why sorting is O(n log n) but the set is O(n)* is the senior-level point.

## Review checklist

- Why does a hash map turn Two Sum from O(n²) into O(n), and why store *after* checking?
- For longest consecutive sequence, contrast the sort (O(n log n)) vs hash-set (O(n)) approaches.
- In the hash-set version, why only start a streak when `x-1` is not in the set?
- What's the average vs worst-case cost of a hash lookup, and what causes the worst case?
