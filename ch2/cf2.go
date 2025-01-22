package main

import (
	"cf/lengconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf2: %v\n", err)
			os.Exit(1)
		}
		m := lengconv.Meter(l)
		f := lengconv.Foot(l)
		fmt.Printf("%s = %s, %s = %s\n", m, lengconv.MtoF(m), f, lengconv.FtoM(f))
	}
}
