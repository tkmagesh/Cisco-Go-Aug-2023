/* concurrent safe state management */

package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count :", count)
}

func increment() {
	mutex.Lock() // allow access to only one goroutine at any point of time
	{
		count += 1
	}
	mutex.Unlock()
}
