package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8000, "The server's port.")
var location = flag.String("loc", "Local", "Your location.")

func main() {
	flag.Parse()
	server := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn1(conn) // handle one connection at a time
	}
}

func handleConn1(c net.Conn) {
	defer c.Close()
	for {
		utc := time.Now().UTC()
		loc, _ := time.LoadLocation(*location)
		Time := utc.In(loc)
		_, err := io.WriteString(c, Time.Format("2006-01-02 15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
