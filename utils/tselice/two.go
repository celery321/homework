package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr[2:5]
	fmt.Println(s, len(s), cap(s))

	s2 := s[2:7]
	fmt.Println(s2, len(s2), cap(s2))
}
