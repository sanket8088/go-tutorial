package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sanket/mongoapi/router"
)

func main() {

	fmt.Println("Mongo db")
	r := router.Router()
	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":9000", r))
	fmt.Println("Listening on port 9000")

}
