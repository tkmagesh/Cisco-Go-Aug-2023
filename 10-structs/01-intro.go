package main

import "fmt"

func main() {
	/*
		var product struct {
			id   int
			name string
			cost float32
		}
		product.id = 100
		product.name = "Pen"
		product.cost = 20
	*/

	/*
		var product struct {
			id   int
			name string
			cost float32
		} = struct {
			id   int
			name string
			cost float32
		}{
			id:   100,
			name: "pen",
			cost: 20,
		}
	*/

	//type inference
	product := struct {
		id   int
		name string
		cost float32
	}{
		id:   100,
		name: "pen",
		cost: 20,
	}

	// fmt.Printf("%#v\n", product)
	printProduct(product)
}

func printProduct(p struct {
	id   int
	name string
	cost float32
}) {
	fmt.Printf("id = %d, name = %q, cost = %.2f\n", p.id, p.name, p.cost)
}
