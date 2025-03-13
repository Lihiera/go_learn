package main

import (
	"bufio"
	"fmt"
	"strings"
)

type counter int

type myWriter int

func (c *counter) wordCount(words string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(words))
	scanner.Split(bufio.ScanWords)
	ret := 0
	for scanner.Scan() {
		*c++
		ret++
	}
	return ret, nil
}

func (c *counter) lineCount(words string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(words))
	ret := 0
	for scanner.Scan() {
		*c++
		ret++
	}
	return ret, nil
}

func main() {
	var c counter
	c.wordCount("hello world")
	fmt.Println(c)
	c.wordCount("hello world secondly")
	fmt.Println(c)
	c.lineCount("hello\nworld\n\n\n")
	fmt.Println(c)
}
