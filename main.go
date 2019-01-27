package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	log.Printf("Listening on %s...\n", port)
	http.ListenAndServe(port, nil)
}
