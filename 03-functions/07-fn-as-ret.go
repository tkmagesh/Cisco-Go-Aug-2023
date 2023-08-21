package main

import "fmt"

func main() {
	fn := getFn()
	fn()
}

func getFn() func() {
	// return f1
	// return f2
	return func() {
		fmt.Println("anon fn invoked")
	}
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
