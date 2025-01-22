package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var counts = make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		fmt.Println(filename, "\n", string(data), "\n")
		for _, lines := range strings.Split(string(data), "\n") {
			counts[lines]++
		}
	}

	for str, n := range counts {
		if n > 1 {
			fmt.Printf("%v个%v\n", n, str)
		}
	}
}
