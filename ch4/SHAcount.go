package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	s1, s2 := os.Args[1], os.Args[2]
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	fmt.Println(SHAcount(&c1, &c2))
}

func SHAcount(c1, c2 *[32]byte) int {
	count := 0
	for i := 0; i < len(c1); i++ {
		count += CountBit(c1[i], c2[i])
	}
	return count
}

func CountBit(b1, b2 byte) int {
	count := 0
	for i := 0; i < 8; i++ {
		if b1&1 != b2&1 {
			count++
		}
		b1 >>= 1
		b2 >>= 1
	}
	return count
}
