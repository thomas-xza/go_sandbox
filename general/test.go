package main

import (
	"fmt"
)

func main() {

	k_v_store := make(map [int]bool)

	k_v_store[1] = true
	k_v_store[1] = false

	for k, v := range k_v_store {

		fmt.Println(k, v)

	}

}
