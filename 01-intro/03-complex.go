// complex types
package main

import "fmt"

func main() {
	// var c1 complex64 = 4 + 5i
	// var c1 = 4 + 5i
	c1 := 4 + 5i
	// fmt.Println(c1)

	c2 := 10 + 11i
	result := c1 + c2

	fmt.Println(result)
	fmt.Printf("Real : %f\n", real(result))
	fmt.Printf("Imaginary : %f\n", imag(result))
}
