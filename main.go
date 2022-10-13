package main

import "fmt"

func main() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	allNames := names[1:]
	someNames := make([]string, 2)
	someNames = append(someNames, allNames...)
	someNames[1]="Boots"

	fmt.Println(allNames)
	fmt.Println(someNames)
	fmt.Println(cap(allNames), len(allNames))
	fmt.Println(cap(someNames), len(someNames))

	// names := make([]string, 3, 6)

	// names[0] = "Kayak"
	// names[1] = "LifeJacket"
	// names[2] = "Paddle"

	// moreNames := []string{"Hat", "Gloves"}

	// appendedNames := append(names, moreNames...)
	// names[0] = "Canoe"

	// fmt.Println(len(names))
	// fmt.Println(cap(names))
	// fmt.Println(names)
	// fmt.Println(appendedNames)
	// fmt.Println(moreNames)
}
