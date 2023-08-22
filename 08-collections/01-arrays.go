package main

import "fmt"

func main() {
	// var nos [5]int // memory allocated, initialized with the default value of 'int'
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}

	//type inference
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating an array using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating an array using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	/*
		p1 := [3]string{"pen", "pencil", "marker"}
		p2 := [3]string{"pen", "pencil", "marker"}
		fmt.Println(p1 == p2)
	*/

	fmt.Println("Before sorting, nos = ", nos)
	sortInts(&nos)
	fmt.Println("After sorting, nos = ", nos)
}

func sortInts(list *[5]int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i] // no need to derefernce
			}
		}
	}
}
