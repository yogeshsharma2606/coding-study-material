# Go Idioms Cheatsheet (for coding interviews)

Fast, correct Go for the whiteboard and the editor. Standard library only.

## Slices

```go
s := []int{}                 // empty
s := make([]int, n)          // n zeros
s := make([]int, 0, n)       // len 0, cap n (avoids reallocs)
s = append(s, x)             // push back (amortized O(1))
s = s[:len(s)-1]             // pop back
top := s[len(s)-1]           // peek back
s = s[1:]                    // pop front (does NOT reclaim memory immediately)

cp := append([]int(nil), s...)  // copy a slice (crucial in backtracking!)
copy(dst, src)                  // element copy

// 2D slice
grid := make([][]int, rows)
for i := range grid { grid[i] = make([]int, cols) }
```

Gotchas:
- `append` may reallocate; a subslice shares the backing array until it does. Copy when you need independence.
- Appending to a slice passed by value won't be seen by the caller if it reallocates - return the slice.
- Range makes a **copy** of each element: `for _, v := range s { v = ... }` doesn't mutate; use `s[i]`.

## Maps and sets

```go
m := map[string]int{}
m[k]++                        // counter (missing key defaults to 0)
v, ok := m[k]                 // presence check
delete(m, k)

set := map[int]struct{}{}     // set (struct{} is zero-width)
set[x] = struct{}{}
_, in := set[x]
```

Gotchas:
- **Map iteration order is randomized** - never rely on it; sort output for determinism.
- Reading a missing key returns the zero value (no panic); writing to a nil map panics.
- For small known alphabets (a-z), prefer `[26]int` over a map (faster, O(1) space, comparable with `==`).

## min / max / abs

```go
// Go 1.21+ builtins:
min(a, b); max(a, b, c)      // variadic, work on ordered types

func abs(x int) int { if x < 0 { return -x }; return x }
```

## Sorting

```go
sort.Ints(a); sort.Strings(a)
sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })   // not stable
sort.SliceStable(a, less)                                    // stable
sort.Sort(sort.Reverse(sort.IntSlice(a)))                    // descending
i := sort.Search(n, func(i int) bool { return a[i] >= target }) // lower bound
```

## Heap (`container/heap`)

`container/heap` is algorithms over a type implementing `heap.Interface`:

```go
type IntHeap []int
func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] } // min-heap
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any           { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }

h := &IntHeap{}
heap.Init(h)
heap.Push(h, 5)
top := heap.Pop(h).(int)
```

For a max-heap, flip `Less` to `>`. See [11 - Heap](../categories/11-heap-priority-queue/) for a generic `PQ[T]`.

## Strings and bytes

```go
b := []byte(s)               // mutable copy (ASCII)
r := []rune(s)               // Unicode code points
s := string(b)

var sb strings.Builder       // efficient concatenation
sb.WriteString("x"); sb.WriteByte('y'); result := sb.String()

strconv.Atoi("42"); strconv.Itoa(42)
strings.Split, strings.Join, strings.Fields, strings.Repeat
```

Gotchas:
- Strings are **immutable**; `s[i] = x` won't compile. Convert to `[]byte`/`[]rune`.
- `len(s)` is **bytes**, not characters; iterate runes for Unicode.
- Never build strings with `+=` in a loop (O(n^2)); use `strings.Builder`.

## Integers

```go
const MaxInt = math.MaxInt          // platform int max (also MaxInt32/64)
a % b                                // sign follows the dividend! normalize: ((a%b)+b)%b
// overflow: Go ints are 64-bit on most platforms; still guard 32-bit-constrained problems
```

## Generics (Go 1.18+)

```go
func Map[T, U any](s []T, f func(T) U) []U {
    r := make([]U, len(s))
    for i, v := range s { r[i] = f(v) }
    return r
}

import "cmp"
import "slices"    // slices.Sort, slices.Contains, slices.Index, slices.Reverse
import "maps"      // maps.Keys, maps.Values
```

## Concurrency quick reference

```go
var wg sync.WaitGroup
wg.Add(1); go func(){ defer wg.Done(); /* ... */ }(); wg.Wait()

var mu sync.Mutex; mu.Lock(); /* ... */ mu.Unlock()
var once sync.Once; once.Do(func(){ /* init */ })

ch := make(chan int)        // unbuffered (synchronous)
ch := make(chan int, n)     // buffered (semaphore of size n)
for v := range ch { }       // reads until closed
close(ch)                   // producer closes; sending after close panics

select {
case v := <-ch:    use(v)
case <-ctx.Done(): return ctx.Err()
default:            /* non-blocking */
}
```

Run concurrency tests with `go test -race` (needs CGO enabled).

## Testing (table-driven)

```go
func TestFoo(t *testing.T) {
    tests := []struct{ in, want int }{{1, 2}, {3, 4}}
    for _, tt := range tests {
        if got := Foo(tt.in); got != tt.want {
            t.Errorf("Foo(%d)=%d want %d", tt.in, got, tt.want)
        }
    }
}
```

- `reflect.DeepEqual(a, b)` compares slices/maps.
- `go test ./...` runs everything; `go vet ./...` catches common bugs; `gofmt`/`go fmt` formats.
