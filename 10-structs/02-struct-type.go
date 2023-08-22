package main

import "fmt"

type Product struct {
	id   int
	name string
	cost float32
}

func main() {
	/*
		var product Product
		product.id = 100
		product.name = "Pen"
		product.cost = 20
	*/

	/*
		var product Product = Product {
			id:   100,
			name: "pen",
			cost: 20,
		}
	*/

	//type inference
	/*
		product := Product{
			id:   100,
			name: "pen",
			cost: 20,
		}
	*/

	product := Product{100, "pen", 10} //=> not advisable

	// fmt.Printf("%#v\n", product)
	printProduct(product)

	//equality
	p1 := Product{200, "Pencil", 5}
	p2 := Product{200, "Pencil", 5}
	fmt.Println(p1 == p2)

}

func printProduct(p Product) {
	fmt.Printf("id = %d, name = %q, cost = %.2f\n", p.id, p.name, p.cost)
}
