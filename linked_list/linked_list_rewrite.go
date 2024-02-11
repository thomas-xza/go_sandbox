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

	linked_list := make([]Linked_list_elem, 6)

	fmt.Println(linked_list)

	initialise_list(linked_list, 6 - 2)

	fmt.Println(linked_list)
	
	_ = find_nth_element(linked_list, 0, 3)

	// insert_element(linked_list, element_pt, 20)

}

func find_nth_element(linked_list []Linked_list_elem, start_pos *Linked_list_elem, n int) *Linked_list_elem {

	//  This function follows the list, from the start position,
	//  until it reaches the nth link in the list.

	var next_element_pt *Linked_list_elem

	for i := 0 ; i < n - 1; i++ {

		// if *next_element_pt.next != nil {

		// 	next_element_pt = *next_element_pt.next

		// }

		fmt.Println(next_element_pt)

	}

	return next_element_pt

}

func insert_element(linked_list []Linked_list_elem, element_source *Linked_list_elem, n int) {

	//  This function is for inserting anywhere after the initial element.

	// empty_pt := find_unused_element(linked_list)

	// fmt.Printf("%p\n", element_source)

}


func find_unused_element(linked_list []Linked_list_elem) int {

	for i := range linked_list {

		if linked_list[i].linked == false {

			return 0

		}

	}

	return 0

}


func initialise_list(linked_list []Linked_list_elem, size int) {

	for i := range linked_list {

		if i < (size - 1) {

			linked_list[i].next = &linked_list[i + 1]

			linked_list[i].linked = true

		}

		linked_list[i].v = i

	}

}
