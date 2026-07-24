package main

import "fmt"

func hasCycle(graph map[int][]int, node, parent int, visited map[int]bool) bool {
	visited[node] = true
	for _, nei := range graph[node] {
		if !visited[nei] {
			if hasCycle(graph, nei, node, visited) {
				return true
			}
		} else if nei != parent {
			return true
		}
	}
	return false
}

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {5},
		2: {4, 3},
		3: {6},
	}
	visited := make(map[int]bool)
	fmt.Println(hasCycle(graph, 0, -1, visited))
}
