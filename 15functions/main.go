package main

import "fmt"

func main() {
	fmt.Println("Values here")
	greeter()
	greeter2()
	fmt.Println(adder(2, 3))
	fmt.Println(proAdder(2, 3, 4, 5))

}

func greeter2() {
	fmt.Println("Another method")
}

func greeter() {
	fmt.Println("Namaste yfrom golang")
}

func adder(num1 int, num2 int) int {
	return num1 + num2

}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val

	}
	return total
}
