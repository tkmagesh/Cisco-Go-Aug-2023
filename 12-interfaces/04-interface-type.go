package main

import (
	"fmt"
	"math"
)

// v1.0
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

// v2.0
type Rectangle struct {
	Length  float32
	Breadth float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Breadth)
}

// v3.0
/*
func printArea(x interface{}) {
	// fmt.Println("Area :", x.Area())
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Type incompatible with Area() behavior")
	}
}
*/

/*
func printArea(x interface{}) {
	// fmt.Println("Area :", x.Area())
	switch val := x.(type) {
	case interface{ Area() float32 }:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Type incompatible with Area() behavior")
	}
}
*/

type AreaFinder interface {
	Area() float32
}

func printArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

type PerimeterFinder interface {
	Perimeter() float32
}

func printPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

// v4.0
func printShape(x ShapeStatsFinder) {
	printArea(x)      //=> x should be interface{ Area() float32 }
	printPerimeter(x) //=> x should interface{ Perimeter() float32 }
}

func main() {
	c := Circle{12}
	// fmt.Println("Area :", c.Area())
	/*
		printArea(c)
		printPerimeter(c)
	*/
	printShape(c)

	r := Rectangle{10, 12}
	// fmt.Println("Area :", r.Area())
	/*
		printArea(r)
		printPerimeter(r)
	*/
	printShape(r)

}
