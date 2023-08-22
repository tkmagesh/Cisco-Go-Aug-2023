package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

// receive a "value" when the state need not be changed
func (p Product) toString() string {
	return fmt.Sprintf("id = %d, name = %q, cost = %.2f", p.Id, p.Name, p.Cost)
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

// inherited type
type PerishableProduct struct {
	Product
	Expiry string
}

// overriding the Product's toString() method
func (pp PerishableProduct) toString() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.toString(), pp.Expiry)
}

func main() {
	grapes := PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Grapes",
			Cost: 50,
		},
		Expiry: "2 Days",
	}

	fmt.Println(grapes.toString())
	fmt.Println("After applying 10% discount")
	grapes.applyDiscount(10)
	fmt.Println(grapes.toString())
}
