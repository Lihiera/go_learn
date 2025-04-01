package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	done := make(chan struct{})
	timer := time.NewTimer(10 * time.Second)
	go func() {
		input := bufio.NewScanner(c)
		var wg sync.WaitGroup
		for input.Scan() {
			timer.Reset(10 * time.Second)
			wg.Add(1)
			go func() {
				defer wg.Done()
				echo(c, input.Text(), 1*time.Second)
			}()
		}
		wg.Wait()
		done <- struct{}{}
	}()
	select {
	case <-timer.C:
		fmt.Fprintf(c, "Connection closed due to inactivity in 10 seconds\n")
		timer.Stop()
	case <-done:
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
