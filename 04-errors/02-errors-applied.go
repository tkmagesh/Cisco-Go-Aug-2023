package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divide by 0 error")

func main() {

	divisor := 0
	q, r, err := divide(100, divisor)
	// handle the error
	if err == DivideByZeroError {
		fmt.Println("Please do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("unknown error :", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)

}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divide by 0 error")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {

	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
