package main

import (
	"fmt"
	"errors"
)

type Linked_list_elem struct {

	v int

	linked bool

	ptr *Linked_list_elem
}

func main() {

	list_capacity := 6

	list_size := list_capacity - 2

	linked_list := make([]Linked_list_elem, list_capacity)

	for i := 0 ; i < list_size ; i++ {

		if i < (list_size - 1) {

			linked_list[i].ptr = &linked_list[i+1]
			
		}

		linked_list[i].v = i

		linked_list[i].linked = true

	}
		
	add_element(linked_list, 1, 20)
	
}

func add_element(linked_list []Linked_list_elem, pos_to_insert int, value int) {

	new_pos, err_1 := find_unused_element(linked_list)

	//  find_size_of_linked_list()

	elem_target, err_2 := find_nth_element(linked_list, pos_to_insert)

	fmt.Println(linked_list)

	if err_1 == nil && err_2 == nil {

		//  Set value within unused element.

		linked_list[new_pos].v = value

		//  Set new element to point to next in sequence.

		linked_list[new_pos].ptr = elem_target.ptr

		elem_target.ptr = &linked_list[new_pos]

		fmt.Println(elem_target, elem_target.ptr)

		fmt.Println(linked_list)

	}
	
}

func find_nth_element(linked_list []Linked_list_elem, n int) (*Linked_list_elem, error) {

	//  Warning: this assumes position 0 of the linked_list array
	//  is always the start of the linked list.

	//  If n is greater than size of all linked elements in
	//  linked_list array, this will return the last element.

	next_element := linked_list[0]

	prev_element := next_element
	
	for i := 0 ; i <= n ; i++ {

		if next_element.ptr != nil {

			prev_element = next_element

			next_element = *next_element.ptr

		} else {

			return &prev_element, errors.New("Pos argument greater than total quantity of elements within linked list, outputting last element of linked list instead.")

		}

	}

	return &prev_element, nil

}

func find_unused_element(linked_list []Linked_list_elem) (int, error) {

	for i := range linked_list {

		if linked_list[i].linked == false {

			return i, nil

		}

	}

	return 0, errors.New("not found")

}
