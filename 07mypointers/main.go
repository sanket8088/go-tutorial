package main

import "fmt"

func main() {
	fmt.Println("Pointer session")

	// var ptrInt *int
	// fmt.Println("Value of ptrInt is ", ptrInt)

	myNumber := 23

	var memAddress = &myNumber

	fmt.Println("Value of actual pointer is", memAddress)
	fmt.Println("Value of actual pointer is", *memAddress)

	*memAddress = *memAddress * 2

	fmt.Println("Multiplied value is", *memAddress)

}
