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
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s %s\n",req.Method, req.URL, req.Proto)
		fs.ServeHTTP(rw, req)
	})
	log.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
