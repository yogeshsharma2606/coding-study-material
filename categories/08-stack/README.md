# 08 - Stack

> Code: [`stack.go`](stack.go) - Tests: [`stack_test.go`](stack_test.go) - run `go test ./categories/08-stack/`

## Overview & mental model

A stack is **LIFO**: the last thing pushed is the first popped. It answers "what is the most recent unresolved item?" You **push to defer** work and **pop to resolve** it against the current element. In Go, a slice is a natural stack: `append` to push, `s = s[:len(s)-1]` to pop, `s[len(s)-1]` to peek.

Stacks are the right tool for: **matching** (brackets, tags), **evaluating expressions** (postfix directly, infix with two stacks or sign-tracking), **nested structure** (decode/expand), and **collision/undo** semantics.

## How to recognize it

- Nested or paired structure: parentheses, tags, directories, `k[...]` encodings.
- "Evaluate an expression", "most recent matching X".
- Something that "cancels" a previous item (backspace `#`, asteroid collisions).
- (For "next greater/smaller", see [09 - Monotonic Stack](../09-monotonic-stack/).)

## How to think / attack plan

1. What is the "unresolved" thing I need to remember until a later element resolves it? Push that.
2. On each element, decide: does it **open** (push) or **close/resolve** (pop and act)?
3. For expressions, ask if postfix (just a value stack) or infix (track sign/precedence, save state on `(`).
4. At the end, an empty stack usually means "balanced"; leftovers mean unmatched.

---

## Problems (easy -> hard)

### 1. Valid Parentheses (easy)
**Thinking.** Push openers. On a closer, the top must be its matching opener - else invalid. Empty stack at the end = balanced.
**Dry run.** `([)]`: push (, [, then ) needs [ on top but sees ( wait - top is [ so ) mismatches -> false.
**Complexity.** Time O(n), space O(n).

### 2. Min Stack (medium, design)
**Direction of thinking.** `Min()` in O(1) means you can't scan. Keep a **parallel stack** where each entry is the min of everything at or below it. Push mirrors the min; pop stays in sync.
**Complexity.** All ops O(1), space O(n).

### 3. Evaluate Reverse Polish Notation (medium)
**Direction of thinking.** Postfix removes precedence/parentheses entirely: operands push; an operator pops its two operands and pushes the result. Mind operand order for `-` and `/` (second popped is the left operand).
**Dry run.** `2 1 + 3 *`: push 2,1 -> +=3 -> push 3 -> *=9.
**Complexity.** Time O(n), space O(n).

### 4. Simplify Path (medium)
**Direction of thinking.** Split on `/`. `..` pops the last directory, `.`/empty are no-ops, names push. Join the stack. The stack models "go up one level".
**Complexity.** Time O(n), space O(n).

### 5. Decode String (medium)
**Direction of thinking.** Nesting like `3[a2[c]]` calls for two stacks: the **repeat counts** and the **string built so far before each `[`**. On `]`, pop the count and prefix, and set `cur = prefix + repeat(cur, count)`.
**Dry run.** `3[a2[c]]`: at inner `]` cur="cc"; at outer `]` cur = "a"+... = "acc" repeated 3 -> "accaccacc".
**Complexity.** Time O(output length), space O(depth).

### 6. Backspace String Compare (easy)
**Thinking.** Build each final string on a stack (`#` pops), then compare. (An O(1)-space variant scans from the right.)
**Complexity.** Time O(n+m), space O(n+m).

### 7. Asteroid Collision (medium)
**Direction of thinking.** Only a right-mover already on the stack can collide with an incoming left-mover. Loop: while the top is positive and the incoming is negative, compare sizes - the smaller explodes; equal, both go. The stack holds survivors.
**Dry run.** `[10,2,-5]`: push 10,2; -5 vs 2 -> 2 dies; -5 vs 10 -> -5 dies -> `[10]`.
**Complexity.** Time O(n), space O(n).

### 8. Basic Calculator (hard)
**Direction of thinking.** For `+ - ( )`, track a running `result`, current `sign`, and `num`. On `(`, push `result` and `sign` and reset; on `)`, fold the inner result back with the saved sign. This avoids a full parser.
**Dry run.** `(1+(4+5+2)-3)+(6+8)` = 23.
**Complexity.** Time O(n), space O(n).

---

## Common pitfalls & edge cases

- Popping an empty stack (check `len > 0` before peek/pop).
- Operand order for non-commutative operators in RPN.
- Forgetting to flush the final `num`/`result` after the loop in expression parsing.
- Building strings with `+=` inside deep nesting - acceptable here but note `strings.Builder` for hot paths.

## Interview Q&A

- **Why a stack for bracket matching?** Nesting is inherently LIFO - the most recently opened must close first.
- **Postfix vs infix evaluation?** Postfix needs only a value stack (no precedence); infix needs precedence handling or sign/state stacks.
- **How does Min Stack keep O(1) min?** A parallel stack storing the running minimum per level.
- **Slice as stack - any gotcha?** Reslicing keeps the backing array; for long-lived stacks with huge churn, occasionally reallocate to release memory.
