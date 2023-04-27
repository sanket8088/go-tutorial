package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://official-joke-api.appspot.com/random_joke"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error alert", err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error alert", err)
		panic(err)
	}

	fmt.Println("Body is", string(body))
}
