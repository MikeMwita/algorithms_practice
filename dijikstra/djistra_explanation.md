# Dijkstra's Algorithm

This code implements Dijkstra's algorithm for finding the shortest path between nodes in a graph.

## Types

The code  defines the following types:

- `Node`: A node in the graph, with a name and a map of neighbors and their weights.
- `Graph`: A graph, with a map of nodes and their names.
- `Heap`: A heap that implements the `container/heap` interface, with a slice of nodes and their distances, and some methods to push, pop, and update the heap elements.

## Functions

The code defines the following functions:
- `Dijkstra`: A function that takes a graph, a source node, and a destination node, and returns the shortest path and its distance, or an error if there is no path. It uses a heap to store the unvisited nodes and their distances, and a map to store the previous nodes in the path.

