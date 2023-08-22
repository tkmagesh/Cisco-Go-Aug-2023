package calc

import "fmt"

// private to the package (NOT file)
// var opCount int
var opCount map[string]int

func init() {
	opCount = make(map[string]int)
	fmt.Println("calc package initialized - [calc.go]")
}

func OpStats() map[string]int {
	return opCount
}
