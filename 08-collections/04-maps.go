package main

import "fmt"

func main() {
	// var productRanks map[string]int // memory not allocated
	// var productRanks map[string]int = make(map[string]int)
	// var productRanks map[string]int = map[string]int{}

	// var productRanks map[string]int = map[string]int{"pen": 2, "pencil": 1, "marker": 5}
	/*
		var productRanks map[string]int = map[string]int{
			"pen":    2,
			"pencil": 1,
			"marker": 5,
		} */

	productRanks := map[string]int{
		"pen":    2,
		"pencil": 1,
		"marker": 5,
	}

	//adding a new item
	productRanks["notepad"] = 3
	fmt.Println(productRanks)

	fmt.Println("len(productRanks) = ", len(productRanks))

	// check for the existance of a key
	key := "mouse"
	if val, exists := productRanks[key]; exists {
		fmt.Printf("Rank of %q is %d\n", key, val)
	} else {
		fmt.Printf("key %q doesnt exist\n", key)
	}

	// removing an item
	delete(productRanks, "notepad") // delete is a noop if the key doesnt exist
	fmt.Println(productRanks)

	//iterating a map
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

}
