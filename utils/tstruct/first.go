package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name string
	sex  byte
	age  int
}

func test(p *Person) {
	p.sex = 'm'
	p.name = "name1"
	p.age = 33
}

func main() {
	var man Person = Person{"andy", 1, 18}
	fmt.Println(man, man.name)

	var man2 Person = Person{"andy", 1, 19}
	//man = man2
	if man == man2 {
		fmt.Println("ok")
	}

	// 函数传参
	var temp Person
	test(&temp)
	fmt.Println(temp, unsafe.Sizeof(temp))

	var p1 *Person = &Person{"n1", 'f', 19}
	fmt.Println(p1.age)

	var p2 *Person
	p2 = p1
	fmt.Println(p2)

	p3 := new(Person)
	p3.age = 11
	fmt.Println(p3)

}
