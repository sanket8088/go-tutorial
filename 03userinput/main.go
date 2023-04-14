package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to pizza app"
	fmt.Println(welcome)
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)

	// numRating := input + 1
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error input")
		fmt.Println(err)
	} else {
		fmt.Println("You rated", numRating+1)
	}

}
