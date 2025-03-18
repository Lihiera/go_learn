package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/user"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	currentUser, err := user.Current()
	if err != nil {
		io.WriteString(c, fmt.Sprintln(err))
		return
	}
	if err := os.Chdir(currentUser.HomeDir); err != nil {
		io.WriteString(c, fmt.Sprintln(err))
		return
	}
	reader := bufio.NewReader(c)
	for {
		dir, err := os.Getwd()
		if err != nil {
			io.WriteString(c, fmt.Sprintln(err))
			return
		}
		_, err = io.WriteString(c, fmt.Sprintf("$ %s:", dir))

		command, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		paras := strings.Fields(string(command))
		if len(paras) == 1 {
			switch paras[0] {
			case "cd":
			case "ls":
			case "close":
				return
			default:
			}
		}
	}
}
