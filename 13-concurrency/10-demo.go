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
		// go fn(i) // "i" is resolved when the fn is "scheduled"

		go func(i int) {
			fn(i) // "i" is resolved when the fn is "executed"
			wg.Done()
		}(i)
	}
	wg.Wait() // block the execution of the current function (main) until the counter to become zero
	fmt.Println("main completed")
}

func fn(id int) {
	fmt.Printf("fn [%d] started\n", id)
	time.Sleep(3 * time.Second)
	fmt.Printf("fn [%d] completed\n", id)
	// wg.Done() // decrement the counter by 1 (always)
}
