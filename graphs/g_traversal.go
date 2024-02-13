package main

import (
	"fmt"
)

func main() {

	g_size := 6

	graph := initialise_graph(g_size)

	var queue, distances []int

	for i := 0 ; i < g_size ; i++ {

		distances = append(distances, -1)

	}

	distances[0] = 0

	queue = append(queue, 0)

	var vx int
	
	for len(queue) != 0 {

		// fmt.Println(queue)
		vx, queue = queue[0], queue[1:]

		fmt.Println(distances)

		for vx_dest, connection := range graph[vx] {

			if connection > 0 && distances[vx_dest] == -1 {

				queue = append(queue, vx_dest)

				distances[vx_dest] = distances[vx] + 1

			}

			// fmt.Println(v)

		}

		fmt.Println(vx, queue)

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
