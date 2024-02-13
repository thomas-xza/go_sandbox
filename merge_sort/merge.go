package main

import (
	"fmt"
)

func main() {

	fmt.Println("test")

	input := []int{ 98, 13, 274, 7, 12, 981, 5 }

	fmt.Println(input)

	merge_sort(input)

}

func merge_sort(set []int) {

	fmt.Println(divide(set))

}


func divide(set []int) ([]int) {

	mid := len(set) / 2

	if len(set) <= 1 {

		return set
		
	} else {

		return merge(divide(set[:mid]), divide(set[mid:]))

	}

}

func merge(set_a []int, set_b []int) ([]int) {

	fmt.Println("merge", set_a, set_b)
	
	if (len(set_a) == 0) { return set_b }

	if (len(set_b) == 0) { return set_a }

	set_sorted := make([]int, 0)

	a_pt, b_pt := 0, 0

	for i := 0 ; i < len(set_a) + len(set_b) ; i++ {

		if (set_a[a_pt] > set_b[b_pt]) {

			set_sorted = append(set_sorted, set_a[a_pt])

			a_pt += 1

		} else {

			set_sorted = append(set_sorted, set_b[b_pt])

			b_pt += 1

		}			

		i += 1

	}

	fmt.Println(set_sorted)

	return append(set_sorted)

}
