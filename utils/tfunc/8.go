package main

import (
	"errors"
	"fmt"
)

func iter(s []int) func() (int, error) {
	var i = 0
	return func() (int, error) {
		if i >= len(s) {
			return 0, errors.New("end")
		}
		n := s[i]
		i += 1
		return n, nil
	}

}

func main() {
	f := iter([]int{1, 2, 3})

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	a := []int{}
	b := []int{1, 2, 3}
	a = append(a, b[:]...)
	fmt.Println(a)
}
