package main

import (
	"fmt"
)

func main() {

	g_size := 6

	graph := initialise_graph(g_size)

	// bfs_traversal(graph, g_size)

	dfs_traversal(graph, g_size)

}

func dfs_traversal(graph [][]int, g_size int) {

	var stack []int

	visited := make([]int, g_size)

	visited[0] = 1

	stack = append(stack, 0)

	dfs_recurse(graph, visited, 0, stack)
	
}

func dfs_recurse(graph [][]int, visited []int, vx int, stack []int) {

	fmt.Println(stack, visited)

	_, stack = stack[len(stack)-1], stack[:len(stack)-1]

	// fmt.Println("error?")

	// vx, stack := stack[len(stack)-1], stack[:len(stack)-1]

	for vx_dest, connection := range graph[vx] {

		if connection > 0 && visited[vx_dest] == 0 {

			visited[vx_dest] = 1

			stack = append(stack, vx_dest)

			dfs_recurse(graph, visited, vx_dest, stack)

		}
			
	}

}

func bfs_traversal(graph [][]int, g_size int) {

	var queue, distances []int

	for i := 0 ; i < g_size ; i++ {

		distances = append(distances, -1)

	}

	distances[0] = 0

	queue = append(queue, 0)

	var vx int
	
	for len(queue) != 0 {

		vx, queue = queue[0], queue[1:]

		for vx_dest, connection := range graph[vx] {

			if connection > 0 && distances[vx_dest] == -1 {

				queue = append(queue, vx_dest)

				distances[vx_dest] = distances[vx] + 1

			}

		}

	}

	print_graph(graph)

}

func initialise_graph(g_size int) [][]int {

	var graph [][]int

	for i := 0 ; i < g_size ; i++ {

		graph = append(graph, make([]int, g_size))

	}

	connected_points := [][]int{
		{ 1, 2 },
		{ 0, 1 },
		{ 1, 3 },
		{ 1, 4 },
		{ 3, 4 },
		{ 4, 5 },
	}

	for _, points := range connected_points {

		connect_points(graph, points[0], points[1])

	}

	return graph
	
}

func connect_points(graph [][]int, x, y int) [][]int {

	graph[x][y] += 1
	graph[y][x] += 1

	return graph

}

func print_graph(graph [][]int) {

	fmt.Printf("\n   ")

	for i := 0 ; i < len(graph) ; i++ {

		fmt.Printf("%d ", i)

	}

	fmt.Println()

	for i := 0 ; i < len(graph) ; i++ {

		fmt.Println(i, graph[i])

	}

	fmt.Println()

}
