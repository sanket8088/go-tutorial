package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "mango"
	fruitList[1] = "apple"
	fruitList[2] = "banana"
	fruitList[3] = "peach"

	fmt.Println("Fruit list", fruitList)
	fmt.Println("Fruit list", len(fruitList))

	var vegetableList = [3]string{"potato", "beans", "mushrooms"}

	fmt.Println("Veggie list", vegetableList)

}
