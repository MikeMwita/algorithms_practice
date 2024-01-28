package main

import "math"

func main() {

}

func dijkstra(graph [][]int, start int, end int) {
	n := len(graph)
	dist := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		visited[i] = false

	}
	dist[start] = 0

	for count := 0; count < n-1; count++ {
		u := -1

		for i := 0; i < n; i++ {
			if !visited[i] && (u == -1) || dist[i] < dist[u] {
				u = i
			}
		}
	}
}
