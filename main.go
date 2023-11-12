package main

import (
	"fmt"
	"net/http"

	"github.com/Gabrielm3/url-shortner/handlers"
)

func main() {
	// function main.
	http.HandleFunc("/", handlers.HandleForm)
	http.HandleFunc("/shortn", handlers.HandleShorten)
	http.HandleFunc("/short/", handlers.HandleRedirect)

	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}
