package main

import "fmt"

func print() {
	fmt.Println("hello world")
}

func print1() {
	fmt.Println("print1")
}

func func1(n int) int {
	return n + 1

}

func main() {
	var f func()
	var flist [3]func(int) int
	//var fsilce  []func()
	//var fmap map[string]func()

	flist[0] = func1
	flist[0](10)
	f = print
	f()
	f = print1
	f()
	fmt.Println(f)
}
