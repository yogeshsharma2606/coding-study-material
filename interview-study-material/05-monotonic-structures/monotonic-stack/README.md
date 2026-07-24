# Monotonic Stack

A **monotonic stack** keeps its elements sorted (always increasing or always decreasing) by **popping anything that would break the order** before pushing. That single discipline solves an entire family of problems in **O(n)**: "next/previous greater/smaller element", "span", "how many days until…", stock spans, histogram areas.

**Recognition cue (burn this in):** the phrase *"next greater / previous smaller / how far until a bigger value"* → monotonic stack. Brute force is O(n²) (for each element scan outward); the stack removes the rescanning because **each index is pushed and popped at most once**.

**The mental model:** store **indices** (not values) on the stack. When a new element arrives, it "resolves" every pending element it dominates — pop them and record their answer. Anything left unresolved waits for a future element (or never gets one → boundary answer).

## Programs

### [Stock span](stock-span/stock_span.go)

For each day, how many consecutive days (including today) had a price `≤` today's price, counting backward.

**How to think about it:** the span ends at the first earlier day with a **strictly greater** price. So keep a stack of indices with **decreasing** prices. For day `i`, pop all days whose price `≤ prices[i]` (they're engulfed by today), then: if the stack is empty, every prior day qualifies → span `= i+1`; otherwise span `= i − (index now on top)`.

**Dry run** — `[100,80,60,70,60,75,85]` → `[1,1,1,2,1,4,6]`. Day `75` pops `60,70,60` (all ≤75) leaving `80` on top at index 1 → span `5−1=4`.

- **O(n) time, O(n) space.**

### [Daily temperatures](daily-temperatures/daily_temperatures.go)

For each day, how many days until a **warmer** temperature (0 if none).

**How to think about it:** this is "next greater element" measured in distance. Keep a stack of indices of days **still waiting** for a warmer day (decreasing temperatures). When today `i` is warmer than the day on top, that day's wait is over → pop it and set `result[poppedIndex] = i − poppedIndex`. Keep popping (one warm day can resolve several pending colder days), then push `i`.

**Dry run** — `[73,74,75,71,69,72,76,73]` → `[1,1,4,2,1,1,0,0]`. When `72` arrives it resolves days `69` (dist 1) and `71` (dist 2); the never-resolved days `76` and last `73` stay 0.

- **O(n) time, O(n) space.** Indices left on the stack at the end have no warmer future day → their result stays 0.

## Review checklist

- Why does "each index pushed/popped once" make these O(n) instead of O(n²)?
- Why store indices rather than values on the stack?
- Increasing vs decreasing stack: how do you decide which, given "next greater" vs "next smaller"?
- What do the indices *left on the stack* at the end represent?
