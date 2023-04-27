package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 18
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if (loginCount > 10) && (loginCount < 20) {
		result = "Greater"
	} else {
		result = "Final answr here"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Number less than 10")
	} else {
		fmt.Println("Num is greater")
	}

}
