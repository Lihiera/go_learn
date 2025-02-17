package main

import (
	"fmt"
)

func main() {
	list := []string{"af", "br", "2c", "34"}
	spring1 := fmt.Sprintf("%q", list)
	fmt.Println(string(spring1[1]))
	fmt.Printf("%T\n", spring1)
	fmt.Println(list)
}
