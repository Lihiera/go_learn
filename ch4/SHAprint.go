package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hash = flag.String("hash", "sha256", "hash type")

func main() {
	flag.Parse()
	s := []byte(flag.Args()[0])
	switch *hash {
	case "sha256":
		fmt.Println(sha256.Sum256(s))
	case "sha384":
		fmt.Println(sha512.Sum384(s))
	case "sha512":
		fmt.Println(sha512.Sum512(s))
	default:
		fmt.Println("Invalid hash type!")
	}
}
