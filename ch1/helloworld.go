package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var sep, s string
	strList := make([]string, 0, 100)
	for i := 0; i < 1000; i++ {
		strList = append(strList, strconv.Itoa(i))
	}

	start1 := time.Now()
	for _, arg := range strList {
		s += sep + arg
		sep = " "
	}
	elapse1 := time.Since(start1).Nanoseconds()
	s = ""
	start2 := time.Now()
	s = strings.Join(strList, " ")
	elapse2 := time.Since(start2).Nanoseconds()
	fmt.Println(elapse1, elapse2)
}
