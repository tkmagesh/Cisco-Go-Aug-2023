package main

import "fmt"

func main() {
	var intVar int32 = 100
	var floatVar float32

	floatVar = float32(intVar) // syntax for type conversion

	fmt.Println(floatVar)
}
