package main

import "fmt"

func main() {

	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var getGreetMsg func(string) string
	getGreetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
	}
	greetMsg := getGreetMsg("Suresh")
	fmt.Print(greetMsg)

	/*
		result := func(x, y int) int {
			return x + y
		}(100, 200)
		fmt.Println(result)

		q, r := func(x, y int) (quotient int, remainder int) {
			quotient, remainder = x/y, x%y
			return
		}(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/
}
