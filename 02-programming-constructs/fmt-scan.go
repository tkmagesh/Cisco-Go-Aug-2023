package main

import "fmt"

func main() {
	/*
		fmt.Println("Enter a no:")
		var no int
		fmt.Scanln(&no)
	*/
	fmt.Println("Enter 2 numbers (separated by a space):")
	var no1, no2 int
	fmt.Scanln(&no1, &no2)
	fmt.Printf("%d + %d = %d\n", no1, no2, no1+no2)

}
