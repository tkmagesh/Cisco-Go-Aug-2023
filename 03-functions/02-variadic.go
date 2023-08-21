package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum(10, 20, 30, 40, 50))
}

/* NOTE: the variadic parameter should be LAST one in the parameter list */
func sum(nos ...int) int {
	result := 0
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}
