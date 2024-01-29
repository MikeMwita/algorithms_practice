package main

import (
	"container/heap"
	"fmt"
)

// Node represents a node in the graph, with a name and a map of neighbors and their weights.
type Node struct {
	name      string
	neighbors map[*Node]int
}

// Graph represents a graph, with a map of nodes and their names.
type Graph struct {
	nodes map[string]*Node
}

// Heap represents a heap that implements the container/heap interface, with a slice of nodes and their distances.
type Heap struct {
	nodes []*Node       // slice of nodes
	dist  map[*Node]int // map of node distances
	index map[*Node]int // map of node indices in the slice
}

// Len returns the length of the heap.
func (h *Heap) Len() int {
	return len(h.nodes)
}

// Less returns true if the node at index i has a lower distance than the node at index j.
func (h *Heap) Less(i, j int) bool {
	return h.dist[h.nodes[i]] < h.dist[h.nodes[j]]
}

// Swap swaps the nodes at indices i and j.
func (h *Heap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
	h.index[h.nodes[i]] = i
	h.index[h.nodes[j]] = j
}

// Push adds a node to the heap.
func (h *Heap) Push(x interface{}) {
	n := x.(*Node)
	h.index[n] = len(h.nodes)
	h.nodes = append(h.nodes, n)
}

// Pop removes and returns the node with the lowest distance from the heap.
func (h *Heap) Pop() interface{} {
	n := h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	delete(h.index, n)
	return n
}

// Update updates the distance and position of a node in the heap.
func (h *Heap) Update(n *Node, dist int) {
	h.dist[n] = dist
	heap.Fix(h, h.index[n])
}

// Dijkstra returns the shortest path and its distance from the source node to the destination node in the graph, or an error if there is no path.
func Dijkstra(g *Graph, source, destination string) ([]string, int, error) {
	// Initialize the heap with the source node and its distance (zero).
	h := &Heap{
		dist:  make(map[*Node]int),
		index: make(map[*Node]int),
	}
	h.dist[g.nodes[source]] = 0
	heap.Push(h, g.nodes[source])

	// Initialize a map to store the previous node in the path for each node.
	prev := make(map[*Node]*Node)

	// Loop until the heap is empty or the destination node is visited.
	for h.Len() > 0 {
		// Pop the node with the lowest distance from the heap.
		u := heap.Pop(h).(*Node)

		// If the node is the destination, we are done.
		if u.name == destination {
			break
		}

		// Loop over the neighbors of the node.
		for v, w := range u.neighbors {
			// Calculate the distance to the neighbor through the node.
			alt := h.dist[u] + w

			// If the distance is smaller than the current distance, update the distance and the previous node, and push the neighbor to the heap.
			if alt < h.dist[v] {
				h.Update(v, alt)
				prev[v] = u
			}
		}
	}

	// If the destination node is not in the previous map, there is no path.
	if _, ok := prev[g.nodes[destination]]; !ok {
		return nil, 0, fmt.Errorf("no path from %s to %s", source, destination)
	}

	// Construct the path by following the previous nodes from the destination to the source.
	path := []string{destination}
	for n := prev[g.nodes[destination]]; n != nil; n = prev[n] {
		path = append(path, n.name)
	}

	// Reverse the path to get the correct order.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	// Return the path and its distance.
	return path, h.dist[g.nodes[destination]], nil
}

// addEdge adds an edge with a weight between two nodes in the graph, creating the nodes if they don't exist.
func (g *Graph) addEdge(u, v string, w int) {
	// If the graph is nil, initialize it.
	if g.nodes == nil {
		g.nodes = make(map[string]*Node)
	}

	// If the node u is not in the graph, create it.
	if _, ok := g.nodes[u]; !ok {
		g.nodes[u] = &Node{name: u, neighbors: make(map[*Node]int)}
	}

	// If the node v is not in the graph, create it.
	if _, ok := g.nodes[v]; !ok {
		g.nodes[v] = &Node{name: v, neighbors: make(map[*Node]int)}
	}

	// Add the edge with the weight between the nodes.
	g.nodes[u].neighbors[g.nodes[v]] = w
	g.nodes[v].neighbors[g.nodes[u]] = w
}

// newGraph returns a new graph with some example edges.
func newGraph() *Graph {
	g := &Graph{}
	g.addEdge("A", "B", 7)
	g.addEdge("A", "C", 9)
	g.addEdge("A", "F", 14)
	g.addEdge("B", "C", 10)
	g.addEdge("B", "D", 15)
	g.addEdge("C", "D", 11)
	g.addEdge("C", "F", 2)
	g.addEdge("D", "E", 6)
	g.addEdge("E", "F", 9)
	return g
}

func main() {
	// Create a new graph with some example edges.
	g := newGraph()

	// Find the shortest path from A to E.
	path, dist, err := Dijkstra(g, "A", "E")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Path:", path, "Distance:", dist)
	}
}
