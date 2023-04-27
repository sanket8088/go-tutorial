package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=fdgsdfag"

func main() {
	fmt.Println("Welcome to hadnling URL")
	fmt.Println(myUrl)

	// parsing url
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Query()["coursename"][0])

}
