package main

import "fmt"

func main() {
	names := [3]string {"Kayak", "LifeJacket","Paddle"}

	otherArray := &names
	
	names[0]="safd"

	fmt.Println(names)
	fmt.Println(otherArray)
}
