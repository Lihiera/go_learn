package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "foo, foo, foofoo"
	fmt.Println(s)
	s = expand(s, exchange)
	fmt.Println(s)
}

func expand(s string, f func(string) string) string {

	index := strings.Index(s, "foo")
	if index == -1 {
		return s
	}
	return s[:index] + f("foo") + expand(s[index+3:], f)
}

func exchange(s string) string {
	return strings.ToUpper(s) + "foo"
}
