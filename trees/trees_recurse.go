package main

import (
	"fmt"
)

type Tree struct {

	v int
	l *Tree
	r *Tree

}

func main() {

	// t_layers := 3 - 1

	// t_elems := two_pow_summation(t_elements)

	// t_elems_linked := two_pow_summation(t_elements - 1)

	var t_set []Tree

	fmt.Println(t_set)
	
}

func tree_gen(t_root Tree, t Tree, t_elems int, t_elems_linked int) Tree {

	

}

func two_pow_summation(n int) int {

	//  Generate a summation of powers of 2 up to n layers of tree.

	if n == 0 {

		return 1

	} else {

		return two_pow_summation(n - 1) + int_pow(2, n)

	}

}

func int_pow(n, m int) int {
	
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	
	for i := 2; i <= m; i++ {
		
		result *= n
		
	}
	
	return result
	
}
