// Package matrix contains 2D matrix / grid interview problems.
//
// A grid is a 2D array indexed [row][col]. Common moves: TRANSPOSE + reverse for
// rotation, LAYER-BY-LAYER traversal for spirals, using the matrix itself as
// O(1) storage for flags, and treating the grid as an implicit graph for
// flood/BFS. When rows and columns are sorted, binary-search-like staircase
// walks give sublinear search.
package matrix

// RotateImage rotates an n x n matrix 90 degrees clockwise in place.
// Trick: transpose (swap across the main diagonal), then reverse each row.
func RotateImage(m [][]int) {
	n := len(m)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j] // transpose
		}
	}
	for i := 0; i < n; i++ {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			m[i][l], m[i][r] = m[i][r], m[i][l] // reverse row
		}
	}
}

// SpiralOrder returns matrix elements in clockwise spiral order by shrinking the
// four boundaries after traversing each edge.
func SpiralOrder(m [][]int) []int {
	if len(m) == 0 {
		return nil
	}
	var res []int
	top, bottom, left, right := 0, len(m)-1, 0, len(m[0])-1
	for top <= bottom && left <= right {
		for c := left; c <= right; c++ {
			res = append(res, m[top][c])
		}
		top++
		for r := top; r <= bottom; r++ {
			res = append(res, m[r][right])
		}
		right--
		if top <= bottom {
			for c := right; c >= left; c-- {
				res = append(res, m[bottom][c])
			}
			bottom--
		}
		if left <= right {
			for r := bottom; r >= top; r-- {
				res = append(res, m[r][left])
			}
			left++
		}
	}
	return res
}

// SetZeroes: if a cell is 0, zero its entire row and column, in place with O(1)
// extra space. Use the first row/col as flags; handle their own state separately.
func SetZeroes(m [][]int) {
	rows, cols := len(m), len(m[0])
	firstRowZero, firstColZero := false, false
	for c := 0; c < cols; c++ {
		if m[0][c] == 0 {
			firstRowZero = true
		}
	}
	for r := 0; r < rows; r++ {
		if m[r][0] == 0 {
			firstColZero = true
		}
	}
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if m[r][c] == 0 {
				m[r][0] = 0
				m[0][c] = 0
			}
		}
	}
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if m[r][0] == 0 || m[0][c] == 0 {
				m[r][c] = 0
			}
		}
	}
	if firstRowZero {
		for c := 0; c < cols; c++ {
			m[0][c] = 0
		}
	}
	if firstColZero {
		for r := 0; r < rows; r++ {
			m[r][0] = 0
		}
	}
}

// Transpose returns the transpose of an m x n matrix (rows become columns).
func Transpose(m [][]int) [][]int {
	rows, cols := len(m), len(m[0])
	out := make([][]int, cols)
	for i := range out {
		out[i] = make([]int, rows)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			out[c][r] = m[r][c]
		}
	}
	return out
}

// SearchMatrix: each row sorted, first of each row > last of previous. Treat the
// grid as one sorted array of size m*n and binary search with index mapping.
func SearchMatrix(m [][]int, target int) bool {
	if len(m) == 0 || len(m[0]) == 0 {
		return false
	}
	rows, cols := len(m), len(m[0])
	lo, hi := 0, rows*cols-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		v := m[mid/cols][mid%cols]
		if v == target {
			return true
		} else if v < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

// SearchMatrixII: rows and columns each sorted ascending. Start at the top-right
// corner: too big -> move left; too small -> move down. Each step eliminates a
// row or column. O(m+n).
func SearchMatrixII(m [][]int, target int) bool {
	if len(m) == 0 || len(m[0]) == 0 {
		return false
	}
	r, c := 0, len(m[0])-1
	for r < len(m) && c >= 0 {
		switch {
		case m[r][c] == target:
			return true
		case m[r][c] > target:
			c--
		default:
			r++
		}
	}
	return false
}

// FloodFill changes the color of the connected region containing (sr, sc).
func FloodFill(image [][]int, sr, sc, newColor int) [][]int {
	old := image[sr][sc]
	if old == newColor {
		return image
	}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != old {
			return
		}
		image[r][c] = newColor
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	dfs(sr, sc)
	return image
}

// GameOfLife updates the board in place. Encode transitions in the 2nd bit so
// neighbor counts still read the old state (bit 0): 01->11 means "was live, now
// live"; 00->10 means "was dead, now live". Finally shift each cell right by 1.
func GameOfLife(board [][]int) {
	rows, cols := len(board), len(board[0])
	countLive := func(r, c int) int {
		cnt := 0
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr == 0 && dc == 0 {
					continue
				}
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && board[nr][nc]&1 == 1 {
					cnt++
				}
			}
		}
		return cnt
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			live := countLive(r, c)
			if board[r][c]&1 == 1 {
				if live == 2 || live == 3 {
					board[r][c] |= 2 // stays live
				}
			} else if live == 3 {
				board[r][c] |= 2 // becomes live
			}
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			board[r][c] >>= 1
		}
	}
}

// DiagonalTraverse returns matrix elements in a zig-zag diagonal order.
func DiagonalTraverse(m [][]int) []int {
	if len(m) == 0 {
		return nil
	}
	rows, cols := len(m), len(m[0])
	res := make([]int, 0, rows*cols)
	for d := 0; d < rows+cols-1; d++ {
		if d%2 == 0 { // going up-right
			r := min(d, rows-1)
			c := d - r
			for r >= 0 && c < cols {
				res = append(res, m[r][c])
				r--
				c++
			}
		} else { // going down-left
			c := min(d, cols-1)
			r := d - c
			for c >= 0 && r < rows {
				res = append(res, m[r][c])
				r++
				c--
			}
		}
	}
	return res
}
