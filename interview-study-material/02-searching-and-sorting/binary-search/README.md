# Binary Search

Binary search is the art of **halving the search space every step**. If you can answer *"is the answer to the left or the right of here?"* in O(1), you find the answer in **O(log n)**.

**Recognition cues (the "hint decoder"):**
- The input is **sorted** (or can be sorted / is monotonic in some property).
- Constraints are huge (`n` up to 10^9-10^18) and brute force is too slow → you need O(log n).
- The phrase *"minimize the maximum"* / *"smallest value that works"* → **binary search on the answer**.

**The one invariant to never forget:** decide your loop condition and boundary updates so the search space **strictly shrinks** every iteration, or you infinite-loop. The standard, bug-resistant form is `l <= r` with `l = mid+1` / `r = mid-1`.

**Why `mid := l + (r-l)/2` and not `(l+r)/2`?** In languages with fixed-width ints, `l+r` can overflow. The subtraction form can't. Say it in the interview — it's a free signal of maturity.

## Programs

### [Binary search](binary-search/binary_search.go)

Find a target's index in a sorted array.

**How to think about it:** look at the middle. If it's the target, done. If the middle is **too small**, the answer can only be to the **right**, so discard the left half (`l = mid+1`). If too big, discard the right half (`r = mid-1`). Repeat on the surviving half.

**Dry run** — `arr = [1..10]`, target `7`:

```
l=0 r=9  mid=4 arr[4]=5 <7 -> l=5
l=5 r=9  mid=7 arr[7]=8 >7 -> r=6
l=5 r=6  mid=5 arr[5]=6 <7 -> l=6
l=6 r=6  mid=6 arr[6]=7 == target -> return 6
```

- **O(log n) time, O(1) space.** Returns `-1` when not found (loop ends with `l > r`).

### [Search a rotated sorted array](search-rotated-sorted-array/search_rotated_sorted_array.go)

A sorted array was rotated at an unknown pivot (e.g. `[4,5,6,7,0,1,2]`). Find a target in O(log n).

**Direction of thinking:** you've lost global sortedness, but here's the key insight: **when you split a rotated sorted array at the middle, at least one half is still perfectly sorted.** So each step: (1) figure out which half is sorted (compare `nums[left]` to `nums[mid]`), (2) check if the target lies **within that sorted half's range** — if yes, search there; if no, search the other half. You still halve every step → still O(log n).

**Dry run** — `nums = [4,5,6,7,0,1,2]`, target `0`:

```
l=0 r=6 mid=3 nums[3]=7. left half [4..7] sorted (4<=7).
   is 0 in [4,7)? no -> go right: l=4
l=4 r=6 mid=5 nums[5]=1. left half [0..1] sorted (0<=1).
   is 0 in [0,1)? yes -> go left: r=4
l=4 r=4 mid=4 nums[4]=0 == target -> return 4
```

- **O(log n) time, O(1) space.**
- **Edge cases:** `nums[left] <= nums[mid]` uses `<=` so a two-element window where left==mid is treated as sorted; duplicates can degrade this to O(n) (worth mentioning).

## Review checklist

- Why `l + (r-l)/2` instead of `(l+r)/2`?
- In a rotated array, how do you know which half is sorted, and why does one half always is?
- What loop-termination condition avoids both infinite loops and off-by-one misses?
- When would you "binary search on the answer" instead of on an array index?
