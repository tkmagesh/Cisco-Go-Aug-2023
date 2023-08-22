package calc

import "fmt"

func init() {
	fmt.Println("calc package initialized - [add.go]")
}

// public
func Add(x, y int) int {
	opCount["add"] += 1
	return x + y
}
