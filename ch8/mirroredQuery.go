package main

import "context"

var done chan struct{}

func mirroredQuery() string {
	responses := make(chan string, 3)
	var response string
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	response = <-responses
	close(done)
	return response // return the quickest response
}

func fetch(hostname string) (response string) {
	cancelCtx, cancel := context.WithCancel(context.Background())
	go func() {
		<-done
		cancel()
	}()
	/*...*/
	return "ok"
}

func request(hostname string) (response string) { /* ... */ }
