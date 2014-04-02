package main

import (
	"fmt"
	"net/http"

	"./ourstack"
)

func main() {
	fmt.Println("Starting server...")
	ourstack.GetDBConnection("postgres://testing:password@localhost/ourstack?sslmode=disable")

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/techs", ourstack.TechHandler)
	http.HandleFunc("/", ourstack.IndexHandler)

	http.ListenAndServe(":7777", nil)
}
