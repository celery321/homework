package main

import "fmt"

func main() {

	s1 := []int{1, 2, 4, 6}

	s1 = append(s1, 888)
	s1 = append(s1, 888)
	fmt.Println(s1)
}
