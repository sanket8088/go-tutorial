package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	sanket := User{"Sanket", "sanket@gmail.com", true, 26}
	fmt.Println(sanket)
	fmt.Printf("Sanket details are : %+v\n", sanket)
	fmt.Printf("Sanket age is : %v", sanket.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
