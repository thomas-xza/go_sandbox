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

	midpoint := len(set) / 2

	fmt.Println(midpoint)

	if len(set) <= 1 {

		return set
		
	} else {

		set_a := set[:midpoint]

		set_b := set[midpoint:]

		return merge(divide(set_a), divide(set_b))

	}

}

func merge(set_a []int, set_b []int) ([]int) {

	fmt.Println(set_a, set_b)

	return append(set_a, set_b...)

	// } else {

	// 	return append(set_b, set_a...)

	// }

}
