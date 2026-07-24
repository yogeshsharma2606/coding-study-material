package main

import "fmt"

func dfs(graph map[int][]int, node int, visited map[int]bool) {
	visited[node] = true
	fmt.Print(node, " ")

	for _, nei := range graph[node] {
		if !visited[nei] {
			dfs(graph, nei, visited)
		}
	}
}

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {2},
		2: {0, 3},
		3: {3},
	}
	visited := make(map[int]bool)
	dfs(graph, 0, visited)
}
