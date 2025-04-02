package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	mu    sync.Mutex
)

func main() {
	ping := make(chan struct{})
	pong := make(chan struct{})
	timer := time.Tick(1 * time.Second)
	go func() {
		for {
			ping <- struct{}{}
		}
	}()
	go func() {
		for {
			pong <- struct{}{}
		}
	}()
	for {
		select {
		case <-ping:
			count++
		case <-pong:
			count++
		case <-timer:
			fmt.Println(count)
		}
	}

}
