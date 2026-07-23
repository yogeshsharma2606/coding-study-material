# 20 - Intervals

> Code: [`intervals.go`](intervals.go) - Tests: [`intervals_test.go`](intervals_test.go) - run `go test ./categories/20-intervals/`

## Overview & mental model

Interval problems almost always start with **sorting**, then a single **left-to-right sweep** with a simple rule:

- Sort by **start** to **merge/insert** overlapping intervals.
- Sort by **end** for **greedy scheduling** (max non-overlapping, min removals, min arrows).
- Use a **sweep line** over start (+1) and end (-1) events to count **concurrency** (rooms needed).

The overlap test you must know cold: intervals `[a,b]` and `[c,d]` overlap iff `a <= d && c <= b`.

## How to recognize it

- Input is a list of `[start, end]` pairs.
- "Merge", "insert", "overlap", "meeting rooms", "minimum removals/arrows", "free time".
- The answer depends on how ranges relate on a timeline.

## How to think / attack plan

1. Decide the sort key. Merging/inserting -> by **start**. Greedy "keep the most that fit" / "cover all" -> by **end**.
2. Sweep once, maintaining the current merged interval or the last kept end.
3. Counting simultaneous intervals -> split into start/end events and sweep (or use a min-heap of end times).
4. Two sorted lists to combine -> two pointers.

---

## Problems (easy -> hard)

### 1. Merge Intervals (medium)
**Direction of thinking.** Sort by start. Walk through; if the current interval starts within the last merged one, extend the end; otherwise start a new merged interval.
**Dry run.** `[1,3],[2,6],[8,10]` -> `[1,6],[8,10]`.
**Complexity.** Time O(n log n), space O(n).

### 2. Insert Interval (medium)
**Direction of thinking.** The list is already sorted and non-overlapping. Three phases: copy intervals entirely **before** the new one, **merge** the overlapping run into the new interval, then copy the ones entirely **after**. No re-sort needed.
**Complexity.** Time O(n), space O(n).

### 3. Non-overlapping Intervals - min removals (medium)
**Direction of thinking.** This is the classic **activity selection**: to keep the most non-overlapping intervals, always keep the one that **ends earliest** (it leaves the most room). Sort by end; remove any interval that starts before the kept end.
**Why by end, not start.** Ending earliest maximizes remaining space for future intervals - provable by exchange argument.
**Complexity.** Time O(n log n), space O(1).

### 4. Meeting Rooms - can attend all? (easy)
**Direction of thinking.** Sort by start; if any meeting starts before the previous ends, there's a conflict.
**Complexity.** Time O(n log n).

### 5. Meeting Rooms II - min rooms (medium)
**Direction of thinking.** Rooms needed = **maximum number of concurrent meetings**. Sweep sorted starts and ends: a start claims a room, an end frees one; track the peak. (Equivalent: a min-heap of end times, size = rooms.)
**Dry run.** `[0,30],[5,10],[15,20]` -> at time 5, two concurrent -> 2 rooms.
**Complexity.** Time O(n log n), space O(n).

### 6. Interval List Intersections (medium)
**Direction of thinking.** Both lists are sorted. The intersection of the two current intervals is `[max(starts), min(ends)]` if valid. Advance the pointer of whichever interval **ends first** (it can't intersect anything later).
**Complexity.** Time O(m+n), space O(m+n).

### 7. Minimum Number of Arrows to Burst Balloons (medium)
**Direction of thinking.** Same greedy as #3: sort by **end**, shoot an arrow at the first balloon's end (bursting all overlapping ones), and only add a new arrow when a balloon starts after the current arrow's position.
**Complexity.** Time O(n log n), space O(1).

---

## Common pitfalls & edge cases

- Sorting by the wrong key (start vs end) - it changes correctness for greedy problems.
- Touching endpoints: is `[1,4]` and `[4,5]` an overlap? Clarify inclusive vs exclusive with the interviewer (this repo treats `<=` as overlap in Merge).
- Mutating shared slices while merging (aliasing the same backing array).
- Empty input and single-interval cases.

## Interview Q&A

- **Why sort by end for scheduling?** Finishing earliest leaves maximum room for subsequent intervals - the greedy-choice property behind activity selection.
- **How to compute max concurrency?** Sweep start/end events (+1/-1) or use a min-heap of end times.
- **Overlap condition?** `[a,b]` and `[c,d]` overlap iff `a <= d && c <= b`.
- **Merge vs insert complexity?** Merge is O(n log n) (needs a sort); insert into an already-sorted list is O(n).
