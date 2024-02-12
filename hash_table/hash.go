package main

import (
	"fmt"
	"errors"
)

//  Could (or should) be replaced with generics

type key_value struct {

	key int
	data int
	save bool

}

func main() {

	k_v_store := make([]key_value, 100)

	for i := 0 ; i < 105 ; i++ {
		
		_ = data_to_key_value_store(k_v_store, i, i * 10)

	}

	data, _ := retrieve_from_key_value_store(k_v_store, 60)

	fmt.Println(data)
	
	_ = delete_from_key_value_store(k_v_store, 60)

	_, err_2 := retrieve_from_key_value_store(k_v_store, 60)

	fmt.Println(err_2)
	
}

func delete_from_key_value_store(k_v_store []key_value, key int) (error) {

	key_pos, err := seek_to_key(k_v_store, key)

	if err == nil {

		k_v_store[key_pos].key = 0
		k_v_store[key_pos].data = 0
		k_v_store[key_pos].save = false

		return nil

	} else {

		return errors.New("Could not delete from key_value store, which could be for various reasons.")

	}

}

func retrieve_from_key_value_store(k_v_store []key_value, key int) (int, error) {

	key_pos, err := seek_to_key(k_v_store, key)

	if err == nil {

		return k_v_store[key_pos].data, nil

	} else {

		return 0, errors.New("Key not found")

	}

}


func seek_to_key(k_v_store []key_value, key int) (int, error) {

	hashed_key := hash_key(len(k_v_store), key)

	target := 0

	for i := 0 ; i < len(k_v_store) ; i++ {

		target = (hashed_key + i) % len(k_v_store)

		if k_v_store[target].key == key {

			return target, nil

		}

	}

	return 0, errors.New("Not found")
	

}


func data_to_key_value_store(k_v_store []key_value, key int, data int) error {

	hashed_key := hash_key(len(k_v_store), key)

	target := 0

	for i := 0 ; i < len(k_v_store) ; i++ {

		target = (hashed_key + i) % len(k_v_store)

		if k_v_store[target].save == false {

			k_v_store[hashed_key].key = key
			k_v_store[hashed_key].data = data
			k_v_store[hashed_key].save = true

			return nil

		}

	}

	return errors.New("No available space in k_v_store")

}

func hash_key(limit int, key int) int {

	//  Todo: replace with true hashing function.

	return key * 9339487593 % limit

}
