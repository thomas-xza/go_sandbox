package main

import (
	"fmt"
	"errors"
)

type Linked_list_elem struct {

	v int
	linked bool
	next *Linked_list_elem
	
}

func main() {

	linked_list_1 := make([]Linked_list_elem, 6)

	// initialise_mem(linked_list_1, 6 - 2)

	fmt.Println(linked_list_1)

	empty_pt := find_unused_element(linked_list_1)

	element_pt = find_nth_element(&linked_list_1[0], 3)
	
	insert_element(element_pt, 20)

}

func insert_element(element_source *Linked_list_elem, n int) {

	//  This function is for inserting anywhere after the initial element.


	// fmt.Printf("%p\n", element_source)

}


func find_unused_element(allocated_mem []Linked_list_elem) (int, error) {

	for i := range allocated_mem {

		if allocated_mem[i].linked == false {

			return &allocated_mem[i], nil

		}

	}

	return &allocated_mem[0], errors.New("No space found")

}

func find_nth_element(start_pos *Linked_list_elem, n int) *Linked_list_elem {

	//  This function follows the list, from the start position,
	//  until it reaches the nth link in the list.

	var next_element_pt *Linked_list_elem = start_pos

	fmt.Printf("%p\n", next_element_pt)

	for i := 0 ; i < n - 1; i++ {

		next_element_pt = next_element_pt.next

		fmt.Printf("%p\n", next_element_pt)

	}

	return next_element_pt

}

// func initialise_mem(linked_list []Linked_list_elem, size int) {

// 	for i := range linked_list {

// 		if i < (size - 1) {

// 			linked_list[i].next = &linked_list[i + 1]

// 			linked_list[i].linked = true

// 		}

// 		linked_list[i].v = i

// 	}

// }
