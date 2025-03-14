package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	walls := os.Args[1:]
	var parts [][]string
	var clocks []string
	var conns []io.Reader
	for i, _ := range walls {
		parts = append(parts, strings.Split(walls[i], "="))
	}
	for _, part := range parts {
		conn, err := net.Dial("tcp", part[1])
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		conns = append(conns, conn)
	}
	for {
		clocks = nil
		for i, part := range parts {
			buf := make([]byte, 1024)
			conns[i].Read(buf)

			clocks = append(clocks, fmt.Sprintf("%s: %s", part[0], string(buf)))
		}
		fmt.Println(clocks)
	}
}

// func display(tz string, url string) string {

// 	conn, err := net.Dial("tcp", url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer conn.Close()
// 	buf := make([]byte, 1024)
// 	conn.Read(buf)
// 	return fmt.Sprintf("%s: %s", tz, string(buf))

// 	// if _, err := io.Copy(os.Stdout, conn); err != nil {
// 	// 	log.Fatal(err)
// 	// }
// }
