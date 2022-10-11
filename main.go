package main

import "fmt"

func main() {
	first := 100
	second := &first

	first++
	*second++

	var myNewPointer *int
	fmt.Println(myNewPointer)
	myNewPointer = second
	*myNewPointer++

	fmt.Println("First:", first)
	fmt.Println("Second:", *second)
}
