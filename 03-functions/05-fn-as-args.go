package main

import "fmt"

func main() {
	/*
		exec("F1")
		exec("f2")
		exec("f3")
	*/

	exec(f1)
	exec(f2)

}

/*
func exec(fnName string) {
	// execute either f1 or f2 based on the given argument
	switch fnName {
	case "f1":
		f1()
	case "f2":
		f2()
	default:
		fmt.Println("invalid function name")
	}
}
*/

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
