# Stack and Queue

Two dual structures that differ only in *which end you remove from*:

- **Stack = LIFO** (Last In, First Out). Push/pop the **same** end. Think "undo history", call stack, matching brackets.
- **Queue = FIFO** (First In, First Out). Add to one end, remove from the other. Think "waiting line", BFS frontier.

**Recognition cue:** the moment a problem involves *"most recent unmatched thing"*, *"nesting"*, or *"reverse order"*, reach for a **stack**. When it's *"process in arrival order"* or *"level by level"*, reach for a **queue**.

## Programs

### [Stack implementation](stack-implementation/stack.go)

A stack of ints backed by a Go slice: `Push`, `Pop`, `Peek`.

**How to think about it:** the *end* of a slice is a perfect stack top — `append` grows it and reslicing `s[:len-1]` shrinks it, both amortized O(1). No need for a linked list.

- `Push` = `append`. `Pop` = read last element, then drop it. `Peek` = read last element without removing.
- **Design choice to notice:** `Pop`/`Peek` return `-1` on empty. That's a *sentinel*; in production you'd return `(int, bool)` or an error so `-1` isn't mistaken for real data. Interviewers love this discussion.
- All operations **O(1) amortized, O(n) space**.

### [Queue implementation](queue-implementation/queue.go)

A queue of ints: `Enqueue` (add at back), `Dequeue` (remove from front), `Front`.

**How to think about it:** add at the tail with `append`, remove from the head with `q.items[1:]`.

- **Gotcha worth saying out loud:** `q.items[1:]` reslices but the underlying array's front elements are never reclaimed, so a long-lived queue can leak memory. Fixes: a **ring buffer** (head/tail indices) or `container/list`. This is the senior-level nuance.
- Enqueue is O(1) amortized; this simple Dequeue is O(1) time but not memory-tidy.

### [Valid parentheses](valid-parentheses/valid_parentheses.go)

Given a string of `()[]{}`, decide whether every bracket is closed by the correct type in the correct order.

**Direction of thinking:** "correctly nested / matching pairs" is the canonical **stack** signal. The key realization: a closing bracket must match the **most recently opened** unmatched bracket — that's LIFO exactly. Push opens; on a close, the top must be its partner.

**Approach:** keep a map `close -> open`. For each char: if it opens, push it; if it closes, the stack must be non-empty **and** its top must equal the expected open — otherwise fail immediately. At the end the stack must be **empty** (no unclosed opens left).

**Dry run** on `([)]`:

```
'(' push        stack=[ (
'[' push        stack=[ ( [
')' expects '(' but top is '[' -> return false
```

Dry run on `{[]}` → push `{`, push `[`, `]` matches `[` (pop), `}` matches `{` (pop), stack empty → **true**.

- **O(n) time, O(n) space** (worst case all opens).
- **Edge cases:** a close with an empty stack (`(]` → false at first close), leftover opens at the end (`(((` → stack non-empty → false).

## Review checklist

- Why is a slice a natural stack but a *leaky* queue, and how would you fix the queue?
- Why must the stack be empty at the end of valid-parentheses, not just balanced in count?
- What's wrong with returning `-1` from `Pop`, and what would you return instead?
- Which classic algorithms are "just a queue" (BFS) vs "just a stack" (DFS / expression eval)?
