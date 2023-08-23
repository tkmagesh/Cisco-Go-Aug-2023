package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	go f1() // schedule the execution of f1() through the scheduler
	f2()

	// the following ways of synchronizing goroutines are very primitive. One should not use these methods in their application

	time.Sleep(time.Second) // block the execution of the main() function for 1 second so that the scheduler can look for other goroutines that are scheduled and execute them
	// fmt.Scanln()

	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
