package main

import "fmt"

func RemainingLeaves(width int, height int, leaves [][]int, winds string) int {
	for _, w := range winds {
		// Create new empty grid
		newGrid := make([][]int, height)
		for i := 0; i < height; i++ {
			newGrid[i] = make([]int, width)
		}

		for r := 0; r < height; r++ {
			for c := 0; c < width; c++ {
				count := leaves[r][c]
				if count == 0 {
					continue
				}

				nr, nc := r, c
				switch w {
				case 'U':
					nr--
				case 'D':
					nr++
				case 'L':
					nc--
				case 'R':
					nc++
				}

				// If inside grid, move leaves
				if nr >= 0 && nr < height && nc >= 0 && nc < width {
					newGrid[nr][nc] += count
				}
				// else: leaves are lost
			}
		}

		leaves = newGrid
	}

	// Sum remaining leaves
	total := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			total += leaves[r][c]
		}
	}

	return total
}

func main() {
	fmt.Println(RemainingLeaves(3, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, "UDLR"))
}
