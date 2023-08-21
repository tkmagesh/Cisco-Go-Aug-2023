package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		log.Println("Operation started")
		add(100, 200)
		log.Println("Operation completed!")

		log.Println("Operation started")
		subtract(100, 200)
		log.Println("Operation completed")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
}

func logOperation(oper func(int, int), x, y int) {
	log.Println("Operation started")
	oper(x, y)
	log.Println("Operation completed!")
}

/*
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed!")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}
*/

// 3rd party library (you cannot change the implementation)
func add(x, y int) {
	fmt.Printf("Add Result : %d\n", x+y)
}

func subtract(x, y int) {
	fmt.Printf("Subtract Result : %d\n", x-y)
}
