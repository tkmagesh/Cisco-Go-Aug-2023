/* understanding channel behaviors */
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Println("len(ch) :", len(ch))
	ch <- 100
	fmt.Println("len(ch) :", len(ch))
	ch <- 200
	fmt.Println("len(ch) :", len(ch))

	ch <- 300 // block forever.. leading to a deadlock
	fmt.Println(<-ch)
	fmt.Println("len(ch) :", len(ch))

	fmt.Println(<-ch)
	fmt.Println("len(ch) :", len(ch))

}
