package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]
	fmt.Println(comma(input))
}

func comma(s string) string {
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		s = s[:dot]
	}
	for i := len(s); i > 3; i -= 3 {
		s = s[:i-3] + "," + s[i-3:]
	}
	return s
}
