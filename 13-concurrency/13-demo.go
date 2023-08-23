/* understanding channel behaviors */
package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)
}

/*
step-1. goroutine (line:8) is scheduled for future execution
step-2. line:11, a receive operation is initiated (opening the gate) and blocked until the data arrives
step-3. since the line:11 is blocked, the scheduler wil look for other goroutines scheduled and execute them (line:8)
step-4. goroutine (line:8) is executed, since the channel receive operation is already initiated (gate is already open) (in step:2), the send operation (line:9) will be a non-blocking operation and the goroutine completes its execution after sending the data to the channel
step-5. the control returns back to line:11, the receive operation becomes un-blocked as the data was already sent (line:9)
step-6. receive operation is successful

*/
