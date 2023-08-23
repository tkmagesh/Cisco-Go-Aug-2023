/* using select..case for signaling */

package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	fibCh := genFib(stopCh)
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		stopCh <- struct{}{}
	}()
	for no := range fibCh {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func genFib(stopCh <-chan struct{}) <-chan int {
	fibCh := make(chan int)

	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
				fmt.Println("Stop signal received")
				break LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				fibCh <- x
				x, y = y, x+y
			}
		}
		close(fibCh)
	}()
	return fibCh
}
