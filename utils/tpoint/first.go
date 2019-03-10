package main

import "fmt"

func test() {
	b := 10
}

func main() {
	a := 10
	var p *int = &a

	a = 100
	*p = 250
	fmt.Println(a, *p)
	test()
	a = 1000
	fmt.Println(a, *p)

}
