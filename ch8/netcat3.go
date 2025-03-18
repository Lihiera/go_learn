package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	_, err = io.Copy(conn, os.Stdin)
	fmt.Println("err is", err)
	tcpc := conn.(*net.TCPConn)
	tcpc.CloseWrite()
	<-done // wait for background goroutine to finish
	tcpc.Close()
}
