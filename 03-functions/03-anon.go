package main

import "fmt"

/* anonymous functions
- functions that are created and invoked
- do not have a name
*/

func main() {

	func() {
		fmt.Println("Hi there!")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	greetMsg := func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
	}("Suresh")
	fmt.Print(greetMsg)

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(result)

	q, r := func(x, y int) (quotient int, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
