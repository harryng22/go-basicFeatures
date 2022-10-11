package main

import (
	"fmt"
	// "math/rand"
)

func main() {
	// fmt.Println("Value:", rand.Int())
	const price float32 = 275.00
	const tax float32 = 27.5
	fmt.Println(price + tax)
}
