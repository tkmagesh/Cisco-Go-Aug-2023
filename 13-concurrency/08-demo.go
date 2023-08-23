/* concurrent safe state management using a custom type*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

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
	atomic.AddInt64(&count, 1)
}
