package main

import "fmt"

func main() {
	var no int
	no = 100

	var noPtr *int
	noPtr = &no // address of no
	fmt.Println(noPtr)

	// deferencing = address -> no
	fmt.Println(*noPtr)

	fmt.Println(no == *(&no))

	fmt.Println("Before incrementing, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping n1 = %d and n2 = %d\n", n1, n2)
}

func increment(x *int) {
	fmt.Println("[increment], x =", x)
	*x += 1
}

func swap(x, y *int) /* DO NOT return anything */ {
	*x, *y = *y, *x
}
