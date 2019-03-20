package main

import (
	"fmt"
	"strings"
)

func main() {
	var f func(int) int
	f = func(i int) int {
		return 10
	}
	fmt.Println(f(1))

	s := strings.Map(
		func(r rune) rune {
			return r - 32
		}, "hello")

	fmt.Println(s)
}
