package calc

import "fmt"

func init() {
	fmt.Println("calc package initialized - [subtract.go]")
}

// public
func Subtract(x, y int) int {
	opCount["subtract"] += 1
	return x - y
}
