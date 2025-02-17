package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()
	in := bufio.NewScanner(file)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		word := in.Text()
		counts[word]++
	}

	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
