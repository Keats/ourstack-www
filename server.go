package main

import (
	"fmt"
	"net/http"

	"./ourstack"
)

func main() {
	fmt.Println("Starting server...")
	ourstack.GetDBConnection("postgres://testing:password@localhost/ourstack?sslmode=disable")

	http.HandleFunc("/", ourstack.HomeHandler)
	http.ListenAndServe(":7777", nil)

}
