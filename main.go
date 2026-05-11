package main

import (
	"fmt"
	"net/http"

	"alati-projekat/handler"
)

func main() {

	http.HandleFunc("/health", handler.Health)
	http.HandleFunc("/configs", handler.CreateConfig)
	fmt.Println("Server pokrenut na :8080")

	http.ListenAndServe(":8080", nil)

}
