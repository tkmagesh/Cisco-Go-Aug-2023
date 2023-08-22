package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos []int = []int{3, 1, 4}
	// var nos  = []int{3, 1, 4}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating a slice using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating a slice using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// adding more data
	nos = append(nos, 10)
	nos = append(nos, 20, 30, 40)

	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	// slicing
	fmt.Println("nos[2:5] = ", nos[2:5])
	fmt.Println("nos[:5] = ", nos[:5])
	fmt.Println("nos[2:] = ", nos[2:])

	// subset := nos[2:5]
	// var subset []int = []int{0, 0, 0}
	var subset []int = make([]int, 3)
	copy(subset, nos[2:5])
	subset[0] = 1000
	fmt.Println("subset = ", subset)
	fmt.Println("nos = ", nos)

	sortInts(nos)
	fmt.Println(nos)

}

func sortInts(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i] // no need to derefernce
			}
		}
	}
}
