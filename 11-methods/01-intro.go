package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

// receive a "value" when the state need not be changed
func (p Product) printProduct() {
	fmt.Printf("id = %d, name = %q, cost = %.2f\n", p.Id, p.Name, p.Cost)
}

// receive a "pointer" when the state has to be modified
func (p *Product) applyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

func NewProduct(id int, name string, cost float32) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}
func main() {
	/*
		pen := Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/
	pen := NewProduct(100, "Pen", 10)

	// printProduct(pen)
	pen.printProduct()
	fmt.Println("After applying 10% discount")
	pen.applyDiscount(10)
	pen.printProduct()
}
