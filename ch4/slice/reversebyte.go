package main

import (
	"fmt"
)

func main() {
	bytes := []byte("Hello, 世界, 我已经醒了")
	bytes = reverseByte(bytes)
	fmt.Println(string(bytes))
}

func reverseByte(s []byte) []byte {
	for i := 0; i < len(s); {
		switch {
		case s[i] <= 0x7F:
			i++
		case s[i] >= 0xC0 && s[i] <= 0xDF:
			s[i], s[i+1] = s[i+1], s[i]
			i += 2
		case s[i] >= 0xE0 && s[i] <= 0xEF:
			s[i], s[i+2] = s[i+2], s[i]
			i += 3
		case s[i] >= 0xF0 && s[i] <= 0xF7:
			s[i], s[i+3] = s[i+3], s[i]
			s[i+1], s[i+2] = s[i+2], s[i+1]
			i += 4
		}
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
