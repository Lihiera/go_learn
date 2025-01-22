package main

import (
	"fmt"
)

func main() {
	b := "\xe4\xb8\x96"
	c := "\xe7\x95\x8c"
	d := "\u4e16"
	e := "\u754c"
	fmt.Printf("%x, %x, %x, %x\n", []byte(b), []byte(c), []byte(d), []byte(e))
}
