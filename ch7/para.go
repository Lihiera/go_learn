package main

import (
	"fmt"
	"sort"
)

type intSlice []int

func (t *intSlice) Len() int {
	return len(*t)
}

func (t *intSlice) Less(i, j int) bool {
	fmt.Println((*t)[i], (*t)[j])
	return (*t)[i] < (*t)[j]
}

func (t *intSlice) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func IsPalindrome(s sort.Interface) bool {
	len := s.Len()
	for i := 0; i < len/2; i++ {
		if s.Less(i, len-i-1) || s.Less(len-1-i, i) {
			return false
		}
	}
	return true
}

func main() {
	sliceptr := intSlice([]int{2, 2, 5, 5, 2, 1})
	fmt.Println(IsPalindrome(&sliceptr))
}
