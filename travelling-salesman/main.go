package main

import (
	"fmt"
	"math"
)

// calculates the total distance of a tour
func tourCost(tour []int, dist [][]int) int {
	cost := 0
	n := len(tour)
	for i := 0; i < n-1; i++ {
		cost += dist[tour[i]][tour[i+1]]
	}
	cost += dist[tour[n-1]][tour[0]] // add the cost of returning to the starting city
	return cost
}

// swaps two elements in a slice
func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}

// A recursive function to generate all permutations of a slice
func permute(s []int, l int, r int, minTour *[]int, minCost *int, dist [][]int) {
	if l == r { // base case: a permutation is generated
		cost := tourCost(s, dist) // calculate the cost of the tour
		if cost < *minCost {      // update the minimum cost and tour if needed
			*minCost = cost
			copy(*minTour, s)
		}
	} else {
		for i := l; i <= r; i++ {
			swap(s, l, i)                              // swap the current element with the leftmost element
			permute(s, l+1, r, minTour, minCost, dist) // recur for the remaining slice
			swap(s, l, i)                              // backtrack and restore the original slice
		}
	}
}

func main() {
	// A sample distance matrix
	dist := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}

	n := len(dist) // number of cities

	// A slice to store the current tour
	tour := make([]int, n)
	for i := 0; i < n; i++ {
		tour[i] = i
	}

	// A slice to store the minimum tour
	minTour := make([]int, n)

	// A variable to store the minimum cost
	minCost := math.MaxInt32

	// Call the permute function to find the optimal tour
	permute(tour, 0, n-1, &minTour, &minCost, dist)

	// Print the optimal tour and cost
	fmt.Println("Optimal tour:", minTour)
	fmt.Println("Optimal cost:", minCost)
}
