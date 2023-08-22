package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func main() {
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	/*
		Write the following APIs for the products collection:
			- filter
				example use cases:
					1. filter costly products (cost > 1000)
					OR
					2. filter stationary products
					etc
			- groupBy
				example use cases:
					1. group products by category
					OR
					2. group products by cost (costly & affordable)
	*/
}
