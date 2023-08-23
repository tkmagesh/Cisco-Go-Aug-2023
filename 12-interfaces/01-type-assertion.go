package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	x = "Mollit esse tempor adipisicing quis occaecat velit."
	x = true
	x = 199.99
	x = struct{}{}
	x = []int{10, 20, 30}
	fmt.Println(x)

	var data interface{}
	data = 100
	// fmt.Println(data + 200)
	// data = "Eiusmod ex magna incididunt commodo laborum."
	// fmt.Println(data.(int) + 200) // there is a possibility of panic
	if val, ok := data.(int); ok {
		fmt.Println(val + 200)
	} else {
		fmt.Println("data is not an int")
	}

	var z interface{}
	// z = 100
	// z = "Culpa ut deserunt pariatur ullamco voluptate."
	// z = true
	// z = 99.99
	z = []int{10, 20, 30}
	switch val := z.(type) {
	case int:
		fmt.Println("z is an int, z + 200 =", val+200)
	case string:
		fmt.Println("z is an string, len(z) =", len(val))
	case float64:
		fmt.Println("z is a float64, 90% of z = ", val*0.9)
	case bool:
		fmt.Println("z is a bool, !z =", !val)
	default:
		fmt.Println("z is of unknown type", val)
	}
}
