package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	strings := strings.Fields(scanner.Text())
	fmt.Println(strings)
	strings = removeSame(strings)
	fmt.Println(strings)
}

func removeSame(s []string) []string {
	length := len(s)
	for i, j := 0, 0; i < length-1; i++ {
		if s[j] == s[j+1] {
			s[j+1] = s[length-1]
			length--
		} else {
			j++
		}
	}
	return s[:length]
}
