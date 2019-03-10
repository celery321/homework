package main

import "fmt"

type person struct {
	name string
	sex  byte
	age  int
}

type sayer interface {
	sayhello()
}

func (s *person) sayhello() {
	fmt.Println(s.name)
}

func swap(a, b int) {
	a, b = b, a
}

func swap2(x, y *int) {
	*x, *y = *y, *x
}
func main() {

	a := 10
	b := 20
	fmt.Println(a, b)

	swap(10, 20)
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println(a, b)
	p := person{"aaa", 1, 18}
	var s sayer
	fmt.Printf("%p %p %p %p", &p, &a, &s, sayer.sayhello)

}
