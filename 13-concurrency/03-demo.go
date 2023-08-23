/* synchronizing goroutines using sync.WaitGroup */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	for i := 0; i < 10; i++ {
		wg.Add(1) // increment the counter by 1
		go f1(wg)
	}
	f2()
	wg.Wait() // block the execution of the current function (main) until the counter to become zero
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")

}

func f2() {
	fmt.Println("f2 invoked")
}
