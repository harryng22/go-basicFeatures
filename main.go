package main

import "fmt"

func main() {
	// fmt.Println("Value:", rand.Int())
	var price, tax = 275.00, 27.5
	const quantity, inStock = 2, true
	fmt.Println("Total:", quantity*(price+tax))
	fmt.Println("In stock:", inStock)
	price = 300
	fmt.Println("Total:", quantity*(price+tax))
	fmt.Println("Price:", price)
}
