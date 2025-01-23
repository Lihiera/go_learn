package main

import (
	"fmt"
	"os"
)

func main() {
	dir := os.Args[1]
	fmt.Println(dir)
	fmt.Println(len(dir))
	fmt.Println(basename(dir))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			fmt.Println(i)
			s = s[i+1:]
			break
		}
	}
	fmt.Println()
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			fmt.Println(i)
			s = s[:i]
			break
		}
	}
	fmt.Println()
	return s
}
