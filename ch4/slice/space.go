package main

import (
	"fmt"
)

func main() {
	strings := "412 5  0otasd 3245  213   "
	strings = removeSpace([]byte(strings))
	fmt.Println(strings)
}

func removeSpace(s []byte) string {
	length := len(s)
	for i := 0; i < length-1; i++ {
		if s[i] == ' ' {
			if s[i+1] == ' ' {
				fmt.Println("found")
				copy(s[i:], s[i+1:])
				length--
			}
		}
	}
	return string(s[:length])
}
