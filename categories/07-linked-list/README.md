# 07 - Linked List

> Code: [`linked_list.go`](linked_list.go) - Tests: [`linked_list_test.go`](linked_list_test.go) - run `go test ./categories/07-linked-list/`

## Overview & mental model

A singly linked list is nodes each pointing to the next. **No random access** (indexing is O(n)), but **O(1) insert/delete** once you hold the node before the target. Three techniques dominate:

1. **Dummy head** - a throwaway node before `head` so inserting/deleting the first element needs no special case.
2. **Pointer reversal** - iteratively flip `next` pointers with a trailing `prev`.
3. **Fast/slow (tortoise-hare)** - two pointers at different speeds find the middle, detect cycles, or locate the nth-from-end in a single pass.

## How to recognize it

- Input is a `*ListNode`; the task is to reverse, merge, reorder, or find a positional node.
- "In one pass", "O(1) space", "detect a loop", "nth from the end" - fast/slow.
- "Remove/insert at the front conditionally" - reach for a dummy head.

## How to think / attack plan

1. Draw the nodes and arrows. Pointer bugs come from losing a reference - **save `next` before rewiring**.
2. Will an operation touch the head? If so, use a dummy to unify the logic.
3. Need a positional node without knowing the length? Use a **gap** between two pointers or fast/slow speeds.
4. Recursive vs iterative: recursion is elegant but O(n) stack; iterative is O(1) space.

## Core template - iterative reversal

```go
var prev *ListNode
for cur := head; cur != nil; {
    next := cur.Next
    cur.Next = prev
    prev, cur = cur, next
}
return prev
```

---

## Problems (easy -> hard)

### 1. Reverse Linked List (easy)
**Thinking.** Flip each pointer to face backward; `prev` accumulates the reversed prefix. Always cache `next` first or you lose the rest of the list.
**Dry run.** `1->2->3`: prev nil; after each step prev = 1, 2->1, 3->2->1.
**Complexity.** Time O(n), space O(1).

### 2. Linked List Cycle - detect (easy)
**Direction of thinking.** If fast moves 2x and slow 1x, inside a loop fast gains one step per iteration on slow and must eventually land on it (mod cycle length). If fast hits nil, no cycle.
**Complexity.** Time O(n), space O(1).

### 3. Cycle II - find the start (medium)
**Direction of thinking.** Floyd's math: distance from head to cycle start = distance from meeting point to cycle start (mod cycle length). So after they meet, reset one pointer to head and advance both at 1x; they meet at the entrance.
**Complexity.** Time O(n), space O(1).

### 4. Merge Two Sorted Lists (easy)
**Thinking.** Classic merge with a dummy tail; append the smaller head each step, then attach the leftover.
**Complexity.** Time O(n+m), space O(1).

### 5. Remove Nth Node From End (medium)
**Direction of thinking.** To find the nth-from-end in one pass, advance `fast` n steps first, then move `fast` and `slow` together; when `fast` reaches the end, `slow` sits just before the target. Dummy handles removing the head.
**Dry run.** `1..5, n=2`: gap 2; slow stops at node 3, unlink 4.
**Complexity.** Time O(n), space O(1).

### 6. Middle of the Linked List (easy)
**Thinking.** Fast/slow: when fast reaches the end, slow is at the middle (second middle for even length).
**Complexity.** Time O(n), space O(1).

### 7. Reorder List (medium)
**Direction of thinking.** Target pattern L0->Ln->L1->Ln-1... Decompose: (1) find the middle, (2) reverse the second half, (3) zip the two halves alternately. Each sub-step is a known technique.
**Dry run.** `1 2 3 4 5` -> halves `1 2 3` and reversed `5 4` -> `1 5 2 4 3`.
**Complexity.** Time O(n), space O(1).

### 8. Add Two Numbers (medium)
**Direction of thinking.** Digits are stored reversed, so you add from the head with a running **carry** - exactly like grade-school addition. Loop while either list or the carry remains.
**Dry run.** `2->4->3` (342) + `5->6->4` (465) = 807 -> `7->0->8`.
**Complexity.** Time O(max(n,m)), space O(max(n,m)).

### 9. Palindrome Linked List (medium)
**Direction of thinking.** For O(1) space: find the middle, reverse the second half, compare node-by-node with the first half. (Simpler O(n) space: dump to a slice and two-pointer it.)
**Complexity.** Time O(n), space O(1).

---

## Common pitfalls & edge cases

- Losing the rest of the list by rewiring before saving `next`.
- Not using a dummy when the head can change (remove/insert at front).
- Fast/slow loop conditions: `fast != nil && fast.Next != nil` (else nil-deref).
- Even vs odd length changes which node is "the middle".
- Creating an accidental cycle when reordering (forgetting to set a `next` to nil).

## Interview Q&A

- **Why a dummy head?** It removes special-casing for operations that affect the first node.
- **How does Floyd's cycle detection work / why O(1) space?** Two pointers at different speeds meet inside a loop; no extra structure needed (vs a hash set which is O(n) space).
- **Recursion vs iteration for reversal?** Same O(n) time; recursion uses O(n) stack, iteration O(1).
- **Array vs linked list?** Arrays: O(1) index, cache-friendly, O(n) middle insert. Lists: O(1) middle insert/delete given the node, O(n) access, pointer overhead.
