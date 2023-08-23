package main

import (
	"fmt"
	"strings"
)

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

type Products []Product

func (products Products) filter(predicate func(Product) bool) Products {
	var result Products
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

/*
func (products Products) String() string {
	var result string
	for _, product := range products {
		result += fmt.Sprintln(product.String())
	}
	return result
}
*/

// using strings.Builder
func (products Products) String() string {
	var builder strings.Builder
	for _, product := range products {
		builder.WriteString(fmt.Sprintf("%s\n", product.String()))
	}
	return builder.String()
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	var costlyProductPredicate = func(product Product) bool {
		return product.Cost > 1000
	}

	costlyProducts := products.filter(costlyProductPredicate)
	fmt.Println("Costly Products")
	/* for _, product := range costlyProducts {
		fmt.Println(product.String())
	} */
	fmt.Println(costlyProducts.String())

	/*
		var stationaryProductPredicate = func(product Product) bool {
			return product.Category == "Stationary"
		}
		stationaryProducts := products.filter(stationaryProductPredicate)
	*/

	stationaryProducts := products.filter(func(product Product) bool {
		return product.Category == "Stationary"
	})
	fmt.Println("Stationary Products")
	fmt.Println(stationaryProducts.String())

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
