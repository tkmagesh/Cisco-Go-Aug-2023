package main

import "fmt"

type MyStr string // MyStr is an alias for "string"

func (str MyStr) Length() int {
	return len(str)
}

func main() {
	s := MyStr("Aliquip exercitation laborum eiusmod dolore non sunt dolor. Aliqua tempor commodo mollit culpa est anim. Est incididunt amet adipisicing minim sit tempor sit amet aliqua nostrud mollit. Labore ullamco dolore dolor velit in est et duis cupidatat. Dolore id exercitation ut fugiat commodo Lorem tempor proident esse eu minim eiusmod excepteur.")
	fmt.Println(s.Length())
}
