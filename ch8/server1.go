package main

import (
	"io"
	"log"
	"net"
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
			log.Print(err)
			continue
		}
		go handler1(conn)
	}
}

func handler1(c net.Conn) {
	defer c.Close()
	io.Copy(c, c)
	time.Sleep(3 * time.Second)
	io.WriteString(c, "over\n")
}
