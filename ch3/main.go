package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/surface", handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	man(w)
}
