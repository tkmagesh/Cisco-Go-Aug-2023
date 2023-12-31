/* communication between goroutines using "Share memory by communicating" strategy */

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("result :", result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
