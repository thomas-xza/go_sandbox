package main

import (
	"fmt"
	// "errors"
)

type Linked_list_elem struct {

	v int
	linked bool
	next *Linked_list_elem
	
}

func main() {

	linked_list := make([]*Linked_list_elem, 6)
	
	fmt.Println(linked_list)

	initialise_list(linked_list, 6 - 2)

	fmt.Println(linked_list)
	
}

func initialise_list(linked_list_slc_pt []*Linked_list_elem, size int) {

	fmt.Println(linked_list_slc_pt)

	for i := 0 ; i < size ; i++ {

		linked_list_slc_pt[i] = &Linked_list_elem{ v: i, linked: true }

		// fmt.Println(linked_list_slc_pt[i])

		// if i < (size - 1) {

		// 	source_node_pt := linked_list_slc_pt[i]

		// 	dest_node_pt := linked_list_slc_pt[i + 1]

			

		// 	fmt.Println(*source_node_pt, *dest_node_pt)

		// 	// source_node_pt.next = dest_node_pt
			
		// }

		// *(linked_list_slc_pt + i).v = i

		// *(linked_list_slc_pt + i).linked = true

	}

}
