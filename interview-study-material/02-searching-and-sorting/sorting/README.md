# Sorting

Sorting is the workhorse pre-processing step: once data is ordered, two pointers, binary search, greedy sweeps, and dedup all become easy. Know the two O(n log n) divide-and-conquer sorts cold, and know how to sort by a **custom key** in Go.

**The two algorithms compared (say this out loud in an interview):**

| | Merge sort | Quick sort |
|---|---|---|
| Idea | Split in half, sort each, **merge** | Pick pivot, **partition**, recurse on sides |
| Time | O(n log n) **guaranteed** | O(n log n) average, **O(n²) worst** (bad pivots) |
| Space | O(n) extra (merge buffer) | O(log n) stack, **in-place** |
| Stable? | **Yes** | No |
| Work happens... | on the way **up** (merge) | on the way **down** (partition) |

**How to remember which is which:** merge sort is *lazy on the way down, works on the way up*. Quick sort is *works on the way down (partition), trivial on the way up*.

## Programs

### [Merge sort](merge-sort/merge_sort.go)

Classic divide-and-conquer: split the slice in half, recursively sort each half, then merge two sorted halves.

**How to think about it:** the recursion is trivial — split until pieces are length ≤ 1 (a single element is already sorted). **All the real work is in `merge`**: walk two sorted slices with two pointers, always copying the smaller front element, then append whatever remains.

**Dry run of merge** on `[1,3,5]` + `[2,4,6]`: compare 1<2→take 1; 3>2→take 2; 3<4→take 3; 5>4→take 4; 5<6→take 5; left empty→append remaining `[6]` → `[1,2,3,4,5,6]`.

- **O(n log n) time** (log n levels × O(n) merge per level), **O(n) space**, **stable**.
- **When to prefer it:** you need stability, or worst-case guarantees, or you're sorting a linked list / external data.

### [Quick sort](quick-sort/quick_sort.go)

In-place sort using **Lomuto partition** (pivot = last element).

**How to think about it:** partition rearranges the array so everything `< pivot` is on the left and everything `≥ pivot` is on the right, then puts the pivot in its **final sorted position**. Now recurse on the two sides — the pivot never moves again. The `i` pointer marks "boundary of the less-than region"; whenever `arr[j] < pivot` you grow that region and swap the element in.

The file contains a full step-by-step **dry run** on `[10,7,8,9,1,5]` — trace it once and partitioning will click forever.

- **O(n log n) average, O(n²) worst** (already-sorted input with last-element pivot), **O(log n)** stack, **in-place**, not stable.
- **Mitigation to mention:** randomized or median-of-three pivot avoids the worst case on sorted input.

### [Custom field sort](custom-field-sort/custom_field_sort.go)

Sort a slice of structs by different fields using `sort.Slice` and a comparator closure.

**How to think about it:** you rarely reimplement sorting in real code — you supply a **less function** `func(i, j int) bool` that answers *"should element i come before element j?"* Return `people[i].Age < people[j].Age` to sort ascending by age; swap the field to sort by name.

- **Multi-key tip:** for tie-breaking, return the primary comparison first, and only when equal fall through to the secondary key. Use `sort.SliceStable` when equal elements must keep their original order.
- `sort.Slice` runs in **O(n log n)** (introsort under the hood).

## Review checklist

- Merge vs quick sort: which is stable, which is in-place, which has an O(n²) worst case and why?
- Where does the actual work happen in each (merge-up vs partition-down)?
- How do you write a comparator for "sort by age ascending, then name alphabetically"?
- When would you pick merge sort despite its O(n) extra space?
