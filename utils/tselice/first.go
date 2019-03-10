package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//s := arr[1:3:5]
	//fmt.Println(s,len(s),cap(s))

	// cap 是根据原理的数字长度
	s := arr[1:5:7]
	fmt.Println(s, len(s), cap(s))
	// len = 4, cap = 7-1
	s2 := s[0:6]
	fmt.Println(s2, len(s2), cap(s2))

}
