/* synchronizing goroutines using sync.WaitGroup */

package main

import (
	"fmt"
	"sync"
	"time"
)

// create an instance of a waitgroup
var wg sync.WaitGroup // default value of the internal counter is 0

func main() {
	fmt.Println("main started")
	for i := 0; i < 10; i++ {
		wg.Add(1) // increment the counter by 1
		go f1()
	}
	f2()
	wg.Wait() // block the execution of the current function (main) until the counter to become zero
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter by 1 (always)
}

func f2() {
	fmt.Println("f2 invoked")
}
