package main

import (
	"log"
	"net/http"

	"example/frameworks"
)

func main() {
	http.HandleFunc("/", frameworks.Router)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
