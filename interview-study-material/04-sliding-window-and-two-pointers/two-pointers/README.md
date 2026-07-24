# Two Pointers

Two pointers replace a nested loop with a single coordinated scan. The most powerful variant is **converging pointers**: start at both ends and move inward, using a rule that lets you *discard possibilities without checking them*. That's how these O(n²) brute forces become **O(n)**.

**Recognition cue:** sorted array + "find a pair", or *"area/water between two lines"*, or *"palindrome check"*, or anything where you can argue *"moving the worse pointer inward can only help"*.

## Programs

### [Container with most water](container-with-most-water/container_with_most_water.go)

Given bar heights, pick two lines forming a container holding the most water: `area = min(h[l], h[r]) * (r - l)`.

**The key insight (the exchange argument):** area is limited by the **shorter** wall. If you move the *taller* wall inward, width shrinks and the height is still capped by the same short wall → area can only drop. So the only move that could ever help is to **advance the shorter wall**, hoping for a taller one. That justification is why greedily moving the shorter side inward finds the optimum in one pass.

**Dry run** — `[1,8,6,2,5,4,8,3,7]`:

```
l=0(1) r=8(7): area=min(1,7)*8=8 ; move l (shorter)
l=1(8) r=8(7): area=min(8,7)*7=49; move r (shorter)  <- max
... nothing beats 49
```

Answer **49**.

- **O(n) time, O(1) space.**

### [Trapping rain water](trapping-rain-water/trapping_rain_water.go)

How much water is trapped between bars after rain.

**The insight:** water above any bar = `min(highestWallToLeft, highestWallToRight) − thisBarHeight`. The two-pointer trick avoids precomputing both max arrays: process from whichever side currently has the **smaller** boundary, because that side's answer is *fully determined* by its running max — the opposite (taller) side guarantees the `min` is on the side you're processing. Advance that pointer, updating its running max or adding trapped water.

**Dry run** — `[4,2,0,3,2,5]`: pointers close in; at the low middle bars water accumulates `2 + 4 + 1 + 2 = 9`. Answer **9**. (The file's comment also traces the classic `[0,1,0,2,1,0,1,3,2,1,2,1]` = 6.)

- **O(n) time, O(1) space.** The array-based version (precompute `leftMax[]` and `rightMax[]`) is O(n) time but O(n) space — a good comparison to state.
- **Why process the smaller side?** Because `min(leftMax, rightMax)` is then pinned to the side you're on, so `leftMax − height[left]` (or the mirror) is provably correct.

## Review checklist

- State the exchange argument for why moving the shorter wall is the only useful move (container problem).
- In trapping rain water, why is it always safe to process the side with the smaller boundary?
- Two-pointer O(1)-space vs precomputed-max O(n)-space: name the trade-off.
- What property (sorted / bounded-by-ends) makes converging pointers valid for a problem?
