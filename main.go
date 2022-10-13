package main

import "fmt"

func main() {
	// names := []string{"Kayak", "Lifejacket", "Paddle"}

	// appendedNames := append(names, "Hat", "Gloves")
	// names[0] = "Canoe"

	// fmt.Println(names)
	// fmt.Println(appendedNames)

	names := make([]string, 3, 6)

	names[0] = "Kayak"
	names[1] = "LifeJacket"
	names[2] = "Paddle"

	moreNames := []string{"Hat", "Gloves"}

	appendedNames := append(names, moreNames...)
	names[0] = "Canoe"

	fmt.Println(len(names))
	fmt.Println(cap(names))
	fmt.Println(names)
	fmt.Println(appendedNames)
	fmt.Println(moreNames)
}
