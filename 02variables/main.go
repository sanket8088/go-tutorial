package main

import "fmt"

func main() {
	var username string = "sanket8088"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isVerified bool = true
	fmt.Println(isVerified)
	fmt.Printf("Variable is of type: %T \n", isVerified)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	bigVal := 2889
	fmt.Println("Value is", bigVal)

}
