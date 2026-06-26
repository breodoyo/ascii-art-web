package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art" AsciiHandler)

	log.Println("Server running on : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}