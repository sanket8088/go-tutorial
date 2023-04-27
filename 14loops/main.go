package main

import "fmt"

func main() {
	fmt.Println("Loops running")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// For loop

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and the day is %v\n", index, day)
	}

	rgValue := 1

	for rgValue < 10 {

		if rgValue == 2 {
			goto lco
		}
		fmt.Println("Value is ", rgValue)
		rgValue++
	}

lco:
	fmt.Println("Jumping at lco.in")

}
