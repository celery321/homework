package main

import "fmt"

func print() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	var p *int
	fmt.Println(*p)
}

func main() {
	print()
	//panic("bu 想执行了")

	//var i = 3
	//var slice [3]int
	//fmt.Println(slice[i])
}
