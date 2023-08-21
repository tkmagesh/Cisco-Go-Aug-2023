package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	//LOG ONLY
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	// PROFILE ONLY
	/*
		profiledAdd := getProfiledOperation(add)
		profiledAdd(100, 200)
	*/

	// PROFILE & LOG
	/*
		logAdd := getLogOperation(add)
		profiledLogAdd := getProfiledOperation(logAdd)
	*/

	profiledAdd := getProfiledOperation(add)
	profiledLogAdd := getLogOperation(profiledAdd)
	profiledLogAdd(100, 200)
}

func getProfiledOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed : ", elapsed)
	}
}

func getLogOperation(oper func(int, int)) func(int, int) {
	logOperation := func(x, y int) {
		log.Println("Operation started")
		oper(x, y)
		log.Println("Operation completed!")
	}
	return logOperation
}

// 3rd party library (you cannot change the implementation)
func add(x, y int) {
	fmt.Printf("Add Result : %d\n", x+y)
}

func subtract(x, y int) {
	fmt.Printf("Subtract Result : %d\n", x-y)
}
