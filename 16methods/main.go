package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	sanket := User{"Sanket", "sanket@gmail.com", true, 26}
	fmt.Println(sanket)
	fmt.Printf("Sanket details are : %+v\n", sanket)
	fmt.Printf("Sanket age is : %v", sanket.Age)
	sanket.GetStatus()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}
