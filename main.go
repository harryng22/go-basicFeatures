package main

import "fmt"

func main() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	deleted := append(names[:2], names[3:]...)
	
	fmt.Println(deleted)

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
