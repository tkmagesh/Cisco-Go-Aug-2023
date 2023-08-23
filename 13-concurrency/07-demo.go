/* concurrent safe state management using a custom type*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count += 1
	}
	c.Unlock()
}

var counter Counter

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
	fmt.Println("count :", counter.count)
}

func increment() {
	counter.Increment()
}
