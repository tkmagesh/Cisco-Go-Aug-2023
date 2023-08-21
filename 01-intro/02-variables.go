package main

import "fmt"

// var dummy int
// var dummy int = 100
// var dummy = 100
// dummy := 100 // NOT ALLOWED

func main() {
	/*
		var name string
		name = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	/*
		var name string = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	// type inference
	/*
		var name = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	/*
		name := "Magesh" // := is allowed ONLY in function scope (NOT in package scope)
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	//multiple variables
	/*
		var x int
		var y int
		var result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var result int = x + y
		var str string = "Sum of %d and %d is %d\n"
		fmt.Printf(str, x, y, result)
	*/

	// multi-assignment
	/*
		var x, y int = 100, 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	// multi-assignment + type inference
	/*
		var (
			x, y   = 100, 200
			str    = "Sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)
}
