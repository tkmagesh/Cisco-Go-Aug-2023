/* Using channels for streaming data */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go generateNos(ch)
	/*
		for {
			time.Sleep(500 * time.Millisecond)
			if data, isOpen := <-ch; isOpen {
				fmt.Println(data)
				continue
			}
			fmt.Println("[consumer] all data received.. channel is closed... exiting")
			break
		}
	*/

	//implementing the above using for..range
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("[consumer] all data received.. channel is closed... exiting")
}

// producer
func generateNos(ch chan int) {
	count := rand.Intn(20)
	for i := 1; i < count; i++ {
		ch <- i * 10
	}
	fmt.Println("[producer] data generation completed.. closing the channel")
	close(ch)
}
