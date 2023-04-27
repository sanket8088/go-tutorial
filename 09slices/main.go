package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruit_list = []string{"Apple", "Tomato", "Peach"}
	fmt.Println(fruit_list)
	fmt.Printf("Type of fruit list is %T\n", fruit_list)

	fruit_list = append(fruit_list, "Mango", "Banana")
	fmt.Println(fruit_list)

	// fruit_list = append(fruit_list[1:])
	fruit_list = fruit_list[1:]
	fmt.Println(fruit_list)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
