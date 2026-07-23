// Package intervals contains interval interview problems.
//
// Interval problems are almost always solved by SORTING first - by start time
// (to merge/insert) or by end time (to schedule the most non-overlapping) - then
// SWEEPING left to right with a simple rule. The overlap test for [a,b] and
// [c,d] is: a <= d && c <= b. A "sweep line" over start/end events counts
// concurrency (e.g., rooms needed).
package intervals

import "sort"

// Merge combines all overlapping intervals. Sort by start; extend the current
// merged interval while the next one overlaps, else push a new one.
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := [][]int{intervals[0]}
	for _, cur := range intervals[1:] {
		last := res[len(res)-1]
		if cur[0] <= last[1] { // overlap
			if cur[1] > last[1] {
				last[1] = cur[1] // extend end
			}
		} else {
			res = append(res, cur)
		}
	}
	return res
}

// Insert adds newInterval into a sorted, non-overlapping list and merges.
// Three phases: intervals entirely before, the overlapping run (merged), and
// intervals entirely after.
func Insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	i, n := 0, len(intervals)
	// before: end < newStart
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	// merge overlapping: start <= newEnd
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	res = append(res, newInterval)
	// after
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}

// EraseOverlapIntervals returns the min removals to make intervals
// non-overlapping. Greedy: sort by END; always keep the interval that finishes
// earliest, removing any that starts before it ends.
func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	removals := 0
	prevEnd := intervals[0][1]
	for _, cur := range intervals[1:] {
		if cur[0] < prevEnd {
			removals++ // overlaps the kept one -> remove cur
		} else {
			prevEnd = cur[1]
		}
	}
	return removals
}

// CanAttendMeetings reports whether a person can attend all meetings (no overlap).
func CanAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

// MinMeetingRooms returns the minimum rooms needed = max concurrent meetings.
// Sweep line: sort starts and ends separately; a start needs a room, an end
// frees one. The peak occupancy is the answer.
func MinMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, iv := range intervals {
		starts[i] = iv[0]
		ends[i] = iv[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)
	rooms, maxRooms, e := 0, 0, 0
	for s := 0; s < n; s++ {
		for e < n && ends[e] <= starts[s] {
			rooms-- // a meeting ended before this one starts
			e++
		}
		rooms++
		if rooms > maxRooms {
			maxRooms = rooms
		}
	}
	return maxRooms
}

// IntervalIntersection returns intersections of two sorted interval lists.
// Two pointers: the intersection is [max(starts), min(ends)] when it's valid;
// advance whichever interval ends first.
func IntervalIntersection(a, b [][]int) [][]int {
	var res [][]int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		lo := max(a[i][0], b[j][0])
		hi := min(a[i][1], b[j][1])
		if lo <= hi {
			res = append(res, []int{lo, hi})
		}
		if a[i][1] < b[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

// FindMinArrowShots: min arrows to burst all balloons (an arrow at x bursts any
// balloon spanning x). Sort by END; shoot at each kept balloon's end, bursting
// all overlapping ones.
func FindMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	arrows := 1
	end := points[0][1]
	for _, p := range points[1:] {
		if p[0] > end { // no overlap with current arrow
			arrows++
			end = p[1]
		}
	}
	return arrows
}
