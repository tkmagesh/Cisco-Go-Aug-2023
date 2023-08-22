package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func main() {
	cisco := &Organization{
		Name: "Cisco",
		City: "Bengaluru",
	}

	e1 := Employee{
		Id:   100,
		Name: "Magesh",
		Org:  cisco,
	}

	e2 := Employee{
		Id:   200,
		Name: "Suresh",
		Org:  cisco,
	}

	fmt.Printf("%+v\n", e1)
	fmt.Printf("%+v\n", e2)

	cisco.City = "Pune"

	fmt.Println("Organization")
	fmt.Printf("%+v\n", cisco)

	fmt.Println("Employees ")
	fmt.Printf("%v - Org: %v\n", e1, e1.Org)
	fmt.Printf("%v - Org: %v\n", e2, e2.Org)
}
