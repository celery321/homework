package main

import "fmt"

func main() {
	//var p *int = 10 // 野指针 无效地址
	//var p *int      // 未初始化指针
	//fmt.Println(&p)

	var p *int
	p = new(int)
	fmt.Println(*p)
}
