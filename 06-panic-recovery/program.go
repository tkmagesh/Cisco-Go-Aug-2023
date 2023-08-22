package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("Divide by 0 error")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("app panicked : ", err)
			return
		}
		fmt.Println("Thank you")
	}()
	divisor := 0
	if divisor == 0 {
		panic(DivideByZeroError)
	}
	result := 100 / divisor
	fmt.Println(result)
}
