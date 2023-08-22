package main

import (
	"fmt"
	"time"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Entity struct {
	Name    string
	Created time.Time
}

type PerishableProduct struct {
	Product
	Entity
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string, entityName string, created time.Time) PerishableProduct {
	return PerishableProduct{
		Product: Product{Id: id, Name: name, Cost: cost},
		Expiry:  expiry,
		Entity:  Entity{Name: entityName, Created: created},
	}
}

func main() {
	// grapes := PerishableProduct{Product{100, "Grapes", 50}, "2 Days"}
	/*
		grapes := PerishableProduct{
			Product: Product{Id: 100, Name: "Grapes", Cost: 50},
			Expiry:  "2 Days",
			Entity:  Entity{Name: "Product", Created: time.Now()},
		}
	*/

	// simplify the above constructor logic using the constructor function
	grapes := NewPerishableProduct(100, "Grapes", 50, "2 Days", "Product", time.Now())
	fmt.Println(grapes)

	// accessing the attributes
	fmt.Println(grapes.Expiry)
	// fmt.Println(grapes.Product.Id, grapes.Product.Cost, grapes.Product.Name)
	fmt.Println(grapes.Id, grapes.Product.Name, grapes.Entity.Name, grapes.Cost)

}
