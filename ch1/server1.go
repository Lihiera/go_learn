package main

import (
	"fmt"
	"log"
	"net/http"

	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		cycles := 5
		for k, v := range r.Form {
			if k == "cycles" {
				cycles, _ = strconv.Atoi(v[0])
			}
		}
		lissajous(w, cycles)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	count++
// 	mu.Unlock()
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %T\n", k, v[0])
	}
}
