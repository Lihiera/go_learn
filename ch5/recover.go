package main

import "fmt"

func main() {
	fmt.Println(f())

}

func f() (re int) {
	defer func() {
		if p := recover(); p != nil {
			re = p.(int)
		}
	}()
	panic(3)
}
