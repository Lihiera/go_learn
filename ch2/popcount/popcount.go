package popcount

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func PopCount1(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

func PopCount2(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount3(x uint64) (count int) {
	for i := 0; i < 64; i++ {
		if x&1 != 0 {
			count++
		}
		x >>= 1
	}
	return
}

func PopCount4(x uint64) (count int) {
	for x != 0 {
		x &= x - 1
		count++
	}
	return
}

func main() {
	s := os.Args[1]
	x, _ := strconv.ParseUint(s, 10, 64)
	start1 := time.Now()
	c1 := PopCount1(x)
	time1 := time.Since(start1).Seconds()
	start2 := time.Now()
	c2 := PopCount4(x)
	time2 := time.Since(start2).Seconds()
	fmt.Printf("PopCount1: %d %vs\nPopCount2: %d %vs\n", c1, time1, c2, time2)
}
