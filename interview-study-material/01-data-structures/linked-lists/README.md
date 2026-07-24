# Linked Lists

A linked list stores each value in its own node and connects nodes with pointers. Unlike an array, elements are **not** contiguous in memory, so you trade O(1) random access for O(1) insert/delete once you already hold the right node.

## When to reach for a linked list

- You insert/delete a lot in the **middle** and rarely index by position.
- You need a structure that grows without reallocating/copying (no capacity doubling like slices).
- The problem screams **"pointer surgery"**: reverse, merge, detect a cycle, remove the k-th node.

**The mental model that never fails:** almost every linked-list bug is a lost pointer. Before you overwrite `curr.next`, ask *"do I still have a handle on the node it used to point to?"* If not, save it in a temp first.

## Programs

### [Singly linked list](singly-linked-list/singly_linked_list.go)

Insert-at-tail, display, search, delete-by-value, and reverse for a one-directional list.

**How to think about it:** each node knows only its successor (`next`). So every operation is "walk from `head` until a condition, then relink." The only hard one is reverse.

**Reverse — the approach & why it works:** you need to flip every `next` arrow to point backward. If you naively do `curr.next = prev`, you lose the rest of the list. So the invariant is a **three-pointer dance**: remember `next` before you clobber the link, flip the link, then slide `prev` and `curr` forward one step.

```
prev=nil  curr=10 -> 20 -> 30 -> nil
save next=20; 10.next=nil;      prev=10, curr=20
save next=30; 20.next=10;       prev=20, curr=30
save next=nil;30.next=20;       prev=30, curr=nil  (stop)
head = prev = 30 -> 20 -> 10 -> nil
```

- **Delete** covers the head-is-the-target case separately, because head has no predecessor to relink.
- Insert/search/delete/reverse are all **O(n) time, O(1) space** (reverse is iterative, no recursion stack).
- **Edge cases:** empty list (`head == nil`), deleting the head, deleting a value that isn't present.

### [Doubly linked list](doubly-linked-list/doubly_linked_list.go)

Same operations plus a `prev` pointer, enabling backward traversal and O(1) delete when you already hold the node.

**How to think about it:** every node now has **two** arrows (`prev` and `next`), so every relink must fix **both directions**. The payoff: `Delete` doesn't need to track the predecessor separately — the node already knows its `prev`.

**Reverse trick:** for each node just swap its `prev` and `next`, then move along what *used to be* `next` (now `prev`). Finally reset `head` to the last processed node.

- Delete is O(1) *given the node*; finding it is still O(n).
- **Gotcha:** when deleting, update `curr.prev.next` **and** `curr.next.prev`, and guard both against `nil` (head/tail cases).

### [Floyd cycle detection](floyd-cycle-detection/floyd_cycle_detection.go)

Detects whether a list loops back on itself using two pointers moving at different speeds ("tortoise and hare").

**Direction of thinking:** "Is there a cycle?" with **O(1) extra space** is the signal. A hash set of visited nodes works but costs O(n) space. The insight: if a fast pointer moves 2 steps and a slow pointer moves 1 step, then *inside a loop the fast one gains 1 step on the slow one every iteration*, so it must eventually land on it — like two runners on a circular track. If there's no loop, fast simply falls off the end (`nil`).

**Dry run** on `1->2->3->4->(back to 1)`:

```
slow=1 fast=1
slow=2 fast=3
slow=3 fast=1   (fast wrapped around)
slow=4 fast=3
slow=1 fast=1   -> slow == fast -> cycle found
```

- **O(n) time, O(1) space.** Loop guard is `fast != nil && fast.next != nil` because you dereference `fast.next.next`.
- To find the *start* of the cycle, after they meet reset one pointer to head and advance both by 1 — they meet at the entry (Floyd's second phase).

### [Merge two sorted lists](merge-two-sorted-lists/merge_two_sorted_lists.go)

Merges two already-sorted lists into one sorted list by splicing nodes (no new nodes allocated).

**Direction of thinking:** two sorted inputs → **two pointers**, always take the smaller head. The classic trick is a **dummy head node**: it removes the "is this the first node?" special case, so the loop body is uniform. You build the result by pointing `curr.next` at whichever input head is smaller, then advancing that input and `curr`.

**Why the tail step is cheap:** when one list runs out, the other is already sorted, so you attach the whole remainder in one link — no more comparisons.

- **O(n + m) time, O(1) extra space** (splices existing nodes; the only allocation is the dummy).
- **Edge cases:** one or both lists empty (the dummy + tail-attach handles them for free).

## Review checklist

- Can you reverse a singly linked list from memory using exactly three pointers?
- Why does Floyd's algorithm guarantee the pointers meet inside a loop?
- Why does a dummy head simplify merge/insert/delete code?
- Which operations are O(1) on a doubly linked list that are O(n) on a singly linked one — and why?
