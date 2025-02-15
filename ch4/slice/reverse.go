package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a list of numbers ending with enter: ")
	scanner.Scan()
	s := scanner.Text()
	var numbers []int
	fields := strings.Fields(s)
	for _, field := range fields {
		number, _ := strconv.Atoi(field)
		numbers = append(numbers, number)
	}
	var arrayNumber [5]int
	for i := 0; i < 5; i++ {
		arrayNumber[i] = numbers[i]
	}
	//reverse(&arrayNumber)
	res := rotate(arrayNumber[:], 4)
	fmt.Println(res) // "[5 4 3 2 1 0]"
}

func reverse(ptr *[5]int) {
	for i, j := 0, 4; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func rotate(s []int, n int) []int {
	n = n % len(s)
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		k := i
		if k-n < 0 {
			k += len(s)
		}
		res[k-n] = s[i]
	}
	return res
}
