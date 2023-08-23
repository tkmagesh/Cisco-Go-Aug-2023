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
	wg := sync.WaitGroup{}
	for _, pair := range data {
		wg.Add(1)
		go func(pair []int) {
			add(pair[0], pair[1])
			wg.Done()
		}(pair)
	}
	wg.Wait()
	fmt.Println(total)

	/*
		Sum all the numbers in "data" concurrently
		   Note : spin one goroutine (add) for each pair of numbers
	*/
}

func add(x, y int) {
	fmt.Printf("x = %d and y = %d\n", x, y)
	sum := x + y
	mutex.Lock()
	{
		total += sum
	}
	mutex.Unlock()
}
