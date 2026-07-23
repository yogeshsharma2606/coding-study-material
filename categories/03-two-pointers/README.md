# 03 - Two Pointers

> Code: [`two_pointers.go`](two_pointers.go) - Tests: [`two_pointers_test.go`](two_pointers_test.go) - run `go test ./categories/03-two-pointers/`

## Overview & mental model

Two pointers turn an O(n^2) nested loop into an O(n) single pass by maintaining two indices with a rule for moving them. Two flavors:

- **Converging** (opposite ends): start `l=0, r=n-1`. Works on **sorted** arrays or symmetric structures (palindromes). Sortedness tells you which pointer to move.
- **Slow/Fast** (same direction): `slow` marks a write/boundary position, `fast` scans ahead. Used for in-place filtering/dedup.

## How to recognize it

- Array is (or can be) **sorted** and you need a pair/triplet with a target relationship.
- "Find a pair that...", "remove in place", "container/area between two lines", palindrome checks.
- You wrote a double loop and the inner loop just searches - two pointers can often replace it.

## How to think / attack plan

1. If unsorted and order doesn't matter for the answer, **sort first** (O(n log n)) to unlock converging pointers.
2. Define the **move rule**: what does the current sum/comparison tell you about which side to advance? (Too small -> move left up; too big -> move right down.)
3. For triplet/k-sum, **fix one element** and two-pointer the rest; skip duplicates to avoid repeated results.

## Core template - converging pointers on sorted data

```go
l, r := 0, len(a)-1
for l < r {
    if cond(a[l], a[r]) { /* record / return */ }
    if tooSmall { l++ } else { r-- }
}
```

---

## Problems (easy -> hard)

### 1. Two Sum II - sorted (easy)
**Thinking.** Sorted input means: if `a[l]+a[r]` is too small, only increasing helps (move `l`); too big, move `r`. Each step eliminates one candidate, so O(n). No hash map needed - O(1) space.
**Dry run.** `[2,7,11,15], t=9`: 2+15=17>9 -> r--; 2+11=13>9 -> r--; 2+7=9 -> [1,2].
**Complexity.** Time O(n), space O(1).

### 2. Valid Palindrome II (easy-medium)
**Problem.** Palindrome allowed after deleting at most one char.
**Direction of thinking.** Converge until the first mismatch. At that point you must "use" your one deletion - try skipping the left char OR the right char and check if either remainder is a palindrome.
**Why both.** You can't know in advance which side's char is the intruder.
**Complexity.** Time O(n), space O(1).

### 3. Container With Most Water (medium)
**Problem.** Two lines form a container; maximize water = `width * min(heights)`.
**Direction of thinking.** Start at maximum width. The area is bounded by the **shorter** wall, so moving the taller wall inward can never increase the area (width shrinks, height still capped). Therefore always move the shorter wall - this greedily explores only promising states.
**Why it's correct.** Discarding the shorter wall loses no better solution: any container using it is narrower and no taller.
**Dry run.** `[1,8,6,2,5,4,8,3,7]` -> best 49 (between the 8 at idx1 and 7 at idx8).
**Complexity.** Time O(n), space O(1).

### 4. 3Sum (medium)
**Problem.** All unique triplets summing to 0.
**Direction of thinking.** Brute force is O(n^3). Sort, then **fix an anchor `i`** and two-pointer the rest for `-nums[i]` in O(n) -> O(n^2) total. Sorting also makes **deduping** easy: skip equal anchors and equal l/r values after a hit.
**Dry run.** Sorted `[-4,-1,-1,0,1,2]`; anchor -1 finds (-1,-1,2) and (-1,0,1).
**Complexity.** Time O(n^2), space O(1) extra (or O(n) for output).

### 5. 3Sum Closest (medium)
**Same structure** as 3Sum but track the sum with the minimum `|sum-target|`. Move pointers by comparing `sum` to `target`.
**Complexity.** Time O(n^2), space O(1).

### 6. Remove Duplicates from Sorted Array (easy)
**Problem.** Dedup in place, return new length.
**Direction of thinking.** Because it's sorted, duplicates are adjacent. `slow` holds the last unique; when `fast` sees something new, write it just after `slow`.
**Complexity.** Time O(n), space O(1).

### 7. Squares of a Sorted Array (easy-medium)
**Problem.** Return sorted squares of a sorted array with negatives.
**Direction of thinking.** After squaring, the largest values are at the **ends** (most negative or most positive). Compare the two ends and fill the output **from the back**. Beats squaring then re-sorting (O(n log n)).
**Dry run.** `[-4,-1,0,3,10]` -> compare 16 vs 100 -> place 100, ... -> `[0,1,9,16,100]`.
**Complexity.** Time O(n), space O(n) for output.

### 8. Trapping Rain Water (hard)
**Problem.** Given elevation bars, compute trapped water.
**Direction of thinking.** Water over column i = `min(maxLeft_i, maxRight_i) - height[i]`. The precompute-arrays solution is O(n) time/space. The two-pointer trick removes the extra arrays: keep `leftMax`/`rightMax`; whichever side has the **smaller** running max is the binding constraint, so process that side (you know its bound is final).
**Why moving the smaller side is safe.** If `leftMax < rightMax`, then for the left pointer the min is `leftMax` regardless of future right bars, so its water is determined now.
**Dry run.** `[0,1,0,2,1,0,1,3,2,1,2,1]` -> 6 units.
**Complexity.** Time O(n), space O(1).

---

## Common pitfalls & edge cases

- Forgetting to skip duplicates in 3Sum -> repeated triplets.
- Moving the wrong pointer in Container With Most Water (must move the shorter).
- Not handling empty input in Trap/RemoveDuplicates.
- Converging pointers require **sorted** input - sorting changes indices (bad if the answer needs original indices; then use hashing instead).

## Interview Q&A

- **When two pointers vs hashing for pair-sum?** Sorted or space-constrained -> two pointers (O(1) space). Need original indices or unsorted-and-can't-sort -> hashing.
- **Why move the shorter wall in Container?** The area is capped by the shorter wall; moving the taller one can't help.
- **How do you generalize to k-Sum?** Recursively fix elements down to 2, then two-pointer the base case: O(n^{k-1}).
