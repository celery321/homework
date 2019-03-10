package main

import "fmt"

type a struct {
	name string
}

func main() {
	// 1 自动推到
	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1)

	s2 := make([]int, 5, 10)
	fmt.Println(len(s2), cap(s2))

	s3 := make([]int, 7)
	fmt.Println(len(s3), cap(s3))

	s4 := make([]a, 10)
	fmt.Println(s4)

}
