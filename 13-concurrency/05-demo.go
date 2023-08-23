/*
Use htop to monitor the resouce usage with varying # of goroutines passed as command line arguments
*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "Number of goroutines to run")
	flag.Parse()

	fmt.Printf("main started - with %d goroutines\n", count)
	fmt.Println("Hit ENTER to start....")
	fmt.Scanln()
	wg := &sync.WaitGroup{}
	for i := 1; i <= count; i++ {
		wg.Add(1) // increment the counter by 1
		go fn(i, rand.Intn(20), wg)
	}

	wg.Wait() // block the execution of the current function (main) until the counter to become zero
	fmt.Println("main completed")
	fmt.Println("Hit ENTER to shutdown....")
	fmt.Scanln()
}

func fn(id int, delay int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn [%d] started\n", id)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Printf("fn [%d] completed\n", id)

}
