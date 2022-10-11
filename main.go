package main

import "fmt"

func main() {
	// fmt.Println("Value:", rand.Int())
	price, tax, quantity, inStock := 275.00, 27.5, 2.0, true
	fmt.Println("Total:", quantity*(price+tax))
	fmt.Println("In stock:", inStock)
	price = 300
	fmt.Println("Total:", quantity*(price+tax))
	fmt.Println("Price:", price)
}
