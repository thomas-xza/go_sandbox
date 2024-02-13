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

	two_pow := 0

	two_pow_prev := 0

	//  Generate a summation of powers of 2 up to n layers of tree.

	for i := 0 ; i < 3 ; i++ {

		two_pow_prev = two_pow

		two_pow += int_pow(2, i)
		
	}

	t_set := make([]Tree, two_pow)

	// for i := range t_set {

	// 	fmt.Printf("%d %p\n", i, &t_set[i])

	// }

	for i := range t_set {

		t_set[i].v = i

		if i < two_pow_prev {

			t_set[i].l = &t_set[i + i + 1]

			t_set[i].r = &t_set[i + i + 2]

		}

	}

	fmt.Println(t_set)

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
