package main

import (
	"fmt"
)

func main() {

	var graph [][]int

	g_size := 6

	fmt.Println("test", graph)

	for i := 0 ; i < g_size ; i++ {

		graph = append(graph, make([]int, g_size))

	}

	print_graph(graph)

	connected_points := [][]int{
		{ 1, 2 },
		{ 0, 1 },
		{ 1, 3 },
		{ 1, 4 },
		{ 3, 4 },
		{ 4, 5 } }

	for _, points := range connected_points {

		fmt.Println(points)

		connect_points(graph, points[0], points[1])

	}

	print_graph(graph)

}

func connect_points(graph [][]int, x, y int) [][]int {

	graph[x][y] += 1
	graph[y][x] += 1

	return graph

}

func print_graph(graph [][]int) {

	for i := 0 ; i < len(graph) ; i++ {

		fmt.Println(graph[i])

	}

	fmt.Println()

}
