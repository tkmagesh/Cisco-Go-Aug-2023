package main

import "fmt"

func main() {
	sayHi()

	greet("Magesh")

	greetMsg := getGreetMsg("Suresh")
	fmt.Print(greetMsg)

	fmt.Println(add(100, 200))

	// fmt.Println(divide(100, 7))

	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	// Print ONLY quotient
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d \n", q)
	*/
}

/* function with no arguments and no return results */
func sayHi() {
	fmt.Println("Hi there!")
}

/* function with 1 argument and no return results */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/* function with 1 argument and 1 return result */
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
}

/* function with 2 arguments and 1 return result */
/*
func add(x int, y int) int {
	return x + y
}
*/
func add(x, y int) int {
	return x + y
}

/* function with 2 arguments and 2 return results */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/* functions with named results */
func divide(x, y int) (quotient int, remainder int /* quotient & remainder are declared and initialized with default value of int (0) */) {
	quotient, remainder = x/y, x%y
	return
}
