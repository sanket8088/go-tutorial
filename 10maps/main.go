package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["py"] = "Python"
	languages["js"] = "Javascript"
	languages["jar"] = "Java"

	fmt.Println("List of languages", languages)
	fmt.Println("py is short for", languages["py"])

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Println("key", key)
		fmt.Println("key", key)
		fmt.Println("value", value)
	}

}
