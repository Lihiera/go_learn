package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"
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
		_, err = io.WriteString(c, fmt.Sprintf("\r\n$ %s:", dir))

		command, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		paras := strings.Fields(string(command))
		if len(paras) == 1 {
			switch paras[0] {
			case "ls":
				var res []string
				files, _ := os.ReadDir("./")
				for i, file := range files {
					if i%2 == 1 {
						res = append(res, fmt.Sprintf("%s\t", file.Name()))
					} else {
						res = append(res, fmt.Sprintf("%s\r\n", file.Name()))
					}

				}
				io.WriteString(c, strings.Join(res, ""))
			case "close":
				return
			default:
				io.WriteString(c, "No such command")
			}
		} else if len(paras) == 2 {
			switch paras[0] {
			case "cd":
				if err := os.Chdir(filepath.Join(dir, paras[1])); err != nil {
					io.WriteString(c, fmt.Sprintln(err))
				}
			default:
				io.WriteString(c, "No such command")
			}

		} else if len(paras) == 3 {
			switch paras[0] {
			case "get":
			case "send":
			default:
				io.WriteString(c, "No such command")
			}

		} else {
			io.WriteString(c, "No such command")
		}
	}
}
