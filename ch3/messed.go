package main

import (
	"fmt"
	"os"
)

func main() {
	s1 := os.Args[1]
	s2 := os.Args[2]
	fmt.Println(messed(s1, s2))
}

func messed(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	compare := make(map[rune]int)
	for _, r := range s1 {
		compare[r]++
	}
	for _, r := range s2 {
		compare[r]--
		if compare[r] < 0 {
			return false
		}
	}
	return true
}
