package main

import "fmt"

func main() {

	divisor := 0
	q, r := divide(100, divisor)
	// handle the error
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)

}

func divide(x, y int) (int, int) {
	if y == 0 {
		// error
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
