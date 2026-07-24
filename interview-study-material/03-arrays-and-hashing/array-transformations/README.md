# Array Transformations

In-place rearrangement problems test whether you can manipulate indices carefully and reason about **auxiliary space**. Rotation is the canonical example, and it hides a beautiful O(1)-space trick.

## Programs

### [Rotate array with extra space](rotate-array-extra-space/rotate_array_extra_space.go)

Rotate right by `k` by slicing: take the last `k` elements and put them in front.

**How to think about it:** rotating right by `k` means the last `k` elements wrap to the front. With extra space that's a one-liner: `append(arr[n-k:], arr[:n-k]...)`. First reduce `k = k % n` so `k > n` doesn't break the slice bounds (rotating by `n` is a no-op).

**Dry run** — `[1..7], k=3`: last 3 = `[5,6,7]`, first 4 = `[1,2,3,4]` → `[5,6,7,1,2,3,4]`.

- **O(n) time, O(n) space** (the new slice).
- This is the "get it working" version; the interviewer's follow-up is almost always *"now do it with O(1) space."*

### [Rotate array in place](rotate-array-in-place/rotate_array_in_place.go)

Rotate right by `k` using the **three-reversal trick** — O(1) extra space.

**The insight (this is the memorable part):** rotating is just moving two blocks past each other. If you (1) reverse the **whole** array, the two blocks are now in the right *order* but each is internally backward; then (2) reverse the **first k** and (3) reverse the **rest**, each block is fixed internally. Net effect: a perfect rotation with no extra array.

**Dry run** — `[1,2,3,4,5,6,7], k=3`:

```
reverse all      -> [7,6,5,4,3,2,1]
reverse [0..k-1] -> [5,6,7,4,3,2,1]
reverse [k..n-1] -> [5,6,7,1,2,3,4]   ✓
```

- **O(n) time, O(1) space.** Remember `k %= n` and guard the empty array.
- **Why it's worth memorizing:** the reverse-the-parts trick reappears in string rotation, "rotate words in a sentence", etc.

## Review checklist

- Why must you compute `k %= n` before rotating?
- Explain the three-reversal trick and why the final array is correctly rotated.
- Extra-space vs in-place: state the space trade-off out loud.
- Right rotation by `k` equals left rotation by what?
