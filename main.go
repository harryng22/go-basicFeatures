package main

import (
	"fmt"
	// "math/rand"
)

func main() {
	// fmt.Println("Value:", rand.Int())
	const price, tax float32 = 275.00, 27.5
	const quantity, inStock = 2, true
	fmt.Println("Total:", quantity*(price+tax))
	fmt.Println("In stock:", inStock)
}
