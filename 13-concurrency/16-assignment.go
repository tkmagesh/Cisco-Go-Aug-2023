/*
get rid of "communicate by sharing memory" strategy (total variable)
use channels and rewrite the program
*/
package main

import (
	"fmt"
	"sync"
)

var total int
var mutex sync.Mutex

func main() {
	data := [][]int{
		{10, 20},
		{30, 20},
		{20, 40},
		{30, 50},
		{50, 60},
		{70, 50},
		{90, 80},
		{80, 70},
		{80, 60},
		{30, 60},
	}
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	for _, pair := range data {
		wg.Add(1)
		go add(pair[0], pair[1], ch, wg)
	}
	resultCh := aggregate(ch)
	wg.Wait()
	close(ch)
	fmt.Println(<-resultCh)

}

func add(x, y int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("x = %d and y = %d\n", x, y)
	sum := x + y
	ch <- sum
}

func aggregate(ch chan int) chan int {
	resultCh := make(chan int)
	go func() {
		result := 0
		for {
			if sum, isOpen := <-ch; isOpen {
				result += sum
				continue
			}
			break
		}
		resultCh <- result
	}()
	return resultCh
}
