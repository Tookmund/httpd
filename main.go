package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":8080"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	fs := http.FileServer(http.Dir("."))
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s %s\n",req.Method, req.URL, req.Proto)
		fs.ServeHTTP(rw, req)
	})
	ip, _ := externalIP()
	if ip != "" {
		log.Printf("External IP: %s\n", ip)
	}
	log.Printf("Listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
