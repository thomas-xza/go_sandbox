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

	for j := 0 ; j < 6 ; j++ { fmt.Printf("%p \n", &linked_list_1[j]) }

	start_pt := &linked_list_1[0]

	for i := 1 ; i < 6 ; i++ {

		empty_pt, _ := find_unused_element(linked_list_1)

		// fmt.Printf("next empty: %p\n", empty_pt)

		start_pt, _ = insert_element(start_pt, start_pt, empty_pt, i * 5)

		for j := 0 ; j < 6 ; j++ { fmt.Printf("%p %d -> %p\n", &linked_list_1[j], linked_list_1[j].v, linked_list_1[j].next) }

		fmt.Println("\n")

// 		fmt.Println(linked_list_1)
		
	}

	// target_pt, _ := find_nth_element(start_pt, 0)

	// _, _ = find_nth_element(&linked_list_1[0], 3)
	
}

func insert_element(element_pt *Linked_list_elem, start_pt *Linked_list_elem, empty_pt *Linked_list_elem, n int) (*Linked_list_elem, error) {

	//  Check if the list is empty.
	
	if ( start_pt.linked == false ) {

		empty_pt.v = n
		
		empty_pt.linked = true

		return empty_pt, nil

	} else {

		//  Copy filled element to empty element
		//         []     [2] <- [1]
		//         [1] -> [2] <- [1]

		empty_pt.v = element_pt.v

		empty_pt.next = element_pt.next

		empty_pt.linked = true

		//  Insert new element and point it to copied element.
		//  [3] -> [1] -> [2]

		element_pt.v = n

		element_pt.next = empty_pt
		
		return start_pt, nil

	}

}


func find_unused_element(allocated_mem []Linked_list_elem) (*Linked_list_elem, error) {

	for i := range allocated_mem {

		if allocated_mem[i].linked == false {

			return &allocated_mem[i], nil

		}

	}

	return &allocated_mem[0], errors.New("No space found")

}

func find_nth_element(start_pos *Linked_list_elem, n int) (*Linked_list_elem, error) {

	//  This function follows the list, from the start position,
	//  until it reaches the nth link in the list.

	var next_element_pt *Linked_list_elem = start_pos

	fmt.Printf("%p %d\n", next_element_pt, next_element_pt.v)

	for i := 0 ; i < n - 1; i++ {

		next_element_pt = next_element_pt.next

		fmt.Printf("%p\n", next_element_pt)

	}

	return next_element_pt, nil

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
