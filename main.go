package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	log.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
